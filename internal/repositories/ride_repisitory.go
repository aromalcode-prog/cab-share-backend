package repositories

import (
	"errors"
	"time"

	"github.com/aromalcode-prog/cab-share-backend/internal/models"
	"gorm.io/gorm"
)

var ErrNotEnoughSeats = errors.New("not enough available seats or ride unavailable")

type RideRepository interface {
	Create(ride *models.Ride) error
	FindRideByID(rideID uint) (*models.Ride, error)
	AvailableRides() ([]models.Ride, error)
	ReserveSeats(rideID uint, seats uint) error
	RestoreSeats(rideID uint, seats uint) error
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

	result := r.db.Where("id = ?", rideID).
		First(ride)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return ride, nil
}

func (r *rideRepository) ReserveSeats(rideID uint, seats uint) error {
	result := r.db.
		Model(&models.Ride{}).
		Where("id = ?", rideID).
		Where("status = ?", models.RideStatusActive).
		Where("available_seats >= ?", seats).
		UpdateColumn(
			"available_seats",
			gorm.Expr("available_seats - ?", seats),
		)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrNotEnoughSeats
	}

	return nil
}

func (r *rideRepository) RestoreSeats(rideID uint, seats uint) error {
	result := r.db.
		Model(&models.Ride{}).
		Where("id = ?", rideID).
		Where("available_seats + ? <= total_seats", seats).
		UpdateColumn(
			"available_seats",
			gorm.Expr("available_seats + ?", seats),
		)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("failed to restore seats")
	}

	return nil
}
