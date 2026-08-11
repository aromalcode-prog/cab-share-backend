package services

import (
	"errors"

	"github.com/aromalcode-prog/cab-share-backend/internal/auth"
	"github.com/aromalcode-prog/cab-share-backend/internal/dto/request"

	"github.com/aromalcode-prog/cab-share-backend/internal/models"
	"github.com/aromalcode-prog/cab-share-backend/internal/repositories"
)

var ErrEmailAlreadyExists = errors.New("email already exists")
var ErrInvalidCredentials = errors.New("invalid credentials")

type UserService interface {
	Register(req request.RegisterRequestDTO) error
	Login(req request.LoginRequestDTO) (string, error)
}

type userService struct {
	userRepo   repositories.UserRepository
	jwtManager *auth.JWTManager
}

func NewUserService(userRepo repositories.UserRepository, jwtManager *auth.JWTManager) UserService {
	return &userService{
		userRepo:   userRepo,
		jwtManager: jwtManager,
	}
}

func (s *userService) Register(req request.RegisterRequestDTO) error {
	_, err := s.userRepo.FindByEmail(req.Email)
	if err == nil {
		return ErrEmailAlreadyExists
	}

	if !errors.Is(err, repositories.ErrUserNotFound) {
		return err
	}
	hashedPassword, _ := auth.HashPassword(req.Password)
	user := &models.User{
		Email:        req.Email,
		Name:         req.Name,
		Phone:        req.Phone,
		PasswordHash: hashedPassword,
	}

	if err := s.userRepo.Create(user); err != nil {
		return err
	}
	return nil
}

func (s *userService) Login(req request.LoginRequestDTO) (string, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if errors.Is(err, repositories.ErrUserNotFound) {
		return "", ErrInvalidCredentials
	}
	if err != nil {
		return "", err
	}
	err = auth.ComparePassword(user.PasswordHash, req.Password)
	if err != nil {
		return "", ErrInvalidCredentials
	}
	return s.jwtManager.GenerateJWT(user.ID)
}
