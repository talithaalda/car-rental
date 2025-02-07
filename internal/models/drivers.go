package models

import (
	"time"
)


type Driver struct {
    ID        uint   `gorm:"primaryKey;autoIncrement;unique" json:"id"`
    Name      string `json:"name"`
    NIK       string `json:"nik" gorm:"unique"`
    Phone     string `json:"phone"`
    DailyCost int    `json:"daily_cost"`
    CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
type InputDriver struct {
    Name      string `json:"name" binding:"required"`
    NIK       string `json:"nik" gorm:"unique" binding:"required"`
    Phone     string `json:"phone" binding:"required"`
    DailyCost int    `json:"daily_cost" binding:"required,gt=0"`
}