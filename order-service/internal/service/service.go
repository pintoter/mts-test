package service

import (
	"context"
	"time"

	"github.com/pintoter/mts-test/order-service/internal/entity"
	"github.com/pintoter/mts-test/order-service/internal/repository"
)

type Service struct {
	broker repository.MessageBroker
}

func New(broker repository.MessageBroker) *Service {
	return &Service{
		broker: broker,
	}
}

func (s *Service) CreateOrder(ctx context.Context, userId, itemId int64) error {
	order := entity.Order{
		UserId:    userId,
		ItemId:    itemId,
		CreatedAt: time.Now(),
	}

	err := s.broker.Publish(ctx, order)
	if err != nil {
		return err
	}

	return nil
}
