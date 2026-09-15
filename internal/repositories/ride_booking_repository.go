package repositories

import (
	"errors"

	"github.com/aromalcode-prog/cab-share-backend/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RideBookingRepository interface {
	Create(booking *models.RideBooking) error
	FindByBookingID(bookingID uint) (*models.RideBooking, error)
	FindByRideAndPassenger(rideID uint, passengerID uint) (*models.RideBooking, error)
	FindByPassengerID(passengerID uint) ([]models.RideBooking, error)
	Update(booking *models.RideBooking) error
}

type rideBookingRepository struct {
	db *gorm.DB
}

func NewRideBookingRepository(db *gorm.DB) RideBookingRepository {
	return &rideBookingRepository{db: db}
}

func (r *rideBookingRepository) Create(booking *models.RideBooking) error {
	result := r.db.Create(booking)
	return result.Error
}

func (r *rideBookingRepository) FindByBookingID(bookingID uint) (*models.RideBooking, error) {
	booking := &models.RideBooking{}
	result := r.db.
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("id = ?", bookingID).
		First(booking)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return booking, nil
}

func (r *rideBookingRepository) Update(booking *models.RideBooking) error {
	return r.db.Save(booking).Error
}

func (r *rideBookingRepository) FindByRideAndPassenger(rideID uint, passengerID uint) (*models.RideBooking, error) {
	rideBooking := &models.RideBooking{}

	result := r.db.Where("ride_id = ?", rideID).
		Where("passenger_id = ?", passengerID).
		Where("status = ?", models.BookingStatusActive).
		First(rideBooking)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return rideBooking, nil
}

func (r *rideBookingRepository) FindByPassengerID(passengerID uint) ([]models.RideBooking, error) {
	bookings := []models.RideBooking{}
	result := r.db.Where("passenger_id = ?", passengerID).
		Find(&bookings)
	return bookings, result.Error
}
