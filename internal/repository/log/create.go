package log

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/s0vunia/chat_microservice/internal/client/db"
	"github.com/s0vunia/chat_microservice/internal/model"
)

// Create creates a new log.
func (r *repo) Create(ctx context.Context, logCreate *model.LogCreate) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(messageColumn).
		Values(logCreate.Message).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "log_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
