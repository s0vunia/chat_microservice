package model

import (
	"database/sql"
	"time"
)

// Chat represents a chat
type Chat struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
