package model

import (
	"database/sql"
	"time"
)

// Message represents a chat message
type Message struct {
	ID        int64
	Info      MessageInfo
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

// MessageInfo represents a chat message info
type MessageInfo struct {
	ChatID int64
	UserID int64
	Text   string
}

// MessageCreate represents a chat message to be created
type MessageCreate struct {
	Info MessageInfo
}

// MessageUpdate represents a chat message to be updated
type MessageUpdate struct {
	ChatID *int64
	UserID *int64
	Text   *string
}
