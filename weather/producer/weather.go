package main

import (
	"context"
	"log"

	"weather/common"

	"github.com/innotechdevops/openmeteo"
)

type Weather struct {
	param *openmeteo.Parameter
}

func NewWeather() *Weather {
	return &Weather{
		param: &openmeteo.Parameter{
			Latitude:  openmeteo.Float32(common.LatitudePunggol),
			Longitude: openmeteo.Float32(common.LatitudePunggol),
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

	respChan := make(chan string)
	errChan := make(chan error)

	go func() {
		resp, err := m.Execute(*w.param)
		if err != nil {
			errChan <- err
			return
		}
		respChan <- resp
	}()

	select {
	case resp := <-respChan:
		return resp
	case err := <-errChan:
		log.Fatalf("Failed to fetch weather data: %s", err)
	case <-ctx.Done():
		log.Fatalf("FetchData timed out or was canceled: %s", ctx.Err())
	}

	return ""
}
