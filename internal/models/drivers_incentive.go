package models

import (
	"time"
)

type DriverIncentive struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;unique" json:"id"`
	BookingID *uint      `json:"booking_id"`
	Incentive int       `json:"incentive"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Booking Booking `gorm:"foreignKey:BookingID" json:"booking,omitempty"`
}

type InputDriverIncentive struct {
	BookingID uint      `json:"booking_id" binding:"required"`
	Incentive int       `json:"incentive" binding:"required,gt=0"`
}