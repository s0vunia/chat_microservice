package service

import (
	"context"
	"github.com/s0vunia/chat_microservices_course_boilerplate/internal/model"
)

// ChatService represents a chat service.
type ChatService interface {
	Create(ctx context.Context, createChat *model.ChatCreate, createParticipants *model.ParticipantsCreate) (int64, error)
	SendMessage(ctx context.Context, createMessage *model.MessageCreate) (string, error)
	Delete(ctx context.Context, id int64) error
}
