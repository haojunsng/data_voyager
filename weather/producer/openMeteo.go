package main

import (
	"weather/common"

	"github.com/innotechdevops/openmeteo"
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
	common.HandleError(err, "Failed to execute API call")
	return resp
}
