package service

import (
	"auth-service/config"
	"auth-service/internal/model"
	"auth-service/internal/repository"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceDeps struct {
	*config.Config
	*repository.UserRepository
}

type AuthService struct {
	*config.Config
	*repository.UserRepository
}

func NewAuthService(deps AuthServiceDeps) *AuthService {
	return &AuthService{
		Config:         deps.Config,
		UserRepository: deps.UserRepository,
	}
}

func (s *AuthService) Register(email, password, username string) (string, string, error) {
	existedEmail, err := s.UserRepository.FindByEmail(email)
	if err != nil {
		return "", "", nil
	}
	if existedEmail != nil {
		return "", "", errors.New(ErrEmailExists)
	}

	existedUsername, err := s.UserRepository.FindByUsername(username)
	if err != nil {
		return "", "", nil
	}
	if existedUsername != nil {
		return "", "", errors.New(ErrUsernameExists)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}

	user := &model.User{
		Email:    email,
		Password: string(hashedPassword),
		Username: username,
	}

	s.UserRepository.Create(user)

	return user.Email, user.Username, nil
}

func (s *AuthService) Login(email, password string) (string, string, error) {
	existedUser, err := s.UserRepository.FindByEmail(email)
	if err != nil {
		return "", "", nil
	}
	if existedUser == nil {
		return "", "", errors.New(ErrWrongCredentials)
	}

	err = bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(password))
	if err != nil {
		return "", "", errors.New(ErrWrongCredentials)
	}

	return existedUser.Email, existedUser.Username, nil
}

func (s *AuthService) ValidateToken(tokenString string) (*UserResponse, error) {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Config.Auth.Secret), nil
	})
	if err != nil || !token.Valid {
		return &UserResponse{IsValid: false}, errors.New(ErrInvalidToken)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return &UserResponse{IsValid: false}, errors.New(ErrInvalidTokenClaims)
	}

	expirationTime := int64(claims["exp"].(float64))
	if time.Now().Unix() > expirationTime {
		return &UserResponse{IsValid: false}, errors.New(ErrTokenExpired)
	}

	email := claims["email"].(string)
	username := claims["username"].(string)
	return &UserResponse{
		Username: username,
		Email:    email,
		IsValid:  true,
	}, nil
}
