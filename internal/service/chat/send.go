package chat

import (
	"context"
	"github.com/s0vunia/chat_microservices_course_boilerplate/internal/model"
)

func (s *serv) SendMessage(ctx context.Context, createMessage *model.MessageCreate) (string, error) {
	isExists, err := s.participantRepository.CheckParticipantInChat(ctx, createMessage.Info.ChatID, createMessage.Info.UserID)
	if !isExists || err != nil {
		return "", err
	}
	id, err := s.messageRepository.Send(ctx, createMessage)
	if err != nil {
		return "", err
	}
	return id, nil
}
