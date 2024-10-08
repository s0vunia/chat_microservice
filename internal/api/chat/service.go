package chat

import (
	"github.com/s0vunia/chat_microservice/internal/service"
	desc "github.com/s0vunia/chat_microservice/pkg/chat_v1"
)

// Implementation represents a chat API implementation.
type Implementation struct {
	desc.UnimplementedChatV1Server
	chatService service.ChatService
}

// NewImplementation creates a new chat API implementation.
func NewImplementation(chatService service.ChatService) *Implementation {
	return &Implementation{
		chatService: chatService,
	}
}
