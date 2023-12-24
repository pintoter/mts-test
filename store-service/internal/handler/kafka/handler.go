package kafka

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/pintoter/mts-test/store-service/internal/entity"
	"github.com/pintoter/mts-test/store-service/internal/service"
	"gopkg.in/Shopify/sarama.v1"
)

type kafkaHandler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *kafkaHandler {
	return &kafkaHandler{
		service: service,
	}
}

func (h *kafkaHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h *kafkaHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h *kafkaHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		var order entity.Order
		err := json.Unmarshal(message.Value, &order)
		if err != nil {
			log.Println("store service: error unmarshalling message:", err)
		}

		log.Printf("store service: got new message: userId - %d, itemId - %d, time - %v\n", order.UserId, order.ItemId, order.CreatedAt)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		err = h.service.Store(ctx, order)
		if err != nil {
			log.Println("store-service: handler err:", err)
		}

		session.MarkMessage(message, "")

		cancel()
	}

	return nil
}
