package converter

import (
	"github.com/s0vunia/chat_microservices_course_boilerplate/internal/model"
)

// ToParticipantsCreateFromDesc converts desc.ParticipantsCreate to model.ParticipantsCreate
func ToParticipantsCreateFromDesc(userIDs []int64) *model.ParticipantsCreate {
	participantsCreate := &model.ParticipantsCreate{
		Participants: make([]model.ParticipantCreate, 0, len(userIDs)),
	}

	for _, userID := range userIDs {
		participantsCreate.Participants = append(participantsCreate.Participants, model.ParticipantCreate{
			UserID: userID,
		})
	}

	return participantsCreate
}
