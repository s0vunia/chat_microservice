package model

// Participant represents a chat participant
type Participant struct {
	ChatID int64
	UserID int64
}

// ParticipantCreate represents a chat participant to be created
type ParticipantCreate struct {
	ChatID int64
	UserID int64
}

// ParticipantsCreate represents a chat participants to be created
type ParticipantsCreate struct {
	Participants []ParticipantCreate
}
