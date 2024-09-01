package chat

import (
	"context"

	"github.com/s0vunia/chat_microservice/internal/converter"
	"github.com/s0vunia/chat_microservice/internal/logger"
	desc "github.com/s0vunia/chat_microservice/pkg/chat_v1"
	"go.uber.org/zap"
)

// Create creates a new chat
func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	logger.Info(
		"creating chat...",
		zap.String("chat name", req.GetChat().Name),
	)
	id, err := i.chatService.Create(ctx, converter.ToChatCreateFromDesc(req.GetChat()), converter.ToParticipantsCreateFromDesc(req.GetUserIds()))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
