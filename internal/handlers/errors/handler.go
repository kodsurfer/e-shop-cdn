package error_handler

import (
	"errors"
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

	sc := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		sc = e.Code
	}

	resp.SetStatus(sc)

	return resp.JSON()
}
