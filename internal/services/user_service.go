package services

import (
	"errors"

	"github.com/aromalcode-prog/cab-share-backend/internal/dto"
	"github.com/aromalcode-prog/cab-share-backend/internal/models"
	"github.com/aromalcode-prog/cab-share-backend/internal/repositories"
)

var ErrEmailAlreadyExists = errors.New("email already exists")

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
	_, err := s.userRepo.FindByEmail(req.Email)
	if err == nil {
		return ErrEmailAlreadyExists
	}

	if !errors.Is(err, repositories.ErrUserNotFound) {
		return err
	}

	user := &models.User{
		Email:        req.Email,
		Name:         req.Name,
		Phone:        req.Phone,
		PasswordHash: req.Password, //store after hashing
	}

	if err := s.userRepo.Create(user); err != nil {
		return err
	}
	return nil
}
