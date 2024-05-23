package configs

import (
	"github.com/WildEgor/e-shop-gopack/pkg/libs/logger/models"
	"github.com/caarlos0/env/v7"
	"log/slog"
)

// ApiKeyConfig for custom api key from headers
type ApiKeyConfig struct {
	Key string `env:"API_KEY,required"`
}

func NewApiKeyConfig(c *Configurator) *ApiKeyConfig {
	cfg := ApiKeyConfig{}

	if err := env.Parse(&cfg); err != nil {
		slog.Error("fail parse api key config", models.LogEntryAttr(&models.LogEntry{
			Err: err,
		}))
	}

	return &cfg
}
