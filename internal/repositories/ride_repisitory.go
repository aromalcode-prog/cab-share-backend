package repositories

import (
	"github.com/aromalcode-prog/cab-share-backend/internal/models"
	"gorm.io/gorm"
)

type RideRepository interface {
	Create(ride *models.Ride) error
}

type rideRepository struct {
	db *gorm.DB
}

func NewRideRepository(db *gorm.DB) RideRepository {
	return &rideRepository{db: db}
}

func (r *rideRepository) Create(ride *models.Ride) error {
	result := r.db.Create(ride)
	return result.Error
}
