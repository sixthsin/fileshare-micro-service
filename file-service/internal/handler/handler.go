package handler

import (
	"file-processor-service/config"

	"github.com/gin-gonic/gin"
)

type FileShareHandlerDeps struct {
	*config.Config
}

type FileShareHandler struct {
	*config.Config
}

func NewFileShareHandler(router *gin.Engine, deps FileShareHandlerDeps) {
	handler := &FileShareHandler{
		Config: deps.Config,
	}
	router.GET("/file/download/{hash}", handler.Download)
	router.POST("/auth/upload", handler.Upload)
}

func (h *FileShareHandler) Upload(c *gin.Context) {

}

func (h *FileShareHandler) Download(c *gin.Context) {

}
