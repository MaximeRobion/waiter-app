package models

import (
	"time"
)

type Table struct {
	ID        uint      `gorm:"primary_key"`
	Name      string    `gorm:"not null"`
	Capacity  int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
