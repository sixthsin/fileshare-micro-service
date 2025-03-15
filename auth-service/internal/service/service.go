package service

import (
	"auth-service/config"
	"auth-service/internal/model"
	"auth-service/internal/repository"
	"errors"
	"strconv"
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

func (s *AuthService) Register(email, password, username string) (string, uint, error) {
	existedEmail, err := s.UserRepository.FindByEmail(email)
	if err != nil {
		return "", 0, nil
	}
	if existedEmail != nil {
		return "", 0, errors.New(ErrEmailExists)
	}

	existedUsername, err := s.UserRepository.FindByUsername(username)
	if err != nil {
		return "", 0, nil
	}
	if existedUsername != nil {
		return "", 0, errors.New(ErrUsernameExists)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", 0, err
	}

	user := &model.User{
		Email:    email,
		Password: string(hashedPassword),
		Username: username,
	}

	s.UserRepository.Create(user)

	return user.Email, user.ID, nil
}

func (s *AuthService) Login(email, password string) (string, uint, error) {
	existedUser, err := s.UserRepository.FindByEmail(email)
	if err != nil {
		return "", 0, nil
	}
	if existedUser == nil {
		return "", 0, errors.New(ErrWrongCredentials)
	}

	err = bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(password))
	if err != nil {
		return "", 0, errors.New(ErrWrongCredentials)
	}

	return existedUser.Email, existedUser.ID, nil
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

	idValue, ok := claims["user_id"]
	if !ok {
		return &UserResponse{IsValid: false}, errors.New(ErrInvalidTokenClaims)
	}

	var id string
	switch v := idValue.(type) {
	case float64:
		id = strconv.FormatFloat(v, 'f', -1, 64)
	case string:
		id = v
	default:
		return &UserResponse{IsValid: false}, errors.New(ErrInvalidTokenClaims)
	}

	email := claims["email"].(string)

	return &UserResponse{
		ID:      id,
		Email:   email,
		IsValid: true,
	}, nil
}
