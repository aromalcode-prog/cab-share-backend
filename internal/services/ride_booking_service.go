package services

import (
	"errors"
	"fmt"

	"github.com/aromalcode-prog/cab-share-backend/internal/models"
	"github.com/aromalcode-prog/cab-share-backend/internal/repositories"
	"gorm.io/gorm"
)

var (
	ErrBookingNotFound         = errors.New("booking not found")
	ErrBookingNotOwned         = errors.New("booking does not belong to the passenger")
	ErrBookingAlreadyCancelled = errors.New("booking has already been cancelled")
	ErrInvalidSeatCount        = errors.New("at least one seat must be booked")
	ErrRideNotFound            = errors.New("ride not found")
	ErrRideNotActive           = errors.New("ride is not active")
	ErrRideAlreadyBooked       = errors.New("passenger has already booked this ride")
	ErrNotEnoughSeats          = repositories.ErrNotEnoughSeats
)

type RideBookingService interface {
	BookRide(rideID uint, passengerID uint, seats uint) (*models.RideBooking, error)
	GetMyBookings(passengerID uint) ([]models.RideBooking, error)
	CancelBooking(bookingID uint, passengerID uint) (*models.RideBooking, error)
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
		return nil, ErrInvalidSeatCount
	}

	var booking *models.RideBooking

	err := s.db.Transaction(func(tx *gorm.DB) error {
		rideRepo := repositories.NewRideRepository(tx)
		rideBookingRepo := repositories.NewRideBookingRepository(tx)

		// Fetch the ride.
		ride, err := rideRepo.FindRideByID(rideID)
		if err != nil {
			return fmt.Errorf("failed to find ride: %w", err)
		}

		if ride == nil {
			return ErrRideNotFound
		}

		if ride.Status != models.RideStatusActive {
			return ErrRideNotActive
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
			return ErrRideAlreadyBooked
		}

		// Atomically reserve the seats.
		if err := rideRepo.ReserveSeats(rideID, seats); err != nil {
			if errors.Is(err, repositories.ErrNotEnoughSeats) {
				return ErrNotEnoughSeats
			}
			return fmt.Errorf("failed to reserve seats: %w", err)
		}

		// Create the booking only after seats are successfully reserved.
		booking = &models.RideBooking{
			RideID:      rideID,
			PassengerID: passengerID,
			SeatsBooked: seats,
			Status:      models.BookingStatusActive,
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

func (s *rideBookingService) CancelBooking(bookingID uint, passengerID uint) (*models.RideBooking, error) {
	var booking *models.RideBooking

	err := s.db.Transaction(func(tx *gorm.DB) error {
		rideRepo := repositories.NewRideRepository(tx)
		rideBookingRepo := repositories.NewRideBookingRepository(tx)

		found, err := rideBookingRepo.FindByBookingID(bookingID)
		if err != nil {
			return fmt.Errorf("failed to find booking: %w", err)
		}
		if found == nil {
			return ErrBookingNotFound
		}
		if found.PassengerID != passengerID {
			return ErrBookingNotOwned
		}
		if found.Status == models.BookingStatusCancelled {
			return ErrBookingAlreadyCancelled
		}

		if err := rideBookingRepo.MarkCancelled(found.ID, passengerID); err != nil {
			return fmt.Errorf("failed to cancel booking: %w", err)
		}

		if err := rideRepo.RestoreSeats(found.RideID, found.SeatsBooked); err != nil {
			return fmt.Errorf("failed to restore seats: %w", err)
		}

		found.Status = models.BookingStatusCancelled
		booking = found
		return nil
	})

	if err != nil {
		return nil, err
	}

	return booking, nil
}
