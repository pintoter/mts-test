package repository

import (
	"context"

	"github.com/pintoter/mts-test/order-service/internal/entity"
)

type MessageBroker interface {
	Publish(ctx context.Context, order entity.Order) error
}
