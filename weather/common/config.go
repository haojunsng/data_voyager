package common

import (
	"context"
	"log"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
)

type Config struct {
	KafkaBroker string `config:"kafka_broker"`
}

var (
	KafkaBroker = "localhost:9092" // default value to be overridden in production using kafka_broker
)

func init() {
	loader := confita.NewLoader(env.NewBackend())

	cfg := Config{}

	err := loader.Load(context.Background(), &cfg)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	if cfg.KafkaBroker != "" {
		KafkaBroker = cfg.KafkaBroker
	}
}
