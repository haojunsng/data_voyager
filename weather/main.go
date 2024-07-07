package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {
	producer, err := createProducer()
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}
	defer producer.Close()

	for {
		ticker := time.NewTicker(1 * time.Minute)
		defer ticker.Stop()

		<-ticker.C
		param := createOpenMeteoParams()
		resp := fetchWeatherData(param)
		var weatherData WeatherData

		// Unmarshal JSON data into struct
		err := json.Unmarshal([]byte(resp), &weatherData)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		message := fmt.Sprintf("Temperature: %.2f, WindSpeed: %.2f",
			weatherData.CurrentWeather.Temperature, weatherData.CurrentWeather.WindSpeed)

		if err := produceMessage(producer, KafkaTopic, message); err != nil {
			log.Printf("Failed to produce message: %v", err)
		}
	}
}
