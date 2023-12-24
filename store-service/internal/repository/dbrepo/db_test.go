package dbrepo

import (
	"context"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/pintoter/mts-test/store-service/internal/entity"
	"github.com/stretchr/testify/assert"
)

func Test_Save(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := New(db)

	type args struct {
		order entity.Order
	}

	type mockBehavior func(args args)

	tests := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		id           int
		wantErr      bool
	}{
		{
			name: "Success",
			mockBehavior: func(args args) {
				mock.ExpectBegin()

				expectedExec := "INSERT INTO store (user_id,item_id,created_at) VALUES ($1,$2,$3) RETURNING id"
				mock.ExpectQuery(regexp.QuoteMeta(expectedExec)).
					WithArgs(args.order.UserId, args.order.ItemId, args.order.CreatedAt).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectCommit()
			},
			args: args{
				order: entity.Order{
					UserId:    1,
					ItemId:    10,
					CreatedAt: time.Now().Round(time.Second),
				},
			},
			id: 1,
		},
		{
			name: "SuccessWithoutDate",
			mockBehavior: func(args args) {
				mock.ExpectBegin()

				expectedExec := "INSERT INTO store (user_id,item_id,created_at) VALUES ($1,$2,$3) RETURNING id"
				mock.ExpectQuery(regexp.QuoteMeta(expectedExec)).
					WithArgs(args.order.UserId, args.order.ItemId, args.order.CreatedAt).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectCommit()
			},
			args: args{
				order: entity.Order{
					UserId: 1,
					ItemId: 10,
				},
			},
			id: 1,
		},
		{
			name: "FailedWithEmpty",
			mockBehavior: func(args args) {
				mock.ExpectBegin()
				expectedExec := "INSERT INTO store (user_id,item_id,created_at) VALUES ($1,$2,$3) RETURNING id"
				mock.ExpectQuery(regexp.QuoteMeta(expectedExec)).
					WithArgs(args.order.UserId, args.order.ItemId, args.order.CreatedAt).WillReturnError(sqlmock.ErrCancelled)
				mock.ExpectRollback()
			},
			args: args{
				order: entity.Order{},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(tt.args)

			got, err := r.Save(context.Background(), tt.args.order)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.id, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
