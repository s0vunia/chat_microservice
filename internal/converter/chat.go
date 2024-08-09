package converter

import (
	"github.com/s0vunia/chat_microservice/internal/model"
	desc "github.com/s0vunia/chat_microservice/pkg/chat_v1"
)

// ToChatCreateFromDesc converts desc.ChatCreate to model.ChatCreate
func ToChatCreateFromDesc(chatCreate *desc.ChatCreate) *model.ChatCreate {
	return &model.ChatCreate{
		Name: chatCreate.Name,
	}
}
