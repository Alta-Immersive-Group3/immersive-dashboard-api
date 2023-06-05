package data

import (
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/team/data"
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/user"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint64    `gorm:"primarykey"`
	FullName  string    `gorm:"not null"`
	Email     string    `gorm:"unique,not null"`
	Password  string    `gorm:"not null"`
	Role      string    `gorm:"enum('admin', 'user');not null"`
	IdTeam    uint64    `gorm:"not null"`
	Status    bool      `gorm:"default:true"`
	IsDeleted bool      `gorm:"default:false"`
	Team      data.Team `gorm:"foreignKey:IdTeam"`
}

func ModelToCore(userData User) user.Core {
	return user.Core{
		Id:        userData.ID,
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
