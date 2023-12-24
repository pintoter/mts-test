package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/pintoter/mts-test/order-service/internal/entity"
	"gopkg.in/Shopify/sarama.v1"
)

type kafkaProducer struct {
	producer sarama.SyncProducer
	topic    string
}

type Config interface {
	GetBrokers() []string
	GetTopic() string
}

func NewProducer(cfg Config) *kafkaProducer {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(cfg.GetBrokers(), config)
	if err != nil {
		log.Fatal("Init kafka producer err:", err)
	}

	return &kafkaProducer{
		producer: producer,
		topic:    cfg.GetTopic(),
	}
}

func (p *kafkaProducer) Publish(ctx context.Context, order entity.Order) error {
	data, err := json.Marshal(order)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic:     p.topic,
		Partition: -1,
		Key:       sarama.StringEncoder(fmt.Sprintf("%d", order.UserId)),
		Value:     sarama.StringEncoder(data),
		Timestamp: time.Now(),
	}

	partition, offset, err := p.producer.SendMessage(msg)
	if err != nil {
		return err
	}
	log.Printf("message sent to partition %d at offset %d\n", partition, offset)
	return nil
}
