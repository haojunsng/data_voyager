package main

type WeatherData struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	GenerationTimeMs     float64 `json:"generationtime_ms"`
	UTCOffsetSeconds     int     `json:"utc_offset_seconds"`
	Timezone             string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	Elevation            float64 `json:"elevation"`
	CurrentWeatherUnits  struct {
		Time          string `json:"time"`
		Interval      string `json:"interval"`
		Temperature   string `json:"temperature"`
		WindSpeed     string `json:"windspeed"`
		WindDirection string `json:"winddirection"`
		IsDay         string `json:"is_day"`
		WeatherCode   string `json:"weathercode"`
	} `json:"current_weather_units"`
	CurrentWeather struct {
		Time          string  `json:"time"`
		Interval      int     `json:"interval"`
		Temperature   float64 `json:"temperature"`
		WindSpeed     float64 `json:"windspeed"`
		WindDirection int     `json:"winddirection"`
		IsDay         int     `json:"is_day"`
		WeatherCode   int     `json:"weathercode"`
	} `json:"current_weather"`
	HourlyUnits struct {
		Time               string `json:"time"`
		Temperature2m      string `json:"temperature_2m"`
		RelativeHumidity2m string `json:"relativehumidity_2m"`
		WindSpeed10m       string `json:"windspeed_10m"`
	} `json:"hourly_units"`
	Hourly struct {
		Time               []string  `json:"time"`
		Temperature2m      []float64 `json:"temperature_2m"`
		RelativeHumidity2m []int     `json:"relativehumidity_2m"`
		WindSpeed10m       []float64 `json:"windspeed_10m"`
	} `json:"hourly"`
}
