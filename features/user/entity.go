package user

import "time"

type Core struct {
	Id        uint64
	FullName  string `validate:"required"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required"`
	Role      string `validate:"required"`
	IdTeam    uint64 `validate:"required"`
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserDataInterface interface {
	Login(email string, password string) (Core, string, error)
	Insert(input Core) (uint64, error)
	SelectAll() ([]Core, error)
	SelectById(id uint64) (Core, error)
	UpdateById(id uint64, input Core) error
	DeleteById(id uint64) error
}

type UserServiceInterface interface {
	Login(email string, password string) (Core, string, error)
	Create(input Core) (Core, error)
	GetAll() ([]Core, error)
	GetById(id uint64) (Core, error)
	UpdateById(id uint64, input Core) (Core, error)
	DeleteById(id uint64) error
}
