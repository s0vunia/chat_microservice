package service

import (
	"context"

	"github.com/s0vunia/chat_microservice/internal/model"
	"github.com/s0vunia/chat_microservice/internal/service/stream"
)

// ChatService represents a chat service.
type ChatService interface {
	Create(ctx context.Context, createChat *model.ChatCreate, createParticipants *model.ParticipantsCreate) (int64, error)
	SendMessage(ctx context.Context, createMessage *model.MessageCreate) (string, error)
	Connect(chatID int64, userID int64, stream stream.Stream) error
	Delete(ctx context.Context, id int64) error
}
