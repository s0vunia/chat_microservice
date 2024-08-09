package message

import (
	"github.com/s0vunia/chat_microservice/internal/client/db"
	"github.com/s0vunia/chat_microservice/internal/repository"
)

const (
	tableName = "messages"

	idColumn        = "id"
	chatIDColumn    = "chat_id"
	userIDColumn    = "user_id"
	textColumn      = "text"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

// NewRepository creates a new user repository.
func NewRepository(db db.Client) repository.MessageRepository {
	return &repo{db: db}
}
