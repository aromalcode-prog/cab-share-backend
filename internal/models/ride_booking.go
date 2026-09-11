package models

import "time"

type RideBooking struct {
	ID          uint `gorm:"primaryKey"`
	RideID      uint `gorm:"not null;uniqueIndex:idx_ride_passenger"`
	PassengerID uint `gorm:"not null;uniqueIndex:idx_ride_passenger"`
	SeatsBooked uint `gorm:"not null;default:1"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
