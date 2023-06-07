package team

import "time"

type Core struct {
	Id        uint64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TeamDataInterface interface {
	SelectAll() ([]Core, error)
}

type TeamServiceInterface interface {
	GetAll() ([]Core, error)
}
