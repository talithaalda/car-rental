package models

import (
	"time"
)

type Car struct {
    ID       uint   `gorm:"primaryKey;autoIncrement;unique" json:"id"`
    Name     string `json:"name"`
    Stock     int    `json:"stock"`
    DailyRent int    `json:"daily_rent"`
    CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
type InputCar struct {
    Name     string `json:"name" binding:"required"`
    Stock     int    `json:"stock" binding:"required,gt=0"`
    DailyRent int    `json:"daily_rent" binding:"required,gt=0"`
}