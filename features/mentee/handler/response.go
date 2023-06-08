package handler

import (
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/mentee"
)

type MenteeResponse struct {
	Id              uint64 `json:"id" form:"id"`
	FullName        string `json:"full_name" form:"full_name"`
	NickName        string `json:"nick_name" form:"nick_name"`
	Email           string `json:"email" form:"email"`
	Phone           string `json:"phone" form:"phone"`
	CurrentAddress  string `json:"current_address" form:"current_address"`
	HomeAddress     string `json:"home_address" form:"home_address"`
	Telegram        string `json:"telegram" form:"telegram"`
	IdStatus        uint64 `json:"id_status" form:"id_status"`
	Gender          string `json:"gender" form:"gender"`
	EducationType   string `json:"education_type" form:"education_type"`
	Major           string `json:"major" form:"major"`
	Graduate        uint32 `json:"graduate" form:"graduate"`
	Institution     string `json:"institution" form:"institution"`
	EmergencyName   string `json:"emergency_name" form:"emergency_name"`
	EmergencyPhone  string `json:"emergency_phone" form:"emergency_phone"`
	EmergencyStatus string `json:"emergency_status" form:"emergency_status"`
}

func CoreToGetMenteeResponse(mentee mentee.Core) MenteeResponse {
	return MenteeResponse{
		FullName:        mentee.FullName,
		NickName:        mentee.NickName,
		Email:           mentee.Email,
		Phone:           mentee.Phone,
		CurrentAddress:  mentee.CurrentAddress,
		HomeAddress:     mentee.HomeAddress,
		Telegram:        mentee.Telegram,
		IdStatus:        mentee.IdStatus,
		Gender:          mentee.Gender,
		EducationType:   mentee.EducationType,
		Major:           mentee.Major,
		Graduate:        mentee.Graduate,
		Institution:     mentee.Institution,
		EmergencyName:   mentee.EmergencyName,
		EmergencyPhone:  mentee.EmergencyPhone,
		EmergencyStatus: mentee.EmergencyStatus,
	}
}
