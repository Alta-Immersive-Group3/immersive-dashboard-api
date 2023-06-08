package data

import (
	"errors"

	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/mentee"
	"gorm.io/gorm"
)

type menteeQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) mentee.MenteeDataInterface {
	return &menteeQuery{
		db: db,
	}
}

func (repo *menteeQuery) Insert(input mentee.Core) (uint64, error) {
	menteeInputGorm := CoreToModel(input)

	tx := repo.db.Create(&menteeInputGorm)
	if tx.Error != nil {
		return 0, tx.Error
	}

	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed, row affected = 0")
	}

	return menteeInputGorm.Id, nil
}

func (repo *menteeQuery) SelectAll() ([]mentee.Core, error) {
	var menteesData []Mentee
	tx := repo.db.Find(&menteesData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, errors.New("error mentees not found")
	}

	var menteesCoreAll []mentee.Core
	for _, value := range menteesData {
		menteeCore := ModelToCore(value)
		menteesCoreAll = append(menteesCoreAll, menteeCore)
	}
	return menteesCoreAll, nil
}

func (repo *menteeQuery) SelectById(id uint64) (mentee.Core, error) {
	var menteeGorm Mentee
	tx := repo.db.First(&menteeGorm, id)
	if tx.Error != nil {
		return mentee.Core{}, errors.New("error mentee not found")
	}

	menteeCore := ModelToCore(menteeGorm)
	return menteeCore, nil
}

func (repo *menteeQuery) UpdateById(id uint64, input mentee.Core) error {
	var menteeGorm Mentee
	tx := repo.db.First(&menteeGorm, id)
	if tx.Error != nil {
		return errors.New("error mentee not found")
	}

	menteeInputGorm := CoreToModel(input)
	tx = repo.db.Model(&menteeGorm).Updates(menteeInputGorm)
	if tx.Error != nil {
		return errors.New(tx.Error.Error() + "failed to update mentee")
	}

	return nil
}

func (repo *menteeQuery) DeleteById(id uint64) error {
	var menteeGorm Mentee
	tx := repo.db.First(&menteeGorm, id)
	if tx.Error != nil {
		return errors.New("error mentee not found")
	}

	tx = repo.db.Delete(&menteeGorm, id)
	if tx.Error != nil {
		return errors.New(tx.Error.Error() + "failed to delete mentee")
	}

	return nil
}
