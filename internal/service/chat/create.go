package chat

import (
	"context"

	"github.com/s0vunia/chat_microservices_course_boilerplate/internal/model"
)

func (s *serv) Create(ctx context.Context, createChat *model.ChatCreate, createParticipants *model.ParticipantsCreate) (int64, error) {
	var id int64

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.chatRepository.Create(ctx, createChat)
		if errTx != nil {
			return errTx
		}
		for i := 0; i < len(createParticipants.Participants); i++ {
			createParticipants.Participants[i].ChatID = id
		}
		errTx = s.participantRepository.CreateParticipants(ctx, createParticipants)
		if errTx != nil {
			return errTx
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return id, nil
}
