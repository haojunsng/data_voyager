package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"weather/common"
)

func main() {
	producer, err := NewProducer()
	if err != nil {
		log.Fatalf("failed to create Kafka Producer: %s", err)
	}
	defer producer.Close()

	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	ctx := context.Background()
	pipeline := NewWeather()

	for {
		select {
		case <-ticker.C:
			resp := pipeline.FetchData(ctx)

			var weatherData common.WeatherData

			// Unmarshal JSON data into struct to quickly validate data
			err = json.Unmarshal([]byte(resp), &weatherData)
			if err != nil {
				log.Fatalf("failed to unmarshal weather data: %s", err)
			}

			message, err := json.Marshal(weatherData)
			if err != nil {
				log.Fatalf("failed to marshal weather data to JSON: %s", err)
			}

			err = producer.ProduceMessage(common.KafkaTopic, string(message))
			if err != nil {
				log.Fatalf("failed to produce Kafka message: %s", err)
			}
		case <-done:
			fmt.Println("received interrupt signal, shutting down...")
			return
		}
	}
}
