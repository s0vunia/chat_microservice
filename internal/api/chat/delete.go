package chat

import (
	"context"
	"strconv"

	"github.com/s0vunia/chat_microservice/internal/logger"
	desc "github.com/s0vunia/chat_microservice/pkg/chat_v1"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Delete deletes a chat
func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	logger.Info(
		"deleting chat...",
		zap.String("chat_id", strconv.Itoa(int(req.GetId()))),
	)
	err := i.chatService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return nil, nil
}
