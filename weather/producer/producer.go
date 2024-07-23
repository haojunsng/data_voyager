package main

import (
	"context"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
	"weather/common"
)

type Producer struct {
	KafkaProducer *kafka.Producer
}

func NewProducer() (*Producer, error) {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":   "localhost:9092",
		"enable.idempotence":  true,
		"acks":                "all",
		"compression.type":    "snappy",
		"batch.num.messages":  1,
		"linger.ms":           10,
		"message.max.bytes":   1000000,
		"delivery.timeout.ms": 60000,
	}

	producer, err := kafka.NewProducer(configMap)
	if err != nil {
		return nil, err
	}
	return &Producer{KafkaProducer: producer}, nil
}

func (p *Producer) ProduceMessage(topic string, message string) error {
	deliveryChan := make(chan kafka.Event, 1)
	defer close(deliveryChan)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := p.KafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
		Headers:        []kafka.Header{{Key: "EventID", Value: []byte(uuid.New().String())}},
	}, deliveryChan)
	if err != nil {
		return err
	}
	select {
	case e := <-deliveryChan:
		m := e.(*kafka.Message)
		if m.TopicPartition.Error != nil {
			return m.TopicPartition.Error
		}
	case <-ctx.Done():
		// TODO : Handle context timeout if needed
	}
	return nil
}

func (p *Producer) Close() {
	p.KafkaProducer.Close()
}