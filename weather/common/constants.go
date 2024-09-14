package common

// Kafka configuration
const (
	KafkaTopic = "weather_topic"
	GroupID    = "weather_consumer"
)

// Geographical constants
const (
	LatitudePunggol  float32 = 1.3984
	LongitudePunggol float32 = 103.9072
)

// Weather data ranges
const (
	MinTemperature float32 = -100.0
	MaxTemperature float32 = 100.0
	MinHumidity    float32 = 0.0
	MaxHumidity    float32 = 100.0
	MinPressure    float32 = 900.0
	MaxPressure    float32 = 1100.0
	MinWindSpeed   float32 = 0.0
	MaxWindSpeed   float32 = 300.0
)

// AWS S3 configuration
const (
	S3Bucket  = "gomu-landing-bucket"
	AWSRegion = "ap-southeast-1"
)
