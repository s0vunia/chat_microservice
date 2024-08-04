package chat

import (
	"context"

	"github.com/s0vunia/chat_microservice/internal/converter"
	desc "github.com/s0vunia/chat_microservice/pkg/chat_v1"
)

// SendMessage sends a new message
func (i *Implementation) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*desc.SendMessageResponse, error) {
	id, err := i.chatService.SendMessage(ctx, converter.ToMessageCreateFromDesc(req.GetMessage()))
	if err != nil {
		return nil, err
	}
	return &desc.SendMessageResponse{
		Id:     id,
		ChatId: req.GetMessage().GetToChatId(),
	}, nil
}
