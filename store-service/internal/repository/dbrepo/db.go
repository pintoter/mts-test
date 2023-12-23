package dbrepo

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/pintoter/mts-test/store-service/internal/entity"
	"github.com/pintoter/mts-test/store-service/internal/repository"
)

const (
	store = "store"

	userId    = "user_id"
	itemId    = "item_id"
	createdAt = "created_at"
)

type dbrepo struct {
	db *sql.DB
}

func New(db *sql.DB) repository.StoreRepository {
	return &dbrepo{db: db}
}

func (r *dbrepo) Save(ctx context.Context, order entity.Order) (int, error) {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
	})
	if err != nil {
		return 0, err
	}
	defer func() { _ = tx.Rollback() }()

	builder := sq.Insert(store).
		Columns(userId, itemId, createdAt).
		Values(order.UserId, order.ItemId, order.CreatedAt).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	var orderId int
	err = tx.QueryRowContext(ctx, query, args...).Scan(&orderId)
	if err != nil {
		return 0, err
	}

	return orderId, tx.Commit()
}
