package configs

import (
	"github.com/WildEgor/e-shop-gopack/pkg/libs/logger/models"
	"github.com/caarlos0/env/v7"
	"log/slog"
	"time"
)

// MongoConfig database config
type MongoConfig struct {
	URI               string        `env:"MONGODB_URI,required"`
	DbName            string        `env:"MONGODB_NAME,required"`
	ConnectionTimeout time.Duration `env:"-"`
	ct                int64         `env:"MONGODB_TIMEOUT" envDefault:"5"`
}

func NewMongoConfig(c *Configurator) *MongoConfig {
	cfg := MongoConfig{}

	if err := env.Parse(&cfg); err != nil {
		slog.Error("fail parse mongo config", models.LogEntryAttr(&models.LogEntry{
			Err: err,
		}))
	}

	cfg.ConnectionTimeout = time.Duration(cfg.ct) * time.Millisecond

	return &cfg
}
