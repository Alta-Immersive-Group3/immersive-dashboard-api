package data

import (
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/team"
	"gorm.io/gorm"
)

type teamQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) team.TeamDataInterface {
	return &teamQuery{
		db: db,
	}
}

func (repo *teamQuery) SelectAll() ([]team.Core, error) {
	var teamsData []Team
	tx := repo.db.Find(&teamsData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var teamsCoreAll []team.Core
	for _, value := range teamsData {
		var teamCore = team.Core{
			Id:        value.Id,
			Name:      value.Name,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}
		teamsCoreAll = append(teamsCoreAll, teamCore)
	}
	return teamsCoreAll, nil
}
