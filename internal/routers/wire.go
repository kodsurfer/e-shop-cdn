package routers

import (
	"github.com/WildEgor/e-shop-cdn/internal/handlers"
	"github.com/WildEgor/e-shop-cdn/internal/services"
	"github.com/google/wire"
)

// RouterSet acts like "controllers" for routing http or etc.
var RouterSet = wire.NewSet(
	handlers.HandlersSet,
	services.ServicesSet,
	NewPublicRouter,
	NewPrivateRouter,
	NewSwaggerRouter,
	NewSocketRouter,
)
