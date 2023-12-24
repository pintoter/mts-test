package service

import (
	"context"
	"log"

	"github.com/pintoter/mts-test/store-service/internal/entity"
	"github.com/pintoter/mts-test/store-service/internal/repository"
)

type Service struct {
	repo repository.StoreRepository
}

func New(repo repository.StoreRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Store(ctx context.Context, order entity.Order) error {
	id, err := s.repo.Save(ctx, order)
	if err != nil {
		return err
	}
	log.Printf("store service: successfully writing new order in DB: %d\n", id)
	return nil
}
