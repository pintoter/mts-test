package app

import (
	"log"

	"github.com/pintoter/mts-test/order-service/internal/broker"
	"github.com/pintoter/mts-test/order-service/internal/config"
	server "github.com/pintoter/mts-test/order-service/internal/server/grpc"
	"github.com/pintoter/mts-test/order-service/internal/service"
)

func Run() {
	cfg := config.Read()
	log.Println("order-service: success reading config")

	producer := broker.NewSyncProducer()
	log.Println("order-service: success creating producer")

	orderService := service.New(producer)
	log.Println("order-service: success creating service")

	if err := server.New(orderService).Run(&cfg.Grpc); err != nil {
		log.Fatal(err)
	}
}
