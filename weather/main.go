package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/innotechdevops/openmeteo"
)

func handleError(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %s", message, err)
	}
}

func createOpenMeteoParams() openmeteo.Parameter {
	return openmeteo.Parameter{
		Latitude:  openmeteo.Float32(LatitudePunggol),
		Longitude: openmeteo.Float32(LatitudePunggol),
		Hourly: &[]string{
			openmeteo.HourlyTemperature2m,
			openmeteo.HourlyRelativeHumidity2m,
			openmeteo.HourlyWindSpeed10m,
		},
		CurrentWeather: openmeteo.Bool(true),
	}
}

func fetchWeatherData(param openmeteo.Parameter) string {
	m := openmeteo.New()
	resp, err := m.Execute(param)
	handleError(err, "Failed to execute API call")
	return resp
}

func produceToKafka(topic string, data []byte, brokers string) {
	// Create a buffered channel to receive delivery events
	deliveryChan := make(chan kafka.Event, 1)
	defer close(deliveryChan) // Ensure channel is closed to avoid leaks

	// Create Kafka producer instance
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":  brokers,
		"message.timeout.ms": 20000,
	})
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer producer.Close() // Ensure producer is closed at the end of function

	// Produce message asynchronously
	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          data,
	}, deliveryChan)
	if err != nil {
		log.Fatalf("Failed to produce message: %v", err)
	}

	// Wait for delivery report
	select {
	case e := <-deliveryChan:
		m := e.(*kafka.Message)

		// Check delivery status
		if m.TopicPartition.Error != nil {
			log.Fatalf("Delivery failed: %v", m.TopicPartition.Error)
		} else {
			fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
				*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		}

	case <-time.After(20 * time.Second): // Adjust timeout as needed
		log.Fatalf("Delivery timed out after 20 seconds")
	}

	// Flush and wait for all messages to be delivered
	queueLength := producer.Flush(30 * 1000)
	if queueLength > 0 {
		log.Printf("Failed to deliver %d messages", queueLength)
	} else {
		fmt.Println("Weather data produced to topic:", topic)
	}
}

func main() {
	param := createOpenMeteoParams()

	resp := fetchWeatherData(param)

	respJSON, err := json.Marshal(resp)
	handleError(err, "Failed to marshal response")

	kafkaTopic := "weather-data"
	kafkaBrokers := "localhost:9092"

	produceToKafka(kafkaTopic, respJSON, kafkaBrokers)
}
