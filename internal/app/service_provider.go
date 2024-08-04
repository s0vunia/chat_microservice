package app

import (
  "context"
	"log"

	"github.com/s0vunia/platform_common/pkg/closer"
	"github.com/s0vunia/platform_common/pkg/db"
	"github.com/s0vunia/platform_common/pkg/db/pg"
	"github.com/s0vunia/platform_common/pkg/db/transaction"

	"github.com/s0vunia/chat_microservice/internal/client/authservice"
	authService2 "github.com/s0vunia/chat_microservice/internal/client/authservice/authservice"

	"github.com/s0vunia/chat_microservice/internal/api/chat"
	"github.com/s0vunia/chat_microservice/internal/config"
	"github.com/s0vunia/chat_microservice/internal/repository"
	chatRepository "github.com/s0vunia/chat_microservice/internal/repository/chat"
	messageRepository "github.com/s0vunia/chat_microservice/internal/repository/message"
	participantRepository "github.com/s0vunia/chat_microservice/internal/repository/participant"
	"github.com/s0vunia/chat_microservice/internal/service"
	chatService "github.com/s0vunia/chat_microservice/internal/service/chat"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig
	authConfig config.AuthServiceConfig

	dbClient              db.Client
	authService           authservice.AuthService
	txManager             db.TxManager
	chatRepository        repository.ChatRepository
	messageRepository     repository.MessageRepository
	participantRepository repository.ParticipantRepository
	logsRepository        repository.LogRepository

	chatService service.ChatService

	chatImpl *chat.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) AuthServiceConfig() config.AuthServiceConfig {
	if s.authConfig == nil {
		cfg, err := config.NewAuthServiceConfig()
		if err != nil {
			log.Fatalf("failed to get auth service config: %s", err.Error())
		}

		s.authConfig = cfg
	}

	return s.authConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) AuthService(_ context.Context) authservice.AuthService {
	if s.authService == nil {
		var err error
		s.authService, err = authService2.NewClient(s.AuthServiceConfig().Address())
		if err != nil {
			log.Fatalf("failed to create auth service: %s", err.Error())
		}
	}
	return s.authService
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = chatRepository.NewRepository(s.DBClient(ctx))
	}

	return s.chatRepository
}

func (s *serviceProvider) MessageRepository(ctx context.Context) repository.MessageRepository {
	if s.messageRepository == nil {
		s.messageRepository = messageRepository.NewRepository(s.DBClient(ctx))
	}

	return s.messageRepository
}

func (s *serviceProvider) ParticipantRepository(ctx context.Context) repository.ParticipantRepository {
	if s.participantRepository == nil {
		s.participantRepository = participantRepository.NewRepository(s.DBClient(ctx))
	}

	return s.participantRepository
}

func (s *serviceProvider) LogsRepository(ctx context.Context) repository.LogRepository {
	if s.logsRepository == nil {
		s.logsRepository = logsRepository.NewRepository(s.DBClient(ctx))
	}

	return s.logsRepository
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewService(
			s.ChatRepository(ctx),
			s.MessageRepository(ctx),
			s.ParticipantRepository(ctx),
			s.AuthService(ctx),
			s.LogsRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.chatService
}

func (s *serviceProvider) ChatImpl(ctx context.Context) *chat.Implementation {
	if s.chatImpl == nil {
		s.chatImpl = chat.NewImplementation(s.ChatService(ctx))
	}

	return s.chatImpl
}
