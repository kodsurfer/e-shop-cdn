package routers

import (
	_ "github.com/WildEgor/e-shop-cdn/api/swagger"
	"github.com/gofiber/fiber/v2"
)

type SwaggerRouter struct {
}

func NewSwaggerRouter() *SwaggerRouter {
	return &SwaggerRouter{}
}

// Setup func for describe group of API Docs routes.
func (sr *SwaggerRouter) Setup(app *fiber.App) {
	// TODO: fiber v3 not impl swagger middleware now
}
