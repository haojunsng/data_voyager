package common

type WeatherData struct {
	Latitude             float64 `json:"latitude" parquet:"type=DOUBLE"`
	Longitude            float64 `json:"longitude" parquet:"type=DOUBLE"`
	GenerationTimeMs     float64 `json:"generationtime_ms" parquet:"type=DOUBLE"`
	UTCOffsetSeconds     int     `json:"utc_offset_seconds" parquet:"type=INT32"`
	Timezone             string  `json:"timezone" parquet:"type=BYTE_ARRAY,encoding=PLAIN"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation" parquet:"type=BYTE_ARRAY,encoding=PLAIN"`
	Elevation            float64 `json:"elevation" parquet:"type=DOUBLE"`
	CurrentWeatherUnits  struct {
		Time          string `json:"time" parquet:"type=BYTE_ARRAY,encoding=PLAIN"`
		Interval      string `json:"interval" parquet:"type=BYTE_ARRAY,encoding=PLAIN"`
		Temperature   string `json:"temperature" parquet:"type=BYTE_ARRAY,encoding=PLAIN"`
		WindSpeed     string `json:"windspeed" parquet:"type=BYTE_ARRAY,encoding=PLAIN"`
		WindDirection string `json:"winddirection" parquet:"type=BYTE_ARRAY,encoding=PLAIN"`
		IsDay         string `json:"is_day" parquet:"type=BYTE_ARRAY,encoding=PLAIN"`
		WeatherCode   string `json:"weathercode" parquet:"type=BYTE_ARRAY,encoding=PLAIN"`
	} `json:"current_weather_units"`
	CurrentWeather struct {
		Time          string  `json:"time" parquet:"type=BYTE_ARRAY,encoding=PLAIN"`
		Interval      int     `json:"interval" parquet:"type=INT32"`
		Temperature   float64 `json:"temperature" parquet:"type=DOUBLE"`
		WindSpeed     float64 `json:"windspeed" parquet:"type=DOUBLE"`
		WindDirection int     `json:"winddirection" parquet:"type=INT32"`
		IsDay         int     `json:"is_day" parquet:"type=INT32"`
		WeatherCode   int     `json:"weathercode" parquet:"type=INT32"`
	} `json:"current_weather"`
	HourlyUnits struct {
		Time               string `json:"time" parquet:"type=BYTE_ARRAY,encoding=PLAIN"`
		Temperature2m      string `json:"temperature_2m" parquet:"type=BYTE_ARRAY,encoding=PLAIN"`
		RelativeHumidity2m string `json:"relativehumidity_2m" parquet:"type=BYTE_ARRAY,encoding=PLAIN"`
		WindSpeed10m       string `json:"windspeed_10m" parquet:"type=BYTE_ARRAY,encoding=PLAIN"`
	} `json:"hourly_units"`
	Hourly struct {
		Time               []string  `json:"time" parquet:"type=LIST,element_type=BYTE_ARRAY,encoding=PLAIN"`
		Temperature2m      []float64 `json:"temperature_2m" parquet:"type=LIST,element_type=DOUBLE"`
		RelativeHumidity2m []int     `json:"relativehumidity_2m" parquet:"type=LIST,element_type=INT32"`
		WindSpeed10m       []float64 `json:"windspeed_10m" parquet:"type=LIST,element_type=DOUBLE"`
	} `json:"hourly"`
}
