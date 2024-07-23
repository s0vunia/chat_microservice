package chat

import (
	"context"

	desc "github.com/s0vunia/chat_microservices_course_boilerplate/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Delete deletes a chat
func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := i.chatService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return nil, nil
}
