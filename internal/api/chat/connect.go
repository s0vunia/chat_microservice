package chat

import (
	"strconv"

	"github.com/s0vunia/chat_microservice/internal/logger"
	desc2 "github.com/s0vunia/chat_microservice/internal/service/stream/desc"
	desc "github.com/s0vunia/chat_microservice/pkg/chat_v1"
	"go.uber.org/zap"
)

// ConnectChat creates connect to a chat
func (i *Implementation) ConnectChat(req *desc.ConnectChatRequest, stream desc.ChatV1_ConnectChatServer) error {
	logger.Info(
		"connecting to chat...",
		zap.String("chat_id", strconv.Itoa(int(req.GetChatId()))),
	)
	err := i.chatService.Connect(req.ChatId, req.UserId, desc2.NewStream(stream))
	if err != nil {
		return err
	}

	return nil
}
