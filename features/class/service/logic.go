package service

import (
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/class"
	"github.com/go-playground/validator/v10"
)

type classService struct {
	classData class.ClassDataInterface
	validate  *validator.Validate
}

func New(repo class.ClassDataInterface) class.ClassServiceInterface {
	return &classService{
		classData: repo,
		validate:  validator.New(),
	}
}

func (service *classService) Create(input class.Core) (class.Core, error) {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return class.Core{}, errValidate
	}

	id, errInsert := service.classData.Insert(input)
	if errInsert != nil {
		return class.Core{}, errInsert
	}

	return service.GetById(id)
}

func (service *classService) GetAll() ([]class.Core, error) {
	data, err := service.classData.SelectAll()
	if err != nil {
		return nil, err
	}
	return data, err
}

func (service *classService) GetById(id uint64) (class.Core, error) {
	data, err := service.classData.SelectById(id)
	if err != nil {
		return class.Core{}, err
	}
	return data, err
}

func (service *classService) UpdateById(id uint64, input class.Core) (class.Core, error) {
	errUpdate := service.classData.UpdateById(id, input)
	if errUpdate != nil {
		return class.Core{}, errUpdate
	}

	return service.GetById(id)
}

func (service *classService) DeleteById(id uint64) error {
	errUpdate := service.classData.DeleteById(id)
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}
