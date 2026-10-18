package services

import (
	"backend/internal/dto"
	"backend/internal/models"
	"backend/internal/repository"
	"backend/internal/utils"
	"errors"
)

type AuthService struct {
	UserRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{UserRepo: userRepo}
}

func (s *AuthService) Register(input dto.RegisterRequest) (*models.User, error) {
	existingUser, _ := s.UserRepo.FindByEmail(input.Email)
	if existingUser != nil {
		return nil, errors.New("email already exists")
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
		Role:     "user",
	}

	err = s.UserRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
