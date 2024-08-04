package chat

import (
	"context"
	"fmt"

	"github.com/s0vunia/chat_microservices_course_boilerplate/internal/model"
)

func (s *serv) Delete(ctx context.Context, id int64) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		errTx = s.chatRepository.Delete(ctx, id)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.logsRepository.Create(ctx, &model.LogCreate{
			Message: fmt.Sprintf("chat %d deleted", id),
		})

		if errTx != nil {
			return errTx
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
