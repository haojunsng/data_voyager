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
		log.Fatalf("Failed to create Kafka Producer: %s", err)
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

			// Unmarshal JSON data into struct
			err = json.Unmarshal([]byte(resp), &weatherData)
			if err != nil {
				log.Fatalf("Failed to unmarshal weather data: %s", err)
			}

			message := fmt.Sprintf("Temperature: %.2f, WindSpeed: %.2f",
				weatherData.CurrentWeather.Temperature, weatherData.CurrentWeather.WindSpeed)
			err = producer.ProduceMessage(common.KafkaTopic, message)
			if err != nil {
				log.Fatalf("Failed to produce Kafka message: %s", err)
			}
		case <-done:
			fmt.Println("Received interrupt signal, shutting down...")
			return
		}
	}
}
