package chat

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/s0vunia/chat_microservice/internal/model"
	"github.com/s0vunia/platform_common/pkg/db"
)

func (r *repo) Create(ctx context.Context, createChat *model.ChatCreate) (int64, error) {
	builderInsert := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn).
		Values(createChat.Name).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "chat_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
