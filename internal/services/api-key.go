package services

import (
	"github.com/WildEgor/e-shop-cdn/internal/configs"
	api_key_middleware "github.com/WildEgor/e-shop-gopack/pkg/core/middlewares/api_key_x"
	"github.com/WildEgor/e-shop-gopack/pkg/libs/logger/models"
	"log/slog"
	"strings"
)

type ApiKeyValidator struct {
	cfg *configs.ApiKeyConfig
}

func NewApiKeyValidator(
	cfg *configs.ApiKeyConfig,
) *ApiKeyValidator {
	return &ApiKeyValidator{
		cfg,
	}
}

func (v *ApiKeyValidator) Validate(key string) error {

	slog.Debug("Handle api key: %s", models.LogEntryAttr(&models.LogEntry{
		Props: map[string]interface{}{
			"key": key,
		},
	}))

	if !strings.EqualFold(v.cfg.Key, key) {
		return api_key_middleware.ErrWrongAPIKey
	}

	return nil
}
