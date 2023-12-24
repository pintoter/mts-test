package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pintoter/mts-test/store-service/internal/entity"
	mock_repository "github.com/pintoter/mts-test/store-service/internal/repository/mocks"
	"github.com/stretchr/testify/require"
)

func TestCreateOrder(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_repository.NewMockStoreRepository(ctl)

	ctx := context.Background()

	repo.EXPECT().Save(ctx, gomock.Any())
	service := New(repo)
	err := service.Store(ctx, entity.Order{})
	require.NoError(t, err)
}
