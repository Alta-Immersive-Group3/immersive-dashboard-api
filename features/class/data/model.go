package data

import (
	"time"

	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/class"
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/user/data"
	"gorm.io/gorm"
)

type Class struct {
	Id           uint64 `gorm:"primarykey"`
	Name         string `gorm:"not null"`
	PIC          uint64 `gorm:"not null"`
	StartDate    string `gorm:"type:date"`
	GraduateDate string `gorm:"type:date"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	User         data.User      `gorm:"foreignKey:PIC"`
}

func ModelToCore(classData Class) class.Core {
	return class.Core{
		Id:           classData.Id,
		Name:         classData.Name,
		PIC:          classData.PIC,
		StartDate:    classData.StartDate,
		GraduateDate: classData.GraduateDate,
		CreatedAt:    classData.CreatedAt,
		UpdatedAt:    classData.UpdatedAt,
	}
}

func CoreToModel(dataCore class.Core) Class {
	return Class{
		Id:           dataCore.Id,
		Name:         dataCore.Name,
		PIC:          dataCore.PIC,
		StartDate:    dataCore.StartDate,
		GraduateDate: dataCore.GraduateDate,
		CreatedAt:    dataCore.CreatedAt,
		UpdatedAt:    dataCore.UpdatedAt,
	}
}
