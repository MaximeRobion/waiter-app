package models

import (
	"time"
)

type Group struct {
	ID        uint      `gorm:"primary_key"`
	TableID   uint      `gorm:"not null"`
	Persons   int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
