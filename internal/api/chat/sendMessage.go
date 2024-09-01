package chat

import (
	"context"
	"strconv"

	"github.com/s0vunia/chat_microservice/internal/converter"
	"github.com/s0vunia/chat_microservice/internal/logger"
	desc "github.com/s0vunia/chat_microservice/pkg/chat_v1"
	"go.uber.org/zap"
)

// SendMessage sends a new message
func (i *Implementation) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*desc.SendMessageResponse, error) {
	logger.Info(
		"sending message...",
		zap.String("chat_id", strconv.Itoa(int(req.GetMessage().GetToChatId()))),
		zap.String("message", req.GetMessage().GetText()),
		zap.String("sender_id", strconv.Itoa(int(req.GetMessage().FromUserId))),
	)
	id, err := i.chatService.SendMessage(ctx, converter.ToMessageCreateFromDesc(req.GetMessage()))
	if err != nil {
		return nil, err
	}
	return &desc.SendMessageResponse{
		Id:     id,
		ChatId: req.GetMessage().GetToChatId(),
	}, nil
}
