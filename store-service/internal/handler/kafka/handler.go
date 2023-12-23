package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

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
			log.Println("error unmarshalling message:", err)
		}

		fmt.Printf("Msg: %d, %d, %d, %v\n", order.ID, order.UserId, order.ItemId, order.CreatedAt)

		h.service.Store(context.Background(), order)

		session.MarkMessage(message, "")
	}
	return nil
}
