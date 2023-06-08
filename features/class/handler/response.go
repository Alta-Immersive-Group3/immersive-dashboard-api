package handler

import (
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/class"
)

type ClassResponse struct {
	Id           uint64 `json:"id" form:"id"`
	Name         string `json:"name" form:"name"`
	PIC          uint64 `json:"pic" form:"pic"`
	StartDate    string `json:"start_date" form:"start_date"`
	GraduateDate string `json:"graduate_date" form:"graduate_date"`
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
