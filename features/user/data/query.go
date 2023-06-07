package data

import (
	"errors"

	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/app/middlewares"
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/user"
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/helper"
	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserDataInterface {
	return &userQuery{
		db: db,
	}
}

func (repo *userQuery) Login(email string, password string) (user.Core, string, error) {
	var userGorm User
	tx := repo.db.Where("email = ?", email).First(&userGorm)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return user.Core{}, "", errors.New("login failed, wrong email and password")
		} else {
			return user.Core{}, "", tx.Error
		}
	}

	checkPassword := helper.CheckPasswordHash(password, userGorm.Password)
	if !checkPassword {
		return user.Core{}, "", errors.New("login failed, wrong password")
	}

	token, errToken := middlewares.CreateToken(int(userGorm.Id))
	if errToken != nil {
		return user.Core{}, "", errToken
	}

	dataCore := ModelToCore(userGorm)
	return dataCore, token, nil
}

func (repo *userQuery) Insert(input user.Core) error {
	hashedPassword, errHash := helper.HashPassword(input.Password)
	if errHash != nil {
		return errors.New("error hash password")
	}
	userInputGorm := CoreToModel(input)
	userInputGorm.Password = hashedPassword

	tx := repo.db.Create(&userInputGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}

	return nil
}

func (repo *userQuery) SelectAll() ([]user.Core, error) {
	var usersData []User
	tx := repo.db.Find(&usersData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, errors.New("error users not found")
	}

	var usersCoreAll []user.Core
	for _, value := range usersData {
		userCore := ModelToCore(value)
		usersCoreAll = append(usersCoreAll, userCore)
	}
	return usersCoreAll, nil
}

func (repo *userQuery) SelectById(id uint64) (user.Core, error) {
	var userData User
	tx := repo.db.Where("id = ?", id).Find(&userData)
	if tx.Error != nil {
		return user.Core{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return user.Core{}, errors.New("error user not found")
	}

	userCore := ModelToCore(userData)
	return userCore, nil
}

func (repo *userQuery) UpdateById(id uint64, input user.Core) error {
	var userGorm User
	tx := repo.db.First(&userGorm, id)
	if tx.Error != nil {
		return errors.New("error user not found")
	}

	userInputGorm := CoreToModel(input)
	if userInputGorm.Password != "" {
		hashedPassword, errHash := helper.HashPassword(userInputGorm.Password)
		if errHash != nil {
			return errors.New("error hash password")
		}
		userInputGorm.Password = hashedPassword
	}

	tx = repo.db.Model(&userGorm).Updates(userInputGorm)
	if tx.Error != nil {
		return errors.New(tx.Error.Error() + "failed to update user")
	}

	return nil
}

func (repo *userQuery) DeleteById(id uint64) error {
	var userGorm User
	tx := repo.db.Delete(&userGorm, id)
	if tx.Error != nil {
		return errors.New(tx.Error.Error() + "failed to delete user")
	}

	return nil
}
