package data

import (
	"time"

	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/team/data"
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/user"
	"gorm.io/gorm"
)

type User struct {
	Id        uint64 `gorm:"primarykey"`
	FullName  string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"enum('admin', 'user');not null"`
	IdTeam    uint64 `gorm:"not null"`
	Status    bool   `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Team      data.Team      `gorm:"foreignKey:IdTeam"`
}

func ModelToCore(userData User) user.Core {
	return user.Core{
		Id:        userData.Id,
		FullName:  userData.FullName,
		Email:     userData.Email,
		Password:  userData.Password,
		Role:      userData.Role,
		IdTeam:    userData.IdTeam,
		Status:    userData.Status,
		CreatedAt: userData.CreatedAt,
		UpdatedAt: userData.UpdatedAt,
	}
}

func CoreToModel(dataCore user.Core) User {
	return User{
		Id:        dataCore.Id,
		FullName:  dataCore.FullName,
		Email:     dataCore.Email,
		Password:  dataCore.Password,
		Role:      dataCore.Role,
		IdTeam:    dataCore.IdTeam,
		Status:    dataCore.Status,
		CreatedAt: dataCore.CreatedAt,
		UpdatedAt: dataCore.UpdatedAt,
	}
}
