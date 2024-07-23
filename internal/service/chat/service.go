package chat

import (
	"github.com/s0vunia/chat_microservices_course_boilerplate/internal/client/db"
	"github.com/s0vunia/chat_microservices_course_boilerplate/internal/repository"
	"github.com/s0vunia/chat_microservices_course_boilerplate/internal/service"
)

type serv struct {
	chatRepository        repository.ChatRepository
	messageRepository     repository.MessageRepository
	participantRepository repository.ParticipantRepository
	txManager             db.TxManager
}

// NewService creates a new chat service.
func NewService(
	chatRepository repository.ChatRepository,
	messageRepository repository.MessageRepository,
	participantRepository repository.ParticipantRepository,
	txManager db.TxManager,
) service.ChatService {
	return &serv{
		chatRepository:        chatRepository,
		messageRepository:     messageRepository,
		participantRepository: participantRepository,
		txManager:             txManager,
	}
}
