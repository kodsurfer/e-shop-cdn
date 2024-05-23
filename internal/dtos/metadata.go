package dtos

import (
	"github.com/gofiber/fiber/v2"
)

type FileMetadataResponseDto struct {
	Filename    string `json:"filename"`
	DownloadUrl string `json:"download_url"`
	FileSize    int64  `json:"file_size"`
}

func (fr *FileMetadataResponseDto) SetDownloadUrl(ctx *fiber.Ctx, filename string) {
	fr.DownloadUrl = string(ctx.BaseURL()) + "/api/v1/cdn/download/" + filename
}
