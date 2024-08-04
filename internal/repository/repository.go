package repository

import (
	"context"

	"github.com/s0vunia/chat_microservice/internal/model"
)

// ChatRepository represents a chat repository.
type ChatRepository interface {
	Create(context context.Context, createChat *model.ChatCreate) (int64, error)
	Delete(context context.Context, id int64) error
}

// MessageRepository represents a message repository.
type MessageRepository interface {
	Send(context context.Context, createMessage *model.MessageCreate) (string, error)
}

// ParticipantRepository represents a participant repository.
type ParticipantRepository interface {
	CreateParticipant(context context.Context, createParticipant *model.ParticipantCreate) error
	CreateParticipants(context context.Context, createParticipants *model.ParticipantsCreate) error
	CheckParticipantInChat(context context.Context, chatID int64, userID int64) (bool, error)
}
