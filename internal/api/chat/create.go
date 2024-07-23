package chat

import (
	"context"
	"github.com/s0vunia/chat_microservices_course_boilerplate/internal/converter"
	desc "github.com/s0vunia/chat_microservices_course_boilerplate/pkg/chat_v1"
	"log"
)

// Create creates a new chat
func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.chatService.Create(ctx, converter.ToChatCreateFromDesc(req.GetChat()), converter.ToParticipantsCreateFromDesc(req.GetUserIds()))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted chat with id: %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
