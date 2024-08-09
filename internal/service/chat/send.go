package chat

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/s0vunia/chat_microservice/internal/model"
)

func (s *serv) SendMessage(ctx context.Context, createMessage *model.MessageCreate) (string, error) {
	var id string
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		isExists, errTx := s.participantRepository.CheckParticipantInChat(ctx, createMessage.Info.ChatID, createMessage.Info.UserID)
		if !isExists || errTx != nil {
			return errors.New("user is not a participant of the chat")
		}
		id, errTx = s.messageRepository.Send(ctx, createMessage)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.logsRepository.Create(ctx, &model.LogCreate{
			Message: fmt.Sprintf("message %s created", id),
		})
		if errTx != nil {
			return errTx
		}

		return nil
	})
	if err != nil {
		return "", err
	}
	return id, nil
}
