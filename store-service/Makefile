MIGRATIONS_DIR = ./migrations

.PHONY: build
build:
	go build -o store-service ./cmd/store-service/main.go

.PHONY: build
run: build
	./store-service

.PHONY: migrations-create
migrations-create:
	migrate create -ext sql -dir $(MIGRATIONS_DIR) store

