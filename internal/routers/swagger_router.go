package routers

import (
	"github.com/gofiber/contrib/swagger"
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

	swaggerCfg := swagger.Config{
		BasePath: "/",
		Path:     "docs",
		FilePath: "./api/swagger/swagger.json",
	}

	app.Use(swagger.New(swaggerCfg))
}
