package main

import (
	"weather/common"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
)

func createProducer() *kafka.Producer {
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
	common.HandleError(err, "Failed to create Kafka Producer")

	return producer
}

func produceMessage(producer *kafka.Producer, topic string, message string) error {
	deliveryChan := make(chan kafka.Event, 1)

	err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
		Headers:        []kafka.Header{{Key: "EventID", Value: []byte(uuid.New().String())}},
	}, deliveryChan)
	common.HandleError(err, "Failed to produce Kafka Message")

	e := <-deliveryChan
	m := e.(*kafka.Message)
	common.HandleError(m.TopicPartition.Error, "Failed to deliver message")

	return nil
}
