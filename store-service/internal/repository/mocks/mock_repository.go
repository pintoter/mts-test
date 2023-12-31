// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/pintoter/mts-test/store-service/internal/entity"
)

// MockStoreRepository is a mock of StoreRepository interface.
type MockStoreRepository struct {
	ctrl     *gomock.Controller
	recorder *MockStoreRepositoryMockRecorder
}

// MockStoreRepositoryMockRecorder is the mock recorder for MockStoreRepository.
type MockStoreRepositoryMockRecorder struct {
	mock *MockStoreRepository
}

// NewMockStoreRepository creates a new mock instance.
func NewMockStoreRepository(ctrl *gomock.Controller) *MockStoreRepository {
	mock := &MockStoreRepository{ctrl: ctrl}
	mock.recorder = &MockStoreRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStoreRepository) EXPECT() *MockStoreRepositoryMockRecorder {
	return m.recorder
}

// Save mocks base method.
func (m *MockStoreRepository) Save(ctx context.Context, order entity.Order) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, order)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockStoreRepositoryMockRecorder) Save(ctx, order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockStoreRepository)(nil).Save), ctx, order)
}
