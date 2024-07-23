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

// ChatCreate represents a chat to be created
type ChatCreate struct {
	Name string
}

// ChatUpdate represents a chat to be updated
type ChatUpdate struct {
	Name *string
}
