package converter

import (
	"github.com/s0vunia/chat_microservices_course_boilerplate/internal/model"
	desc "github.com/s0vunia/chat_microservices_course_boilerplate/pkg/chat_v1"
)

func ToMessageCreateFromDesc(messageCreate *desc.MessageCreate) *model.MessageCreate {
	return &model.MessageCreate{
		Info: model.MessageInfo{
			ChatID: messageCreate.ToChatId,
			UserID: messageCreate.FromUserId,
			Text:   messageCreate.Text,
		},
	}
}
