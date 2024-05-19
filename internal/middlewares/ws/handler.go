package ws_middleware

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"log/slog"
)

func NewWS() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		slog.Debug("ws connection init")

		if websocket.IsWebSocketUpgrade(ctx) {
			ctx.Locals("allowed", true)
			return ctx.Next()
		}

		return ctx.SendStatus(fiber.StatusUpgradeRequired)
	}
}
