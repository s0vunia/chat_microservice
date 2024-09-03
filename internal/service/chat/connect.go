package chat

import (
	"context"
	"errors"

	"github.com/s0vunia/chat_microservice/internal/model"
	"github.com/s0vunia/chat_microservice/internal/service/stream"
)

func (s *serv) Connect(chatID int64, userID int64, streamObj stream.Stream) error {
	var chatChan chan *model.MessageCreate
	err := s.txManager.ReadCommitted(streamObj.Context(), func(ctx context.Context) error {
		exists, err := s.participantRepository.CheckParticipantInChat(ctx, chatID, userID)
		if err != nil {
			return err
		}
		if !exists {
			return errors.New("user is not in chat")
		}
		chat, err := s.chatRepository.Get(ctx, chatID)
		if err != nil {
			return err
		}
		if chat == nil {
			return errors.New("chat not found")
		}
		s.mxChannel.RLock()
		var ok bool
		chatChan, ok = s.channels[chatID]
		if !ok {
			s.channels[chatID] = make(chan *model.MessageCreate, 100)
			chatChan = s.channels[chatID]
		}
		s.mxChannel.RUnlock()

		s.mxChat.Lock()
		if _, okChat := s.chats[chatID]; !okChat {
			s.chats[chatID] = &Chat{
				streams: make(map[int64]stream.Stream),
			}
		}
		s.mxChat.Unlock()
		s.chats[chatID].m.Lock()
		s.chats[chatID].streams[userID] = streamObj
		s.chats[chatID].m.Unlock()

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

			for _, st := range s.chats[chatID].streams {
				if err := st.Send(msg); err != nil {
					return err
				}
			}

		case <-streamObj.Context().Done():
			s.chats[chatID].m.Lock()
			delete(s.chats[chatID].streams, userID)
			s.chats[chatID].m.Unlock()
			return nil
		}
	}
}
