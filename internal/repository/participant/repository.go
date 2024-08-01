package participant

import (
	"github.com/s0vunia/chat_microservices_course_boilerplate/internal/repository"
	"github.com/s0vunia/platform_common/pkg/db"
)

const (
	tableName = "chatparticipants"

	idColumn        = "id"
	chatIDColumn    = "chat_id"
	userIDColumn    = "user_id"
	createdAtColumn = "created_at"
)

type repo struct {
	db db.Client
}

// NewRepository creates a new user repository.
func NewRepository(db db.Client) repository.ParticipantRepository {
	return &repo{db: db}
}
