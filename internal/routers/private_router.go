package routers

import (
	delh "github.com/WildEgor/e-shop-cdn/internal/handlers/delete"
	gf "github.com/WildEgor/e-shop-cdn/internal/handlers/get_files"
	mh "github.com/WildEgor/e-shop-cdn/internal/handlers/metadata"
	uh "github.com/WildEgor/e-shop-cdn/internal/handlers/upload"
	"github.com/WildEgor/e-shop-cdn/internal/services"
	api_key_middleware "github.com/WildEgor/e-shop-gopack/pkg/core/middlewares/api_key_x"
	"github.com/gofiber/fiber/v2"
)

type PrivateRouter struct {
	uh *uh.UploadHandler
	dh *delh.DeleteHandler
	gf *gf.GetFilesHandler
	mh *mh.MetadataHandler

	vs *services.ApiKeyValidator
}

func NewPrivateRouter(
	uh *uh.UploadHandler,
	dh *delh.DeleteHandler,
	gf *gf.GetFilesHandler,
	mh *mh.MetadataHandler,
	vs *services.ApiKeyValidator,
) *PrivateRouter {
	return &PrivateRouter{
		uh,
		dh,
		gf,
		mh,
		vs,
	}
}

func (r *PrivateRouter) Setup(app *fiber.App) {
	v1 := app.Group("/api/v1")

	akm := api_key_middleware.NewApiKeyMiddleware(api_key_middleware.ApiKeyMiddlewareConfig{
		KeyLookup: "header:x-api-key",
		Validator: func(ctx *fiber.Ctx, key string) (bool, error) {
			if err := r.vs.Validate(key); err != nil {
				return false, err
			}

			return true, nil
		},
	})

	fc := v1.Group("/cdn")

	fc.Post("/upload", akm, r.uh.Handle)
	fc.Post("/metadata/:filename", akm, r.mh.Handle)
	fc.Post("/files", akm, r.gf.Handle)
	fc.Delete("/delete/:id", akm, r.dh.Handle)
}
