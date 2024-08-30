package chat

import (
	"context"
	"errors"

	"github.com/s0vunia/chat_microservice/internal/model"
	"github.com/s0vunia/chat_microservice/internal/service/stream"
)

func (s *serv) Connect(chatId int64, userId int64, streamObj stream.Stream) error {
	var chatChan chan *model.MessageCreate
	err := s.txManager.ReadCommitted(streamObj.Context(), func(ctx context.Context) error {
		exists, err := s.participantRepository.CheckParticipantInChat(ctx, chatId, userId)
		if err != nil {
			return err
		}
		if !exists {
			return errors.New("user is not in chat")
		}
		chat, err := s.chatRepository.Get(ctx, chatId)
		if err != nil {
			return err
		}
		if chat == nil {
			return errors.New("chat not found")
		}
		s.mxChannel.RLock()
		var ok bool
		chatChan, ok = s.channels[chatId]
		if !ok {
			s.channels[chatId] = make(chan *model.MessageCreate, 100)
			chatChan = s.channels[chatId]
		}
		s.mxChannel.RUnlock()

		s.mxChat.Lock()
		if _, okChat := s.chats[chatId]; !okChat {
			s.chats[chatId] = &Chat{
				streams: make(map[int64]stream.Stream),
			}
		}
		s.mxChat.Unlock()
		s.chats[chatId].m.Lock()
		s.chats[chatId].streams[userId] = streamObj
		s.chats[chatId].m.Unlock()

		return nil
	})
	if err != nil {
		return err
	}

	for {
		select {
		case msg, okCh := <-chatChan:
			if !okCh {
				return nil
			}

			for _, st := range s.chats[chatId].streams {
				if err := st.Send(msg); err != nil {
					return err
				}
			}

		case <-streamObj.Context().Done():
			s.chats[chatId].m.Lock()
			delete(s.chats[chatId].streams, userId)
			s.chats[chatId].m.Unlock()
			return nil
		}
	}
}
