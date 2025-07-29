package handlers

import (
	"bytes"
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo/v4"

	"fitbyte/internal/models"
	"fitbyte/internal/services"
	"fitbyte/internal/utils"
)

type ValidateFileFunc func(fileHeader *multipart.FileHeader) error

type FileHandler struct {
	fileService  services.FileService
	validateFile ValidateFileFunc
}

func NewFileHanlder(fileService services.FileService, validateFile ValidateFileFunc) *FileHandler {
	return &FileHandler{fileService: fileService, validateFile: validateFile}
}

func (h *FileHandler) PostFile(ctx echo.Context) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return utils.SendErrorResponse(ctx, "File is required", http.StatusBadRequest)
	}

	if err := h.validateFile(file); err != nil {
		return utils.SendErrorResponse(ctx, err.Error(), http.StatusBadRequest)
	}

	src, err := file.Open()
	if err != nil {
		return utils.SendErrorResponse(ctx, "Failed to open file", http.StatusInternalServerError)
	}
	defer src.Close()

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(src); err != nil {
		return utils.SendErrorResponse(ctx, "Failed to read file content", http.StatusInternalServerError)
	}

	fileURL, err := h.fileService.UploadToS3(ctx, file.Filename, buf)
	if err != nil {
		return utils.SendErrorResponse(ctx, err.Error(), http.StatusInternalServerError)
	}

	response := models.FileResponse{
		URI: fileURL,
	}
	return ctx.JSON(http.StatusOK, response)
}
