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
	IsDeleted bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserDataInterface interface {
	Login(email string, password string) (Core, string, error)
}

type UserServiceInterface interface {
	Login(email string, password string) (Core, string, error)
}
