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

	token, errToken := middlewares.CreateToken(int(userGorm.ID))
	if errToken != nil {
		return user.Core{}, "", errToken
	}

	dataCore := ModelToCore(userGorm)
	return dataCore, token, nil
}
