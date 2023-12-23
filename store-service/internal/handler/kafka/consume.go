package kafka

import (
	"context"
	"log"

	"gopkg.in/Shopify/sarama.v1"
)

const (
	groupID = "store_service"
)

type Config interface {
	GetBrokers() []string
	GetTopics() []string
}

func StartConsuming(ctx context.Context, cfg Config, handler sarama.ConsumerGroupHandler) error {
	config := sarama.NewConfig()
	config.Version = sarama.MaxVersion
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumerGroup, err := sarama.NewConsumerGroup(cfg.GetBrokers(), groupID, config)
	if err != nil {
		log.Fatal(err)
		return err
	}
	go func() {
		for {
			if err := consumerGroup.Consume(ctx, cfg.GetTopics(), handler); err != nil {
				log.Println("error from consumerGroup:", err)
			}

			if ctx.Err() != nil {
				return
			}
		}
	}()

	return nil
}
