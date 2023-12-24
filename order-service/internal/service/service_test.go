package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	mock_repository "github.com/pintoter/mts-test/order-service/internal/repository/mocks"
	"github.com/stretchr/testify/require"
)

func TestCreateOrder(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_repository.NewMockMessageBroker(ctl)

	ctx := context.Background()

	repo.EXPECT().Publish(ctx, gomock.Any())
	service := New(repo)
	err := service.CreateOrder(ctx, 1, 1)
	require.NoError(t, err)
}
