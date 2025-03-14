package service

import (
	"auth-service/internal/model"
	"auth-service/internal/repository"
	"errors"
)

type AuthService struct {
	userRepository *repository.UserRepository
}

func NewAuthService(userRepository *repository.UserRepository) *AuthService {
	return &AuthService{userRepository: userRepository}
}

func (s *AuthService) Register(email, password, username string) (string, error) {
	existedEmail, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return "", nil
	}
	if existedEmail != nil {
		return "", errors.New(ErrEmailExists)
	}
	existedUsername, err := s.userRepository.FindByUsername(username)
	if err != nil {
		return "", nil
	}
	if existedUsername != nil {
		return "", errors.New(ErrUsernameExists)
	}
	user := &model.User{
		Email:    email,
		Password: password,
		Username: username,
	}
	s.userRepository.Create(user)
	return user.Email, nil
}
