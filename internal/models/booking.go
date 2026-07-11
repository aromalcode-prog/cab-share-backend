package models

import "time"

type Booking struct {
	ID          uint
	RideID      uint
	PassengerID uint
	SeatsBooked uint
	Status      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
