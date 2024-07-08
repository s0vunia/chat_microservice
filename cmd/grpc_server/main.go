package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/s0vunia/chat_microservices_course_boilerplate/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50052

type server struct {
	desc.UnimplementedChatV1Server
}

// Create chat
func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Create, usernames: %s", req.GetUsernames())
	log.Printf("ctx: %+v", ctx)

	return &desc.CreateResponse{
		Id: int64(gofakeit.Number(1, 100)),
	}, nil
}

// Delete chat
func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Delete id %d", req.GetId())
	log.Printf("ctx: %+v", ctx)

	return nil, nil
}

// SendMessage to chat
func (s *server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("Send message %s from %s", req.GetText(), req.GetFrom())
	log.Printf("ctx: %+v", ctx)

	return nil, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
