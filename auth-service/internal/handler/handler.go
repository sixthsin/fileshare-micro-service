package handler

import (
	"auth-service/config"
	"auth-service/internal/service"
	"auth-service/pkg/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	MsgLoggedSuccessfully     = "User logged in successfully"
	MsgRegisteredSuccessfully = "User registered successfully"
)

type AuthHandlerDeps struct {
	*config.Config
	*service.AuthService
}

type AuthHandler struct {
	*config.Config
	*service.AuthService
}

func NewAuthHandler(router *gin.Engine, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.POST("/auth/register", handler.Register)
	router.POST("/auth/login", handler.Login)
}

func (handler *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	email, username, err := handler.AuthService.Register(req.Email, req.Password, req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(jwt.JWTData{
		Username: username,
		Email:    email,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, RegisterResponse{Token: token})
}

func (handler *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	email, username, err := handler.AuthService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(jwt.JWTData{
		Username: username,
		Email:    email,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	
}
