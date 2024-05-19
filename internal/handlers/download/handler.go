package download_handler

import (
	"context"
	adapters "github.com/WildEgor/e-shop-cdn/internal/adapters/storage"
	"github.com/WildEgor/e-shop-cdn/internal/dtos"
	"github.com/WildEgor/e-shop-cdn/internal/utils"
	core_dtos "github.com/WildEgor/e-shop-gopack/pkg/core/dtos"
	"github.com/gofiber/fiber/v2"
)

type DownloadHandler struct {
	sa adapters.IFileStorage
}

func NewDownloadHandler(
	sa adapters.IFileStorage,
) *DownloadHandler {
	return &DownloadHandler{
		sa,
	}
}

// Handle Download file 				godoc
// @Summary Allow 						download file
// @Description	download file by name
// @Tags								Files Controller
// @Accept								json
// @Produce								json
// @Param								filename	path		string	true	"Filenam"
// @Router								/api/v1/cdn/download/{filename} [get]
func (hch *DownloadHandler) Handle(c *fiber.Ctx) error {
	resp := core_dtos.NewResp(core_dtos.WithOldContext(c))

	query := &dtos.FileQueryDto{
		Filename: c.Params("filename"),
	}
	if err := query.Validate(); err != nil {
		resp.SetStatus(fiber.StatusBadRequest)
		return resp.JSON()
	}

	filename, err := utils.ExtractFilename(query.Filename)
	if err != nil {
		resp.SetStatus(fiber.StatusBadRequest)
		return resp.JSON()
	}

	file, err := hch.sa.Download(context.Background(), query.Filename)
	if err != nil {
		// TODO: add real path to 404 file
		return err
	}

	resp.SetFile(filename, file)

	return resp.File()
}
