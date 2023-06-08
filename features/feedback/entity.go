package feedback

import "time"

type Core struct {
	Id        uint64
	Notes     string `validate:"required"`
	Proof     string
	IdUser    uint64 `validate:"required"`
	IdMentee  uint64 `validate:"required"`
	IdStatus  uint64 `validate:"required"`
	Approved  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type FeedbackDataInterface interface {
	Insert(input Core) (uint64, error)
	SelectAll() ([]Core, error)
	SelectAllByMenteeId(idMentee uint64) ([]Core, error)
	SelectById(id uint64) (Core, error)
	UpdateById(id uint64, input Core) error
	DeleteById(id uint64) error
}

type FeedbackServiceInterface interface {
	Create(input Core) (Core, error)
	GetAll() ([]Core, error)
	GetAllByMenteeId(idMentee uint64) ([]Core, error)
	GetById(id uint64) (Core, error)
	UpdateById(id uint64, input Core) (Core, error)
	DeleteById(id uint64) error
}
