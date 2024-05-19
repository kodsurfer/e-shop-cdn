package configs

import (
	"fmt"
	"github.com/WildEgor/e-shop-gopack/pkg/libs/logger/models"
	"github.com/caarlos0/env/v7"
	"log/slog"
	"runtime"
	"strings"
)

// MinioConfig s3 config
type MinioConfig struct {
	Endpoint  string `env:"MINIO_ENDPOINT,required"`
	AccessKey string `env:"MINIO_ACCESS_KEY,required"`
	Secret    string `env:"MINIO_SECRET_KEY,required"`
	Bucket    string `env:"MINIO_BUCKET,required"`
	Region    string `env:"MINIO_REGION" envDefault:"us-east-1"`
	UseSSL    bool   `env:"MINIO_SECURE" envDefault:"false"`
}

// StorageConfig storage factory config
type StorageConfig struct {
	Type  string `env:"STORAGE_TYPE" envDefault:"s3"`
	Minio MinioConfig

	ac *AppConfig
}

func NewStorageConfig(
	c *Configurator,
	ac *AppConfig,
) *StorageConfig {
	cfg := StorageConfig{}

	if err := env.Parse(&cfg); err != nil {
		slog.Error("fail parse storage config", models.LogEntryAttr(&models.LogEntry{
			Err: err,
		}))
	}

	cfg.ac = ac

	return &cfg
}

// DownloadUrl func create url to file by key
func (m *StorageConfig) DownloadUrl(key string) string {
	if m.ac.IsProduction() {
		return fmt.Sprintf("https://localhost:80/api/v1/cdn/download/%s", key)
	}

	if strings.EqualFold(runtime.GOOS, "windows") {
		return fmt.Sprintf("http://localhost:%s/api/v1/cdn/download/%s", m.ac.Port, key)
	}

	return fmt.Sprintf("http://127.0.0.1:%s/api/v1/cdn/download/%s", m.ac.Port, key)
}
