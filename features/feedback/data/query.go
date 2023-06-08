package data

import (
	"errors"

	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/feedback"
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/mentee"
	_menteeData "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/mentee/data"
	"gorm.io/gorm"
)

type feedbackQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) feedback.FeedbackDataInterface {
	return &feedbackQuery{
		db: db,
	}
}

func (repo *feedbackQuery) Insert(input feedback.Core) (uint64, error) {
	feedbackInputGorm := CoreToModel(input)

	tx := repo.db.Create(&feedbackInputGorm)
	if tx.Error != nil {
		return 0, tx.Error
	}

	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed, row affected = 0")
	}

	return feedbackInputGorm.Id, nil
}

func (repo *feedbackQuery) SelectAll() ([]feedback.Core, error) {
	var feedbacksData []Feedback
	tx := repo.db.Find(&feedbacksData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, errors.New("error feedbacks not found")
	}

	var feedbacksCoreAll []feedback.Core
	for _, value := range feedbacksData {
		feedbackCore := ModelToCore(value)
		feedbacksCoreAll = append(feedbacksCoreAll, feedbackCore)
	}
	return feedbacksCoreAll, nil
}

func (repo *feedbackQuery) SelectAllByMenteeId(idMentee uint64) ([]feedback.Core, mentee.Core, error) {
	var feedbacksData []Feedback
	tx := repo.db.Where("IdMentee = ?", idMentee).Find(&feedbacksData)
	if tx.Error != nil {
		return nil, mentee.Core{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, mentee.Core{}, errors.New("error feedbacks not found")
	}

	var feedbacksCoreAll []feedback.Core
	for _, value := range feedbacksData {
		feedbackCore := ModelToCore(value)
		feedbacksCoreAll = append(feedbacksCoreAll, feedbackCore)
	}

	var feedbackData Feedback
	tx = repo.db.Where("IdMentee = ?", idMentee).Preload("Mentee", "id = ?", idMentee).First(&feedbackData)
	if tx.Error != nil {
		return nil, mentee.Core{}, tx.Error
	}

	menteeCore := _menteeData.ModelToCore(feedbackData.Mentee)
	return feedbacksCoreAll, menteeCore, nil
}

func (repo *feedbackQuery) SelectById(id uint64) (feedback.Core, error) {
	var feedbackGorm Feedback
	tx := repo.db.First(&feedbackGorm, id)
	if tx.Error != nil {
		return feedback.Core{}, errors.New("error feedback not found")
	}

	feedbackCore := ModelToCore(feedbackGorm)
	return feedbackCore, nil
}

func (repo *feedbackQuery) UpdateById(id uint64, input feedback.Core) error {
	var feedbackGorm Feedback
	tx := repo.db.First(&feedbackGorm, id)
	if tx.Error != nil {
		return errors.New("error feedback not found")
	}

	feedbackInputGorm := CoreToModel(input)
	tx = repo.db.Model(&feedbackGorm).Updates(feedbackInputGorm)
	if tx.Error != nil {
		return errors.New(tx.Error.Error() + "failed to update feedback")
	}

	return nil
}

func (repo *feedbackQuery) DeleteById(id uint64) error {
	var feedbackGorm Feedback
	tx := repo.db.First(&feedbackGorm, id)
	if tx.Error != nil {
		return errors.New("error feedback not found")
	}

	tx = repo.db.Delete(&feedbackGorm, id)
	if tx.Error != nil {
		return errors.New(tx.Error.Error() + "failed to delete feedback")
	}

	return nil
}
