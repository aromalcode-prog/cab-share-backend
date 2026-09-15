package models

import "time"

type RideStatus string

const (
	RideStatusActive    RideStatus = "active"
	RideStatusCancelled RideStatus = "cancelled"
	RideStatusCompleted RideStatus = "completed"
)

type Ride struct {
	ID             uint       `gorm:"primaryKey"`
	DriverID       uint       `gorm:"not null"`
	Source         string     `gorm:"not null"`
	Destination    string     `gorm:"not null"`
	DepartureTime  time.Time  `gorm:"not null"`
	TotalSeats     uint       `gorm:"not null"`
	AvailableSeats uint       `gorm:"not null"`
	PricePerSeat   int64      `gorm:"not null"`
	Status         RideStatus `gorm:"not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
