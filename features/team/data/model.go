package data

import (
	"time"

	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	ID        uint64 `gorm:"primarykey"`
	Name      string `gorm:"unique,not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
