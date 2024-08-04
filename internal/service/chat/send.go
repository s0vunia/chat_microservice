package chat

import (
	"context"

	"github.com/pkg/errors"

	"github.com/s0vunia/chat_microservice/internal/model"
)

func (s *serv) SendMessage(ctx context.Context, createMessage *model.MessageCreate) (string, error) {
	isExists, err := s.participantRepository.CheckParticipantInChat(ctx, createMessage.Info.ChatID, createMessage.Info.UserID)
	if !isExists || err != nil {
		return "", errors.New("user is not a participant of the chat")
	}
	id, err := s.messageRepository.Send(ctx, createMessage)
	if err != nil {
		return "", err
	}
	return id, nil
}
