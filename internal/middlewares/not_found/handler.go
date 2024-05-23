package not_found_middleware

import "github.com/gofiber/fiber/v2"

func NewNotFound() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Render("not_found", fiber.Map{
			"AppTitle": "App",
		})
	}
}
