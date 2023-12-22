package broker

import (
	"log"

	"gopkg.in/Shopify/sarama.v1"
)

func NewSyncProducer() sarama.SyncProducer {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{"localhost:9095"}, cfg)
	if err != nil {
		log.Fatal("Init kafka producer err:", err)
	}

	return producer
}
