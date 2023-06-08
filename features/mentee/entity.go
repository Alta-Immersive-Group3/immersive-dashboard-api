package mentee

import "time"

type Core struct {
	Id              uint64
	IdClass         uint64 `validate:"required"`
	FullName        string `validate:"required"`
	NickName        string
	Email           string `validate:"required,email"`
	Phone           string `validate:"required"`
	CurrentAddress  string
	HomeAddress     string
	Telegram        string
	IdStatus        uint64 `validate:"required"`
	Gender          string
	EducationType   string
	Major           string
	Graduate        uint32
	Institution     string
	EmergencyName   string
	EmergencyPhone  string
	EmergencyStatus string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type MenteeDataInterface interface {
	Insert(input Core) (uint64, error)
	SelectAll() ([]Core, error)
	SelectById(id uint64) (Core, error)
	UpdateById(id uint64, input Core) error
	DeleteById(id uint64) error
}

type MenteeServiceInterface interface {
	Create(input Core) (Core, error)
	GetAll() ([]Core, error)
	GetById(id uint64) (Core, error)
	UpdateById(id uint64, input Core) (Core, error)
	DeleteById(id uint64) error
}
