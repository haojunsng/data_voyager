package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
)

func createProducer() (*kafka.Producer, error) {
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

	return producer, nil
}

func produceMessage(producer *kafka.Producer, topic string, message string) error {
	deliveryChan := make(chan kafka.Event, 1)

	err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
		Headers:        []kafka.Header{{Key: "EventID", Value: []byte(uuid.New().String())}},
	}, deliveryChan)

	if err != nil {
		return err
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		return m.TopicPartition.Error
	}

	return nil
}
