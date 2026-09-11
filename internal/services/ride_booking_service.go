package services

import (
	"errors"
	"fmt"

	"github.com/aromalcode-prog/cab-share-backend/internal/models"
	"github.com/aromalcode-prog/cab-share-backend/internal/repositories"
	"gorm.io/gorm"
)

type RideBookingService interface {
	BookRide(rideID uint, passengerID uint, seats uint) (*models.RideBooking, error)
	GetMyBookings(passengerID uint) ([]models.RideBooking, error)
}

type rideBookingService struct {
	db              *gorm.DB
	rideRepo        repositories.RideRepository
	rideBookingRepo repositories.RideBookingRepository
}

func NewRideBookingService(db *gorm.DB, rideRepo repositories.RideRepository, rideBookingRepo repositories.RideBookingRepository) RideBookingService {
	return &rideBookingService{
		db:              db,
		rideRepo:        rideRepo,
		rideBookingRepo: rideBookingRepo,
	}
}

func (s *rideBookingService) BookRide(rideID uint, passengerID uint, seats uint) (*models.RideBooking, error) {
	if seats == 0 {
		return nil, errors.New("at least one seat must be booked")
	}

	var booking *models.RideBooking

	err := s.db.Transaction(func(tx *gorm.DB) error {
		// Use repositories connected to this transaction.
		rideRepo := repositories.NewRideRepository(tx)
		rideBookingRepo := repositories.NewRideBookingRepository(tx)

		// Fetch the ride.
		ride, err := rideRepo.FindRideByID(rideID)
		if err != nil {
			return fmt.Errorf("failed to find ride: %w", err)
		}

		if ride == nil {
			return errors.New("ride not found")
		}

		if ride.Status != models.RideStatusActive {
			return errors.New("ride is not active")
		}

		// Check whether this passenger has already booked the ride.
		existingBooking, err := rideBookingRepo.FindByRideAndPassenger(
			rideID,
			passengerID,
		)
		if err != nil {
			return fmt.Errorf("failed to check existing booking: %w", err)
		}

		if existingBooking != nil {
			return errors.New("passenger has already booked this ride")
		}

		// Atomically reserve the seats.
		if err := rideRepo.ReserveSeats(rideID, seats); err != nil {
			return err
		}

		// Create the booking only after seats are successfully reserved.
		booking = &models.RideBooking{
			RideID:      rideID,
			PassengerID: passengerID,
			SeatsBooked: seats,
		}

		if err := rideBookingRepo.Create(booking); err != nil {
			return fmt.Errorf("failed to create booking: %w", err)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return booking, nil
}

func (s *rideBookingService) GetMyBookings(passengerID uint) ([]models.RideBooking, error) {
	return s.rideBookingRepo.FindByPassengerID(passengerID)
}
