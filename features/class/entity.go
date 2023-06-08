package class

import "time"

type Core struct {
	Id           uint64
	Name         string `validate:"required"`
	PIC          uint64 `validate:"required"`
	StartDate    string
	GraduateDate string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type ClassDataInterface interface {
	Insert(input Core) (uint64, error)
	SelectAll() ([]Core, error)
	SelectById(id uint64) (Core, error)
	UpdateById(id uint64, input Core) error
	DeleteById(id uint64) error
}

type ClassServiceInterface interface {
	Create(input Core) (Core, error)
	GetAll() ([]Core, error)
	GetById(id uint64) (Core, error)
	UpdateById(id uint64, input Core) (Core, error)
	DeleteById(id uint64) error
}
