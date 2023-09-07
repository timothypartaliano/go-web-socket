package model

import (
	"time"
)

type User struct {
	ID            uint      `gorm:"primaryKey"`
	Username      string    `json:"username" gorm:"not null"`
	Password      string    `json:"password" gorm:"not null"`
	DepositAmount int       `json:"deposit_amount" gorm:"not null"`
	CreatedAt     time.Time `json:"created_at" gorm:"-"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"-"`
	DeletedAt     time.Time `json:"deleted_at" gorm:"-"`
}