package main

import (
	"encoding/json"
	"fmt"
	"log"

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
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":  brokers,
		"message.timeout.ms": 20000,
	})
	handleError(err, "Failed to create producer")
	defer producer.Close()

	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          data,
	}, nil)
	handleError(err, "Failed to produce message")

	producer.Flush(30 * 1000)
	fmt.Println("Weather data produced to topic:", topic)
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
