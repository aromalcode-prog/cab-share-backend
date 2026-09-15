package repositories

import (
	"errors"
	"time"

	"github.com/aromalcode-prog/cab-share-backend/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RideRepository interface {
	Create(ride *models.Ride) error
	FindRideByID(rideID uint) (*models.Ride, error)
	AvailableRides() ([]models.Ride, error)
	Update(ride *models.Ride) error
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

func (r *rideRepository) FindRideByID(rideID uint) (*models.Ride, error) {
	ride := &models.Ride{}

	result := r.db.
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("id = ?", rideID).
		First(ride)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return ride, nil
}

func (r *rideRepository) Update(ride *models.Ride) error {
	result := r.db.Save(ride)
	return result.Error
}
