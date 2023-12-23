package postgres

import (
	"database/sql"
	"time"
)

const (
	driverName = "postgres"
)

type Config interface {
	GetDSN() string
	GetMaxOpenConns() int
	GetMaxIdleConns() int
	GetConnMaxIdleTime() time.Duration
	GetConnMaxLifetime() time.Duration
}

func New(cfg Config) (*sql.DB, error) {
	db, err := sql.Open(driverName, cfg.GetDSN())
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.GetMaxOpenConns())
	db.SetMaxIdleConns(cfg.GetMaxIdleConns())
	db.SetConnMaxIdleTime(cfg.GetConnMaxIdleTime())
	db.SetConnMaxLifetime(cfg.GetConnMaxLifetime())

	return db, nil
}
