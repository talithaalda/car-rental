package models

import (
	"time"
)

type Customer struct {
    ID      uint   `gorm:"primaryKey;autoIncrement;unique" json:"id"`
    Name    string `json:"name"`
    NIK     string `json:"nik" gorm:"unique"`
    Phone   string `json:"phone"`
    CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
    // MembershipID *uint  `json:"membership_id" gorm:"default:null"`
    // Membership   *Membership `gorm:"foreignKey:MembershipID"`
}

type InputCustomer struct {
    Name    string `json:"name" binding:"required"`
    NIK     string `json:"nik" gorm:"unique" binding:"required"`
    Phone   string `json:"phone" binding:"required"`
    
    // MembershipID *uint  `json:"membership_id" gorm:"default:null"`
    // Membership   *Membership `gorm:"foreignKey:MembershipID"`
}
