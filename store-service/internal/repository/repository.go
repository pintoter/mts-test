package repository

import (
	"context"

	"github.com/pintoter/mts-test/store-service/internal/entity"
)

type StoreRepository interface {
	Save(ctx context.Context, order entity.Order) (int, error)
}