package services

import (
	"github.com/aromalcode-prog/cab-share-backend/internal/dto/request"
	"github.com/aromalcode-prog/cab-share-backend/internal/models"
	"github.com/aromalcode-prog/cab-share-backend/internal/repositories"
)

type RideService interface {
	CreateRide(req request.CreateRideRequestDTO, driverID uint) error
}

type rideService struct {
	rideRepo repositories.RideRepository
}

func NewRideService(rideRepo repositories.RideRepository) RideService {
	return &rideService{rideRepo: rideRepo}
}

func (s *rideService) CreateRide(req request.CreateRideRequestDTO, driverID uint) error {
	ride := &models.Ride{
		Source:        req.Source,
		Destination:   req.Destination,
		DepartureTime: req.DepartureTime,
		TotalSeats:    req.TotalSeats,
		PricePerSeat:  req.PricePerSeat,
		DriverID:      driverID,
		Status:        models.RideStatusActive,
	}
	return s.rideRepo.Create(ride)
}
