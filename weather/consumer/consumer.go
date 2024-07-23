package main

import (
	"context"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	KafkaConsumer *kafka.Consumer
}

func NewConsumer() (*Consumer, error) {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "your-consumer-group",
		"auto.offset.reset": "earliest",
	}

	consumer, err := kafka.NewConsumer(configMap)
	if err != nil {
		log.Fatalf("Failed to create Kafka Consumer: %s", err)
		return nil, err
	}

	return &Consumer{KafkaConsumer: consumer}, nil
}

func (c *Consumer) Subscribe(topic string) error {
	err := c.KafkaConsumer.Subscribe(topic, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic %s: %s", topic, err)
		return err
	}
	return nil
}

func (c *Consumer) ConsumeMessages(ctx context.Context) {
	for {
		select {
		case msg := <-c.KafkaConsumer.Messages():
			log.Printf("Received message: %s", string(msg.Value))
			// Add processing logic here

		case <-ctx.Done():
			log.Println("Shutting down consumer")
			return
		}
	}
}

func (c *Consumer) Close() {
	c.KafkaConsumer.Close()
}
