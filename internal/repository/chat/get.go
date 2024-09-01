package chat

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/s0vunia/chat_microservice/internal/model"
	"github.com/s0vunia/platform_common/pkg/db"
)

func (r *repo) Get(ctx context.Context, id int64) (*model.Chat, error) {
	builderSelect := sq.Select(idColumn, nameColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "chat_repository.Get",
		QueryRaw: query,
	}

	var chat model.Chat
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&chat.ID, &chat.Name)
	if err != nil {
		return nil, err
	}

	return &chat, nil
}
