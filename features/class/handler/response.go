package handler

import (
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/class"
)

type ClassResponse struct {
	Id           uint64 `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	PIC          uint64 `json:"pic,omitempty"`
	StartDate    string `json:"start_date,omitempty"`
	GraduateDate string `json:"graduate_date,omitempty"`
}

func CoreToGetClassResponse(class class.Core) ClassResponse {
	return ClassResponse{
		Id:           class.Id,
		Name:         class.Name,
		PIC:          class.PIC,
		StartDate:    class.StartDate,
		GraduateDate: class.GraduateDate,
	}
}
