package participant

import (
	"context"
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/s0vunia/platform_common/pkg/db"
)

func (r *repo) CheckParticipantInChat(ctx context.Context, chatID int64, userID int64) (bool, error) {
	builderSelect := sq.Select("1").
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{chatIDColumn: chatID, userIDColumn: userID}).
		Limit(1)

	query, args, err := builderSelect.ToSql()
	if err != nil {
		return false, err
	}

	q := db.Query{
		Name:     "participant_repository.CheckParticipantInChat",
		QueryRaw: query,
	}

	var exists int
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return exists == 1, nil
}
