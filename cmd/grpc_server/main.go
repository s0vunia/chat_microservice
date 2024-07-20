package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/gofrs/uuid"
	"github.com/s0vunia/chat_microservices_course_boilerplate/internal/config"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/s0vunia/chat_microservices_course_boilerplate/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	configPath           string
	countTriesToPostgres = 5
)

func init() {
	configPath = os.Getenv("CONFIG_PATH")
}

type server struct {
	desc.UnimplementedChatV1Server
	pool *pgxpool.Pool
}

// Create chat
func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Create, name: %s", req.GetName())

	if req.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "len of name must be positive")
	}

	builderInsert := sq.Insert("chats").
		PlaceholderFormat(sq.Dollar).
		Columns("name", "user_ids").
		Values(req.GetName(), req.GetUserIds()).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Printf("err: %e", err)
		return nil, status.Error(codes.Internal, "failed to build query")
	}
	var chatID int64
	err = s.pool.QueryRow(ctx, query, args...).Scan(&chatID)
	if err != nil {
		log.Printf("err: %e", err)
		return nil, status.Error(codes.Internal, "failed to insert chat")
	}

	log.Printf("inserted chat with id: %d", chatID)

	return &desc.CreateResponse{
		Id: chatID,
	}, nil
}

// Delete chat
func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Delete chat id: %d", req.GetId())

	builderDelete := sq.Delete("chats").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
		return nil, status.Error(codes.Internal, "failed to build query")
	}

	res, err := s.pool.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("failed to delete chat: %v", err)
		return nil, status.Error(codes.Internal, "failed to delete chat")
	}

	log.Printf("deleted %d rows", res.RowsAffected())
	return nil, nil
}

// SendMessage to chat
func (s *server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*desc.SendMessageResponse, error) {
	log.Printf("Send message from %d to %d", req.GetFromUserId(), req.GetToChatId())

	if req.GetText() == "" {
		return nil, status.Error(codes.InvalidArgument, "len of text must be positive")
	}

	builderInsert := sq.Insert("messages").
		PlaceholderFormat(sq.Dollar).
		Columns("chat_id", "user_id", "text").
		Values(req.GetToChatId(), req.GetFromUserId(), req.GetText()).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Printf("err: %e", err)
		return nil, status.Error(codes.Internal, "failed to build query")
	}
	var messageID uuid.UUID
	err = s.pool.QueryRow(ctx, query, args...).Scan(&messageID)
	if err != nil {
		log.Printf("err: %e", err)
		return nil, status.Error(codes.Internal, "failed to insert message")
	}

	log.Printf("inserted chat with id: %d", messageID)

	return &desc.SendMessageResponse{
		Id:     messageID.String(),
		ChatId: req.GetToChatId(),
	}, nil
}

func main() {
	flag.Parse()
	ctx := context.Background()

	// Считываем переменные окружения
	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := config.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	pgConfig, err := config.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Создаем пул соединений с базой данных
	var pool *pgxpool.Pool
	for i := 0; i < countTriesToPostgres; i++ {
		pool, err = pgxpool.Connect(ctx, pgConfig.DSN())
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database: %v. Attempt #%d\n", err, i+1)
		time.Sleep(2 * time.Second)
	}
	defer pool.Close()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{pool: pool})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
