package routers

import (
	dh "github.com/WildEgor/e-shop-cdn/internal/handlers/download"
	hch "github.com/WildEgor/e-shop-cdn/internal/handlers/health_check"
	rch "github.com/WildEgor/e-shop-cdn/internal/handlers/ready_check"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"log/slog"
)

type PublicRouter struct {
	hch *hch.HealthCheckHandler
	rch *rch.ReadyCheckHandler
	dwh *dh.DownloadHandler
}

func NewPublicRouter(
	hh *hch.HealthCheckHandler,
	rch *rch.ReadyCheckHandler,
	dwh *dh.DownloadHandler,
) *PublicRouter {
	return &PublicRouter{
		hh,
		rch,
		dwh,
	}
}

func (r *PublicRouter) Setup(app *fiber.App) {
	api := app.Group("/api", limiter.New(limiter.Config{
		Max:                    10,
		SkipSuccessfulRequests: true,
	}))
	v1 := api.Group("/v1")

	fc := v1.Group("/cdn")

	fc.Get("/download/:filename", r.dwh.Handle)

	app.Use(healthcheck.New(healthcheck.Config{
		LivenessProbe: func(c *fiber.Ctx) bool {
			if err := r.hch.Handle(c); err != nil {
				slog.Error("error not healthy")
				return false
			}

			slog.Debug("is healthy")

			return true
		},
		LivenessEndpoint: "/api/v1/livez",
		ReadinessProbe: func(c *fiber.Ctx) bool {
			if err := r.rch.Handle(c); err != nil {
				slog.Error("error not ready")
				return false
			}

			slog.Debug("is ready")

			return true
		},
		ReadinessEndpoint: "/api/v1/readyz",
	}))
}
