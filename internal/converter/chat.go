package converter

import (
	"github.com/s0vunia/chat_microservices_course_boilerplate/internal/model"
	desc "github.com/s0vunia/chat_microservices_course_boilerplate/pkg/chat_v1"
)

func ToChatCreateFromDesc(chatCreate *desc.ChatCreate) *model.ChatCreate {
	return &model.ChatCreate{
		Name: chatCreate.Name,
	}
}
