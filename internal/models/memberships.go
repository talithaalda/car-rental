package models

import (
	"time"
)

type Membership struct {
    ID                  uint   `gorm:"primaryKey;autoIncrement;unique" json:"id"`
    MembershipName      string `json:"membership_name"`
    Discount            int    `json:"discount"`
    CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type InputMembership struct {
    MembershipName      string `json:"membership_name"  binding:"required"`
    Discount            int    `json:"discount"          binding:"required,gt=0"`
}