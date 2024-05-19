package domains

import (
	"fmt"
	core_dtos "github.com/WildEgor/e-shop-gopack/pkg/core/dtos"
	"github.com/gofiber/fiber/v2"
)

const defErrCode = 99

var errCodesMessages = map[int]string{
	99:  "unknown error %s",
	100: "wrong api key",
	101: "missing api key",
	102: "fail serve file",
	103: "storage error",
	104: "file (%s) extension not allowed",
	105: "save file error",
}

func SetUnknown(r *core_dtos.ResponseDto, error error) {
	r.SetStatus(fiber.StatusInternalServerError)
	r.SetError(99, fmt.Sprintf(errCodesMessages[99], error.Error()))
}

func SetFileExtErr(r *core_dtos.ResponseDto, filename string) {
	r.SetStatus(fiber.StatusBadRequest)
	r.SetError(104, fmt.Sprintf(errCodesMessages[104], filename))
}

func SetFileServeErr(r *core_dtos.ResponseDto) {
	r.SetStatus(fiber.StatusInternalServerError)
	r.SetError(102, errCodesMessages[102])
}

func SetStorageErr(r *core_dtos.ResponseDto) {
	r.SetStatus(fiber.StatusInternalServerError)
	r.SetError(103, errCodesMessages[103])
}

func SetSaveFileErr(r *core_dtos.ResponseDto) {
	r.SetStatus(fiber.StatusInternalServerError)
	r.SetError(105, errCodesMessages[105])
}
