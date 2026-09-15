package services

import (
	"errors"
	"strings"
	"time"

	"github.com/aromalcode-prog/cab-share-backend/internal/dto/request"
	"github.com/aromalcode-prog/cab-share-backend/internal/models"
	"github.com/aromalcode-prog/cab-share-backend/internal/repositories"
)

var (
	ErrSourceDestinationRequired = errors.New("source and destination are required")
	ErrDepartureTimeInPast       = errors.New("departure time must be in the future")
)

type RideService interface {
	CreateRide(req request.CreateRideRequestDTO, driverID uint) error
	GetAvailableRides() ([]models.Ride, error)
	SearchRides(req request.SearchRideRequestDTO) ([]models.Ride, error)
}

type rideService struct {
	rideRepo repositories.RideRepository
}

func NewRideService(rideRepo repositories.RideRepository) RideService {
	return &rideService{rideRepo: rideRepo}
}

func (s *rideService) CreateRide(req request.CreateRideRequestDTO, driverID uint) error {
	if !req.DepartureTime.After(time.Now()) {
		return ErrDepartureTimeInPast
	}

	ride := &models.Ride{
		Source:         req.Source,
		Destination:    req.Destination,
		DepartureTime:  req.DepartureTime,
		TotalSeats:     req.TotalSeats,
		AvailableSeats: req.TotalSeats,
		PricePerSeat:   req.PricePerSeat,
		DriverID:       driverID,
		Status:         models.RideStatusActive,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	return s.rideRepo.Create(ride)
}

func (s *rideService) GetAvailableRides() ([]models.Ride, error) {
	return s.rideRepo.AvailableRides()
}

func (s *rideService) SearchRides(req request.SearchRideRequestDTO) ([]models.Ride, error) {
	source := strings.TrimSpace(req.Source)
	destination := strings.TrimSpace(req.Destination)
	if source == "" || destination == "" {
		return nil, ErrSourceDestinationRequired
	}

	rides, err := s.rideRepo.Search(source, destination)
	if err != nil {
		return nil, err
	}
	if rides == nil {
		return []models.Ride{}, nil
	}
	return rides, nil
}
