package service

import (
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/mentee"
	"github.com/go-playground/validator/v10"
)

type menteeService struct {
	menteeData mentee.MenteeDataInterface
	validate   *validator.Validate
}

func New(repo mentee.MenteeDataInterface) mentee.MenteeServiceInterface {
	return &menteeService{
		menteeData: repo,
		validate:   validator.New(),
	}
}

func (service *menteeService) Create(input mentee.Core) (mentee.Core, error) {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return mentee.Core{}, errValidate
	}

	id, errInsert := service.menteeData.Insert(input)
	if errInsert != nil {
		return mentee.Core{}, errInsert
	}

	return service.GetById(id)
}

func (service *menteeService) GetAll() ([]mentee.Core, error) {
	data, err := service.menteeData.SelectAll()
	if err != nil {
		return nil, err
	}
	return data, err
}

func (service *menteeService) GetById(id uint64) (mentee.Core, error) {
	data, err := service.menteeData.SelectById(id)
	if err != nil {
		return mentee.Core{}, err
	}
	return data, err
}

func (service *menteeService) UpdateById(id uint64, input mentee.Core) (mentee.Core, error) {
	errUpdate := service.menteeData.UpdateById(id, input)
	if errUpdate != nil {
		return mentee.Core{}, errUpdate
	}

	return service.GetById(id)
}

func (service *menteeService) DeleteById(id uint64) error {
	errUpdate := service.menteeData.DeleteById(id)
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}
