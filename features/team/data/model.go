package data

import (
	"time"
)

type Team struct {
	Id        uint64 `gorm:"primarykey"`
	Name      string `gorm:"unique,not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
