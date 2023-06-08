package handler

import (
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/mentee"
)

type MenteeRequest struct {
	IdClass         uint64 `json:"id_class" form:"id_class"`
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

func MenteeRequestToCore(menteeRequest MenteeRequest) mentee.Core {
	return mentee.Core{
		IdClass:         menteeRequest.IdClass,
		FullName:        menteeRequest.FullName,
		NickName:        menteeRequest.NickName,
		Email:           menteeRequest.Email,
		Phone:           menteeRequest.Phone,
		CurrentAddress:  menteeRequest.CurrentAddress,
		HomeAddress:     menteeRequest.HomeAddress,
		Telegram:        menteeRequest.Telegram,
		IdStatus:        menteeRequest.IdStatus,
		Gender:          menteeRequest.Gender,
		EducationType:   menteeRequest.EducationType,
		Major:           menteeRequest.Major,
		Graduate:        menteeRequest.Graduate,
		Institution:     menteeRequest.Institution,
		EmergencyName:   menteeRequest.EmergencyName,
		EmergencyPhone:  menteeRequest.EmergencyPhone,
		EmergencyStatus: menteeRequest.EmergencyStatus,
	}
}
