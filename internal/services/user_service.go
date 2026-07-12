package services

import (
	"github.com/aromalcode-prog/cab-share-backend/internal/dto"
	"github.com/aromalcode-prog/cab-share-backend/internal/repositories"
)

type UserService interface {
	Register(req dto.RegisterRequest) error
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) Register(req dto.RegisterRequest) error {
	return nil
}
