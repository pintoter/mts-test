package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/pintoter/mts-test/order-service/internal/entity"
	sarama "gopkg.in/Shopify/sarama.v1"
)

type Service struct {
	producer sarama.SyncProducer
}

func New(producer sarama.SyncProducer) *Service {
	return &Service{
		producer: producer,
	}
}

func (s *Service) CreateOrder(ctx context.Context, userId, itemId int64) error {
	order := &entity.Order{
		UserId: userId,
		ItemId: itemId,
	}

	o, err := json.Marshal(order)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: "order_created",
		Key:   sarama.StringEncoder(fmt.Sprintf("%d", userId)),
		Value: sarama.StringEncoder(o),
	}

	partition, offset, err := s.producer.SendMessage(msg)
	if err != nil {
		return err
	}
	log.Printf("message sent to partition %d at offset %d\n", partition, offset)

	return nil
}
