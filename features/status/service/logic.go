package service

import (
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/status"
)

type statusService struct {
	statusData status.StatusDataInterface
}

func New(repo status.StatusDataInterface) status.StatusServiceInterface {
	return &statusService{
		statusData: repo,
	}
}

func (service *statusService) GetAll() ([]status.Core, error) {
	data, err := service.statusData.SelectAll()
	if err != nil {
		return nil, err
	}
	return data, err
}
