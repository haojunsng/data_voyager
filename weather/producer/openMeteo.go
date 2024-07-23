package main

import (
	"context"

	"github.com/innotechdevops/openmeteo"

	"weather/common"
)

func createOpenMeteoParams() openmeteo.Parameter {
	return openmeteo.Parameter{
		Latitude:  openmeteo.Float32(common.LatitudePunggol),
		Longitude: openmeteo.Float32(common.LatitudePunggol),
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
	if err != nil {
		log.Fatalf("%s: %s", message, err)
	}
	return resp
}

func NewWeather() *Weather {
	return &Weather{
		param: &openmeteo.Parameter{
			Latitude:  openmeteo.Float32(LatitudePunggol),
			Longitude: openmeteo.Float32(LatitudePunggol),
			Hourly: &[]string{
				openmeteo.HourlyTemperature2m,
				openmeteo.HourlyRelativeHumidity2m,
				openmeteo.HourlyWindSpeed10m,
			},
			CurrentWeather: openmeteo.Bool(true),
		},
	}
}

func (w *Weather) FetchData(ctx context.Context) string {
	m := openmeteo.New()
	resp, err := m.Execute(*w.param)
	if err != nil {
		log.Fatalf("%s: %s", message, err)
	}
	return resp
}

type IPipeline interface {
	FetchData(ctx context.Context) string
}
