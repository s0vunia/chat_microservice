package converter

import (
	"github.com/s0vunia/chat_microservice/internal/model"
	desc "github.com/s0vunia/chat_microservice/pkg/chat_v1"
)

// ToMessageCreateFromDesc converts desc.MessageCreate to model.MessageCreate
func ToMessageCreateFromDesc(messageCreate *desc.MessageCreate) *model.MessageCreate {
	return &model.MessageCreate{
		Info: model.MessageInfo{
			ChatID: messageCreate.ToChatId,
			UserID: messageCreate.FromUserId,
			Text:   messageCreate.Text,
		},
	}
}
