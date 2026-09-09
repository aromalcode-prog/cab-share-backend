package models

import "time"

type RideStatus string

const (
	RideStatusActive    RideStatus = "active"
	RideStatusCancelled RideStatus = "cancelled"
	RideStatusCompleted RideStatus = "completed"
)

type Ride struct {
	ID             uint
	DriverID       uint
	Source         string
	Destination    string
	DepartureTime  time.Time
	TotalSeats     uint
	AvailableSeats uint
	PricePerSeat   int64
	Status         RideStatus
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
