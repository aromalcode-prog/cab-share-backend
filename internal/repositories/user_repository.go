package repositories

import (
	"github.com/aromalcode-prog/cab-share-backend/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user *models.User) error {
	result := r.db.Create(user)
	return result.Error
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var foundUser models.User
	result := r.db.Where("email = ?", email).First(&foundUser)
	if result.Error != nil {
		return nil, result.Error
	}
	return &foundUser, nil
}
