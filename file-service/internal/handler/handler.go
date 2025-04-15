package handler

import (
	"file-service/config"
	grpcauth "file-service/internal/grpc/auth"
	"file-service/internal/service"
	"file-service/pkg/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FileShareHandlerDeps struct {
	*config.Config
	*service.FileService
	*grpcauth.AuthClient
}

type FileShareHandler struct {
	*config.Config
	*service.FileService
	*grpcauth.AuthClient
}

func NewFileShareHandler(router *gin.Engine, deps FileShareHandlerDeps) {
	handler := &FileShareHandler{
		Config:      deps.Config,
		FileService: deps.FileService,
		AuthClient:  deps.AuthClient,
	}
	protected := router.Group("/file")
	protected.Use(middleware.AuthMiddleware(handler.AuthClient))
	{
		protected.GET("/download/{hash}", handler.Download)
		protected.POST("/upload", handler.Upload)
	}
}

func (h *FileShareHandler) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	openedFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	defer openedFile.Close()

	author := c.GetString("user")
	hash, err := h.FileService.SaveFile(openedFile, file.Filename, file.Size, file.Header.Get("Content-Type"), author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, FileResponse{
		Status:      "ok",
		Code:        http.StatusOK,
		Message:     "File uploaded successfully",
		Filename:    file.Filename,
		FileSize:    file.Size,
		ContentType: file.Header.Get("Content-Type"),
		Hash:        hash,
	})
}

func (h *FileShareHandler) Download(c *gin.Context) {

}
