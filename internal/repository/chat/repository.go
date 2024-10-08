package chat

import (
	"github.com/s0vunia/chat_microservice/internal/client/db"
	"github.com/s0vunia/chat_microservice/internal/repository"
)

const (
	tableName = "chats"

	idColumn        = "id"
	nameColumn      = "name"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

// NewRepository creates a new user repository.
func NewRepository(db db.Client) repository.ChatRepository {
	return &repo{db: db}
}
