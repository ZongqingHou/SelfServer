package utils

import (
	"time"

	"github.com/Shopify/sarama"
)

const ContextKafkaName = "kafka"

func KafkaConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second

	return config
}

func InitKafkaAsyncProducer(address []string) *sarama.AsyncProducer {
	config := KafkaConfig()
	producer, err := sarama.NewAsyncProducer(address, config)

	if err != nil {
		panic(err)
	}

	return &producer
}

func InitKafkaSyncProducer(address []string) *sarama.SyncProducer {
	config := KafkaConfig()
	producer, err := sarama.NewSyncProducer(address, config)

	if err != nil {
		panic(err)
	}

	return &producer
}

func InitKafkaConsumer(address []string) *sarama.Consumer {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer(address, config)
	if err != nil {
		panic(err)
	}

	return &consumer
}
