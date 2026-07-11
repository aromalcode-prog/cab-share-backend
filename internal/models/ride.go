package models

import "time"

type Ride struct {
	ID             uint
	DriverID       string
	Source         string
	Destination    string
	DepartureTime  string
	AvailableSeats string
	PricePerSeat   uint
	Status         bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
