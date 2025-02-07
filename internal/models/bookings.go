package models

import (
	"time"
)

type Booking struct {
    ID              uint       `gorm:"primaryKey;autoIncrement;unique" json:"id"`
    CustomerID      uint       `json:"customer_id"`
    CarID          uint       `json:"car_id"`
    StartRent       time.Time  `json:"start_rent"`
    EndRent         time.Time  `json:"end_rent"`
    DriverID       *uint      `json:"driver_id" gorm:"default:null"`
    BookTypeID     *uint      `json:"book_type_id" gorm:"default:null"`
    TotalCost      int        `json:"total_cost"`
    TotalDriverCost int       `json:"total_driver_cost" gorm:"default:null"`
    Finished       bool       `json:"finished"`
    Discount       int        `json:"discount" gorm:"default:null"`
    CreatedAt      time.Time  `json:"created_at"`
    UpdatedAt      time.Time  `json:"updated_at"`

    Driver         *Driver     `gorm:"foreignKey:DriverID" json:"driver"`
    BookingType    *BookingType `gorm:"foreignKey:BookTypeID" json:"booking_type"`
    Customer       Customer    `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
    Car           Car         `gorm:"foreignKey:CarID" json:"car,omitempty"`
}

type InputBooking struct {
    CustomerID  uint `json:"customer_id" binding:"required"`
    CarID       uint `json:"car_id" binding:"required"`
    StartRent   string `json:"start_rent" binding:"required"`
    EndRent     string `json:"end_rent" binding:"required"`
    DriverID       *uint      `json:"driver_id" gorm:"default:null"`
    BookTypeID     *uint      `json:"book_type_id" gorm:"default:null"`
    Finished    bool `json:"finished"`
}