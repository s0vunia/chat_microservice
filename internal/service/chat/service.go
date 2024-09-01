package chat

import (
	"sync"

	"github.com/s0vunia/chat_microservice/internal/client/authservice"
	"github.com/s0vunia/chat_microservice/internal/model"
	"github.com/s0vunia/chat_microservice/internal/repository"
	"github.com/s0vunia/chat_microservice/internal/service"
	"github.com/s0vunia/chat_microservice/internal/service/stream"
	"github.com/s0vunia/platform_common/pkg/db"
)

type Chat struct {
	streams map[int64]stream.Stream
	m       sync.RWMutex
}

type serv struct {
	chatRepository        repository.ChatRepository
	messageRepository     repository.MessageRepository
	participantRepository repository.ParticipantRepository
	authServiceClient     authservice.AuthService
	logsRepository        repository.LogRepository
	txManager             db.TxManager

	chats  map[int64]*Chat
	mxChat sync.RWMutex

	channels  map[int64]chan *model.MessageCreate
	mxChannel sync.RWMutex
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
		chats:                 make(map[int64]*Chat),
		channels:              make(map[int64]chan *model.MessageCreate),
	}
}
