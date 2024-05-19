package delete_handler

import (
	"github.com/WildEgor/e-shop-cdn/internal/adapters/pubsub"
	adapters "github.com/WildEgor/e-shop-cdn/internal/adapters/storage"
	"github.com/WildEgor/e-shop-cdn/internal/repositories"
	core_dtos "github.com/WildEgor/e-shop-gopack/pkg/core/dtos"
	"github.com/gofiber/fiber/v2"
)

type DeleteHandler struct {
	fr     repositories.IFilesRepository
	sa     adapters.IFileStorage
	pubsub pubsub.IPubSub
}

func NewDeleteHandler(
	fr repositories.IFilesRepository,
	sa adapters.IFileStorage,
	pubsub pubsub.IPubSub,
) *DeleteHandler {
	return &DeleteHandler{
		fr,
		sa,
		pubsub,
	}
}

// Delete file 		godoc
// @Summary 		Allow delete file
// @Description		delete file by name
// @Tags			Files Controller
// @Accept			json
// @Produce			json
// @Param			x-api-key header	string	true	"123"
// @Param			filename	path		string	true	"Filenam"
// @Router			/api/v1/cdn/delete/{id} [delete]
func (hch *DeleteHandler) Handle(c *fiber.Ctx) error {
	resp := core_dtos.NewResp(core_dtos.WithOldContext(c))
	resp.SetStatus(fiber.StatusOK)

	id := c.Params("id")

	oldFile, err := hch.fr.DeleteFileById(id)
	if err != nil {
		resp.SetStatus(fiber.StatusInternalServerError)
	}

	// TODO: remove file in another cron job using is_deleted field (prevent s3 call with mongo)
	if oldFile != nil {
		hch.pubsub.Publish(oldFile.DirPrefix(), "DELETED")

		if err := hch.sa.Delete(oldFile.Name); err != nil {
			resp.SetStatus(fiber.StatusInternalServerError)
		}
	}

	return resp.JSON()
}
