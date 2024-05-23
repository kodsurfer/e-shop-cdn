package error_handler

import (
	domains "github.com/WildEgor/e-shop-cdn/internal/domain"
	core_dtos "github.com/WildEgor/e-shop-gopack/pkg/core/dtos"
	"github.com/gofiber/fiber/v2"
)

// ErrorsHandler acts like global error handler
type ErrorsHandler struct {
}

func NewErrorsHandler() *ErrorsHandler {
	return &ErrorsHandler{}
}

func (hch *ErrorsHandler) Handle(ctx *fiber.Ctx, err error) error {
	resp := core_dtos.NewResp(core_dtos.WithOldContext(ctx))

	domains.SetUnknown(resp, err)

	return resp.JSON()
}
