package models

import (
	"time"
)

type BookingType struct {
    ID       uint   `gorm:"primaryKey;autoIncrement;unique" json:"id"`
    BookingType     string `json:"booking_type"`
    Description     string  `json:"description"`
    CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type InputBookingType struct {
    BookingType     string `json:"booking_type" binding:"required"`
    Description     string  `json:"description" binding:"required"`
}