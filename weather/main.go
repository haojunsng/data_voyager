package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	producer := createProducer()
	defer producer.Close()

	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	for {
		select {
		case <-ticker.C:
			param := createOpenMeteoParams()
			resp := fetchWeatherData(param)
			var weatherData WeatherData

			// Unmarshal JSON data into struct
			err := json.Unmarshal([]byte(resp), &weatherData)
			handleError(err, "Failed to unmarshal JSON into struct")

			message := fmt.Sprintf("Temperature: %.2f, WindSpeed: %.2f",
				weatherData.CurrentWeather.Temperature, weatherData.CurrentWeather.WindSpeed)
			produceMessage(producer, KafkaTopic, message)
		case <-done:
			fmt.Println("Received interrupt signal, shutting down...")
			return
		}
	}
}
