package data

import (
	"time"

	_classData "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/class/data"
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/mentee"
	_statusData "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/status/data"
	"gorm.io/gorm"
)

type Mentee struct {
	Id              uint64 `gorm:"primarykey"`
	IdClass         uint64 `gorm:"not null"`
	FullName        string `gorm:"not null"`
	NickName        string
	Email           string `gorm:"not null"`
	Phone           string `gorm:"not null"`
	CurrentAddress  string
	HomeAddress     string
	Telegram        string
	IdStatus        uint64 `gorm:"not null"`
	Gender          string `gorm:"enum('L', 'P')"`
	EducationType   string `gorm:"enum('informatics', 'non-informatics')"`
	Major           string
	Graduate        uint32
	Institution     string
	EmergencyName   string
	EmergencyPhone  string
	EmergencyStatus string `gorm:"enum('orang tua', 'saudara kandung', 'kakek nenek', 'keluarga')"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt     `gorm:"index"`
	Class           _classData.Class   `gorm:"foreignKey:IdClass"`
	Status          _statusData.Status `gorm:"foreignKey:IdStatus"`
}

func ModelToCore(menteeData Mentee) mentee.Core {
	return mentee.Core{
		Id:              menteeData.Id,
		IdClass:         menteeData.IdClass,
		FullName:        menteeData.FullName,
		NickName:        menteeData.NickName,
		Email:           menteeData.Email,
		Phone:           menteeData.Phone,
		CurrentAddress:  menteeData.CurrentAddress,
		HomeAddress:     menteeData.HomeAddress,
		Telegram:        menteeData.Telegram,
		IdStatus:        menteeData.IdStatus,
		Gender:          menteeData.Gender,
		EducationType:   menteeData.EducationType,
		Major:           menteeData.Major,
		Graduate:        menteeData.Graduate,
		Institution:     menteeData.Institution,
		EmergencyName:   menteeData.EmergencyName,
		EmergencyPhone:  menteeData.EmergencyPhone,
		EmergencyStatus: menteeData.EmergencyStatus,
		CreatedAt:       menteeData.CreatedAt,
		UpdatedAt:       menteeData.UpdatedAt,
	}
}

func CoreToModel(dataCore mentee.Core) Mentee {
	return Mentee{
		Id:              dataCore.Id,
		IdClass:         dataCore.IdClass,
		FullName:        dataCore.FullName,
		NickName:        dataCore.NickName,
		Email:           dataCore.Email,
		Phone:           dataCore.Phone,
		CurrentAddress:  dataCore.CurrentAddress,
		HomeAddress:     dataCore.HomeAddress,
		Telegram:        dataCore.Telegram,
		IdStatus:        dataCore.IdStatus,
		Gender:          dataCore.Gender,
		EducationType:   dataCore.EducationType,
		Major:           dataCore.Major,
		Graduate:        dataCore.Graduate,
		Institution:     dataCore.Institution,
		EmergencyName:   dataCore.EmergencyName,
		EmergencyPhone:  dataCore.EmergencyPhone,
		EmergencyStatus: dataCore.EmergencyStatus,
		CreatedAt:       dataCore.CreatedAt,
		UpdatedAt:       dataCore.UpdatedAt,
	}
}
