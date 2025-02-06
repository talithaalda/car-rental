package models

import (
	"time"
)

type Membership struct {
    ID                  uint   `gorm:"primaryKey;autoIncrement;unique" json:"id"`
    MembershipName      string `json:"membership_name"`
    Discount            int    `json:"dsicount"`
    CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}