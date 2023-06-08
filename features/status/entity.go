package status

import "time"

type Core struct {
	Id        uint64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type StatusDataInterface interface {
	SelectAll() ([]Core, error)
}

type StatusServiceInterface interface {
	GetAll() ([]Core, error)
}
