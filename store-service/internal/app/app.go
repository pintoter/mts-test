package app

import (
	"context"
	"log"

	"github.com/pintoter/mts-test/store-service/internal/config"
	"github.com/pintoter/mts-test/store-service/internal/handler/kafka"
	"github.com/pintoter/mts-test/store-service/internal/migrations"
	"github.com/pintoter/mts-test/store-service/internal/repository/dbrepo"
	"github.com/pintoter/mts-test/store-service/internal/service"
	"github.com/pintoter/mts-test/store-service/pkg/database/postgres"
)

func Run() {
	cfg := config.Read()
	log.Println("store-service: success reading config")

	err := migrations.Do(&cfg.DB)
	if err != nil {
		log.Fatal("store-service: migrations failed:", err)
	}

	db, err := postgres.New(&cfg.DB)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("store-service: success connecting to db")

	storeRepo := dbrepo.New(db)
	storeService := service.New(storeRepo)
	handler := kafka.NewHandler(storeService)
	log.Println("store-service: init handler")

	err = kafka.StartConsuming(context.Background(), &cfg.Kafka, handler)
	for {
	}
}
