package models

import (
	"time"
)

type Booking struct {
    ID          uint   `gorm:"primaryKey;autoIncrement;unique" json:"id"`
    CustomerID  uint `json:"customer_id"`
    CarID       uint `json:"car_id"`
    StartRent   time.Time `json:"start_rent"`
    EndRent     time.Time `json:"end_rent"`

    // DriverID   *uint `json:"driver_id" gorm:"default:null"`
    // DaysOfRent  int   `json:"days_of_rent"`
    // BookType    string `json:"book_type"` 
    TotalCost  int   `json:"total_cost"`
    Finished bool `json:"finished"`
    // Discount int `json:"discount"`
    CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`

    Customer    Customer `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
    Car         Car      `gorm:"foreignKey:CarID" json:"car,omitempty"`
    // Driver      *Driver  `gorm:"foreignKey:DriverID"`
}

type InputBooking struct {
    CustomerID  uint `json:"customer_id" binding:"required"`
    CarID       uint `json:"car_id" binding:"required"`
    StartRent   string `json:"start_rent" binding:"required"`
    EndRent     string `json:"end_rent" binding:"required"`
    TotalCost   int   `json:"total_cost" binding:"required"`
    Finished    bool `json:"finished"`
}