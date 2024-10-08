package chat

import (
	"github.com/s0vunia/chat_microservice/internal/client/authservice"
	"github.com/s0vunia/chat_microservice/internal/client/db"
	"github.com/s0vunia/chat_microservice/internal/repository"
	"github.com/s0vunia/chat_microservice/internal/service"
)

type serv struct {
	chatRepository        repository.ChatRepository
	messageRepository     repository.MessageRepository
	participantRepository repository.ParticipantRepository
	authServiceClient     authservice.AuthService
	logsRepository        repository.LogRepository
	txManager             db.TxManager
}

// NewService creates a new chat service.
func NewService(
	chatRepository repository.ChatRepository,
	messageRepository repository.MessageRepository,
	participantRepository repository.ParticipantRepository,
	authServiceClient authservice.AuthService,
	logsRepository repository.LogRepository,
	txManager db.TxManager,
) service.ChatService {
	return &serv{
		chatRepository:        chatRepository,
		messageRepository:     messageRepository,
		participantRepository: participantRepository,
		authServiceClient:     authServiceClient,
		logsRepository:        logsRepository,
		txManager:             txManager,
	}
}
