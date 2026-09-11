package repositories

import (
	"time"

	"github.com/aromalcode-prog/cab-share-backend/internal/models"
	"gorm.io/gorm"
)

type RideRepository interface {
	Create(ride *models.Ride) error
	AvailableRides() ([]models.Ride, error)
}

type rideRepository struct {
	db *gorm.DB
}

func NewRideRepository(db *gorm.DB) RideRepository {
	return &rideRepository{db: db}
}

func (r *rideRepository) AvailableRides() ([]models.Ride, error) {
	var rides []models.Ride
	result := r.db.
		Where("status = ?", models.RideStatusActive).
		Where("available_seats > ?", 0).
		Where("departure_time > ?", time.Now()).
		Order("departure_time ASC").
		Find(&rides)

	return rides, result.Error
}

func (r *rideRepository) Create(ride *models.Ride) error {
	result := r.db.Create(ride)
	return result.Error
}
