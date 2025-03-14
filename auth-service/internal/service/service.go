package service

import (
	"auth-service/internal/model"
	"auth-service/internal/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user := &model.User{
		Email:    email,
		Password: string(hashedPassword),
		Username: username,
	}
	s.userRepository.Create(user)
	return user.Email, nil
}

func (s *AuthService) Login(email, password string) (string, error) {
	existedUser, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return "", nil
	}
	if existedUser == nil {
		return "", errors.New(ErrWrongCredentials)
	}
	err = bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(password))
	if err != nil {
		return "", errors.New(ErrWrongCredentials)
	}
	return existedUser.Email, nil
}
