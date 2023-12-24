include .env

MIGRATIONS_DIR = ./migrations
POSTGRES_DSN = postgres://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

.PHONY: build
build:
	go build -o store-service ./cmd/store-service/main.go

.PHONY: build
run: build
	./store-service

.PHONY: migrations-create
migrations-create:
	migrate create -ext sql -dir $(MIGRATIONS_DIR) store

.PHONY: migrations-up
migrations-up:
	migrate -path $(MIGRATIONS_DIR) -database $(POSTGRES_DSN) up

.PHONY: migrations-down
migrations-down:
	migrate -path $(MIGRATIONS_DIR) -database $(POSTGRES_DSN) down