package data

import (
	"time"

	"gorm.io/gorm"
)

type Status struct {
	Id        uint64 `gorm:"primarykey"`
	Name      string `gorm:"unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
