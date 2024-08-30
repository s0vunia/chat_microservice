package chat

import (
	"log"

	desc2 "github.com/s0vunia/chat_microservice/internal/service/stream/desc"
	desc "github.com/s0vunia/chat_microservice/pkg/chat_v1"
)

// ConnectChat creates connect to a chat
func (i *Implementation) ConnectChat(req *desc.ConnectChatRequest, stream desc.ChatV1_ConnectChatServer) error {
	err := i.chatService.Connect(req.ChatId, req.UserId, desc2.NewStream(stream))
	if err != nil {
		return err
	}

	log.Printf("connect chat: %d", req.ChatId)

	return nil
}
