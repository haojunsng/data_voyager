package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"time"

	"weather/common"
)

func main() {
	producer, err := NewProducer()
	if err != nil {
		log.Fatalf("%s: %s", message, err)
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
				log.Fatalf("%s: %s", message, err)
			}

			message := fmt.Sprintf("Temperature: %.2f, WindSpeed: %.2f",
				weatherData.CurrentWeather.Temperature, weatherData.CurrentWeather.WindSpeed)
			err = producer.ProduceMessage(KafkaTopic, message)
			if err != nil {
				log.Fatalf("%s: %s", message, err)
			}
		case <-done:
			fmt.Println("Received interrupt signal, shutting down...")
			return
		}
	}
}
