package service

import (
	"errors"

	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/user"
	"github.com/go-playground/validator/v10"
)

type userService struct {
	userData user.UserDataInterface
	validate *validator.Validate
}

func New(repo user.UserDataInterface) user.UserServiceInterface {
	return &userService{
		userData: repo,
		validate: validator.New(),
	}
}

func (service *userService) Login(email string, password string) (user.Core, string, error) {
	if email == "" || password == "" {
		return user.Core{}, "", errors.New("error validation: email or password cannot be empty")
	}
	dataLogin, token, err := service.userData.Login(email, password)
	return dataLogin, token, err
}

func (service *userService) Create(input user.Core) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}
	errInsert := service.userData.Insert(input)
	return errInsert
}

func (service *userService) GetAll() ([]user.Core, error) {
	data, err := service.userData.SelectAll()
	if err != nil {
		return nil, err
	}
	return data, err
}

func (service *userService) GetById(id uint64) (user.Core, error) {
	data, err := service.userData.SelectById(id)
	if err != nil {
		return user.Core{}, err
	}
	return data, err
}

func (service *userService) UpdateById(id uint64, input user.Core) error {
	errUpdate := service.userData.UpdateById(id, input)
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}

func (service *userService) DeleteById(id uint64) error {
	errUpdate := service.userData.DeleteById(id)
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}
