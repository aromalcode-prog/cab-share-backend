package request

import "time"

type CreateRideRequestDTO struct {
	Source        string    `json:"source" binding:"required"`
	Destination   string    `json:"destination" binding:"required"`
	DepartureTime time.Time `json:"departure_time" binding:"required"`
	TotalSeats    uint      `json:"total_seats" binding:"required,min=1"`
	PricePerSeat  int64     `json:"price_per_seat" binding:"required,min=0"`
}

type SearchRideRequestDTO struct {
	Source      string `form:"source" binding:"required"`
	Destination string `form:"destination" binding:"required"`
}
