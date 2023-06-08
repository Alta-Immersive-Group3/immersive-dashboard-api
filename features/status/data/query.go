package data

import (
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/status"
	"gorm.io/gorm"
)

type statusQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) status.StatusDataInterface {
	return &statusQuery{
		db: db,
	}
}

func (repo *statusQuery) SelectAll() ([]status.Core, error) {
	var statussData []Status
	tx := repo.db.Find(&statussData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var statussCoreAll []status.Core
	for _, value := range statussData {
		var statusCore = status.Core{
			Id:        value.Id,
			Name:      value.Name,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}
		statussCoreAll = append(statussCoreAll, statusCore)
	}
	return statussCoreAll, nil
}
