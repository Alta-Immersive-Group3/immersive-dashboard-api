package data

import (
	"time"

	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/team/data"
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/user"
)

type User struct {
	Id        uint64    `gorm:"primarykey"`
	FullName  string    `gorm:"not null"`
	Email     string    `gorm:"unique,not null"`
	Password  string    `gorm:"not null"`
	Role      string    `gorm:"enum('admin', 'user');not null"`
	IdTeam    uint64    `gorm:"not null"`
	Status    bool      `gorm:"default:true"`
	IsDeleted bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
	Team      data.Team `gorm:"foreignKey:IdTeam"`
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
		IsDeleted: userData.IsDeleted,
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
		IsDeleted: dataCore.IsDeleted,
		CreatedAt: dataCore.CreatedAt,
		UpdatedAt: dataCore.UpdatedAt,
	}
}
