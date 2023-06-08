package data

import (
	"time"

	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/feedback"
	_menteeData "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/mentee/data"
	_statusData "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/status/data"
	_userData "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/user/data"
	"gorm.io/gorm"
)

type Feedback struct {
	Id        uint64 `gorm:"primarykey"`
	Notes     string `gorm:"not null"`
	Proof     string `gorm:"default:'placeholder.jpg'"`
	IdUser    uint64 `gorm:"not null"`
	IdMentee  uint64 `gorm:"not null"`
	IdStatus  uint64 `gorm:"not null"`
	Approved  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt     `gorm:"index"`
	User      _userData.User     `gorm:"foreignKey:IdUser"`
	Mentee    _menteeData.Mentee `gorm:"foreignKey:IdMentee"`
	Status    _statusData.Status `gorm:"foreignKey:IdStatus"`
}

func ModelToCore(feedbackData Feedback) feedback.Core {
	return feedback.Core{
		Id:        feedbackData.Id,
		Notes:     feedbackData.Notes,
		Proof:     feedbackData.Proof,
		IdUser:    feedbackData.IdUser,
		IdMentee:  feedbackData.IdMentee,
		IdStatus:  feedbackData.IdStatus,
		Approved:  feedbackData.Approved,
		CreatedAt: feedbackData.CreatedAt,
		UpdatedAt: feedbackData.UpdatedAt,
	}
}

func CoreToModel(dataCore feedback.Core) Feedback {
	return Feedback{
		Id:        dataCore.Id,
		Notes:     dataCore.Notes,
		Proof:     dataCore.Proof,
		IdUser:    dataCore.IdUser,
		IdMentee:  dataCore.IdMentee,
		IdStatus:  dataCore.IdStatus,
		Approved:  dataCore.Approved,
		CreatedAt: dataCore.CreatedAt,
		UpdatedAt: dataCore.UpdatedAt,
	}
}
