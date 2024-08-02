package participant

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/s0vunia/chat_microservices_course_boilerplate/internal/client/db"
	"github.com/s0vunia/chat_microservices_course_boilerplate/internal/model"
)

func (r *repo) CreateParticipant(ctx context.Context, createParticipant *model.ParticipantCreate) error {
	builderInsert := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(chatIDColumn, userIDColumn).
		Values(createParticipant.ChatID, createParticipant.UserID)

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "participant_repository.CreateParticipant",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) CreateParticipants(ctx context.Context, createParticipants *model.ParticipantsCreate) error {
	for i := 0; i < len(createParticipants.Participants); i++ {
		err := r.CreateParticipant(ctx, &createParticipants.Participants[i])
		if err != nil {
			return err
		}
	}
	return nil
}
