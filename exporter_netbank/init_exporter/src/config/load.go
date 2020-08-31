package config

import (
	"github.com/nats-io/nats.go"
	"os"
	"sync"
)

var (
	globalConfig *config
	configOnce   sync.Once
)

func Env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return fallback
}

func GetConfig() *config {
	if globalConfig == nil {
		configOnce.Do(func() {
			globalConfig = new(config)
		})
	}
	return globalConfig
}

func loadConfigFromEnv() {
	cfg := GetConfig()
	cfg.SetNatsURL(Env(envNatsURL, defNatsURL))
	cfg.SetPort(Env(envPort, defPort))
}

const (
	defNatsURL = nats.DefaultURL
	envNatsURL = "NATS_URL"

	defPort = "8180"
	envPort = "PORT"
)

func init() {
	loadConfigFromEnv()
}
