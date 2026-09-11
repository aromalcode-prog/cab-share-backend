package response

import (
	"time"

	"github.com/aromalcode-prog/cab-share-backend/internal/models"
)

type RideResponseDTO struct {
	ID             uint              `json:"id"`
	DriverID       uint              `json:"driver_id"`
	Source         string            `json:"source"`
	Destination    string            `json:"destination"`
	DepartureTime  time.Time         `json:"departure_time"`
	TotalSeats     uint              `json:"total_seats"`
	AvailableSeats uint              `json:"available_seats"`
	PricePerSeat   int64             `json:"price_per_seat"`
	Status         models.RideStatus `json:"status"`
}

func FromRideModel(ride models.Ride) RideResponseDTO {
	return RideResponseDTO{
		ID:             ride.ID,
		DriverID:       ride.DriverID,
		Source:         ride.Source,
		Destination:    ride.Destination,
		DepartureTime:  ride.DepartureTime,
		TotalSeats:     ride.TotalSeats,
		AvailableSeats: ride.AvailableSeats,
		PricePerSeat:   ride.PricePerSeat,
		Status:         ride.Status,
	}
}
