package chat

import (
	"context"
	"errors"
	"fmt"

	"github.com/s0vunia/chat_microservice/internal/model"
)

func (s *serv) Create(ctx context.Context, createChat *model.ChatCreate, createParticipants *model.ParticipantsCreate) (int64, error) {
	var id int64

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		userIDs := make([]int64, 0, len(createParticipants.Participants))
		for _, participant := range createParticipants.Participants {
			userIDs = append(userIDs, participant.UserID)
		}
		exists, errTx := s.authServiceClient.IsUserExists(ctx, userIDs)
		if errTx != nil {
			return errTx
		}
		if !exists {
			return errors.New("some users do not exist")
		}

		id, errTx = s.chatRepository.Create(ctx, createChat)
		if errTx != nil {
			return errTx
		}
		for i := 0; i < len(createParticipants.Participants); i++ {
			createParticipants.Participants[i].ChatID = id
		}

		s.mxChannel.Lock()
		s.channels[id] = make(chan *model.MessageCreate, 100)
		s.mxChannel.Unlock()

		errTx = s.participantRepository.CreateParticipants(ctx, createParticipants)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.logsRepository.Create(ctx, &model.LogCreate{
			Message: fmt.Sprintf("chat %d created", id),
		})

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
