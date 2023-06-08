package handler

import "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/class"

type ClassRequest struct {
	Name         string `json:"name" form:"name"`
	PIC          uint64 `json:"pic" form:"pic"`
	StartDate    string `json:"start_date" form:"start_date"`
	GraduateDate string `json:"graduate_date" form:"graduate_date"`
}

func classRequestToCore(classRequest ClassRequest) class.Core {
	return class.Core{
		Name:         classRequest.Name,
		PIC:          classRequest.PIC,
		StartDate:    classRequest.StartDate,
		GraduateDate: classRequest.GraduateDate,
	}
}
