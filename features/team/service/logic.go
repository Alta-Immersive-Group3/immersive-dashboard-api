package service

import (
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/team"
)

type teamService struct {
	teamData team.TeamDataInterface
}

func New(repo team.TeamDataInterface) team.TeamServiceInterface {
	return &teamService{
		teamData: repo,
	}
}

func (service *teamService) GetAll() ([]team.Core, error) {
	data, err := service.teamData.SelectAll()
	if err != nil {
		return nil, err
	}
	return data, err
}
