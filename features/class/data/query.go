package data

import (
	"errors"

	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/class"
	"gorm.io/gorm"
)

type classQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) class.ClassDataInterface {
	return &classQuery{
		db: db,
	}
}

func (repo *classQuery) Insert(input class.Core) (uint64, error) {
	classInputGorm := CoreToModel(input)

	tx := repo.db.Create(&classInputGorm)
	if tx.Error != nil {
		return 0, tx.Error
	}

	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed, row affected = 0")
	}

	return classInputGorm.Id, nil
}

func (repo *classQuery) SelectAll() ([]class.Core, error) {
	var classsData []Class
	tx := repo.db.Find(&classsData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, errors.New("error classs not found")
	}

	var classsCoreAll []class.Core
	for _, value := range classsData {
		classCore := ModelToCore(value)
		classsCoreAll = append(classsCoreAll, classCore)
	}
	return classsCoreAll, nil
}

func (repo *classQuery) SelectById(id uint64) (class.Core, error) {
	var classGorm Class
	tx := repo.db.First(&classGorm, id)
	if tx.Error != nil {
		return class.Core{}, errors.New("error class not found")
	}

	classCore := ModelToCore(classGorm)
	return classCore, nil
}

func (repo *classQuery) UpdateById(id uint64, input class.Core) error {
	var classGorm Class
	tx := repo.db.First(&classGorm, id)
	if tx.Error != nil {
		return errors.New("error class not found")
	}

	classInputGorm := CoreToModel(input)
	tx = repo.db.Model(&classGorm).Updates(classInputGorm)
	if tx.Error != nil {
		return errors.New(tx.Error.Error() + "failed to update class")
	}

	return nil
}

func (repo *classQuery) DeleteById(id uint64) error {
	var classGorm Class
	tx := repo.db.First(&classGorm, id)
	if tx.Error != nil {
		return errors.New("error class not found")
	}

	tx = repo.db.Delete(&classGorm, id)
	if tx.Error != nil {
		return errors.New(tx.Error.Error() + "failed to delete class")
	}

	return nil
}
