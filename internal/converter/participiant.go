package converter

import (
	"github.com/s0vunia/chat_microservices_course_boilerplate/internal/model"
)

func ToParticipantsCreateFromDesc(userIds []int64) *model.ParticipantsCreate {
	participantsCreate := &model.ParticipantsCreate{
		Participants: make([]model.ParticipantCreate, 0, len(userIds)),
	}

	for _, userID := range userIds {
		participantsCreate.Participants = append(participantsCreate.Participants, model.ParticipantCreate{
			UserID: userID,
		})
	}

	return participantsCreate
}
