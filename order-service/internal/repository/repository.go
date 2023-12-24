package repository

import (
	"context"

	"github.com/pintoter/mts-test/order-service/internal/entity"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type MessageBroker interface {
	Publish(ctx context.Context, order entity.Order) error
}
