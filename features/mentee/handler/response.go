package handler

import (
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/mentee"
)

type MenteeResponse struct {
	Id              uint64 `json:"id,omitempty"`
	IdClass         uint64 `json:"id_class,omitempty"`
	FullName        string `json:"full_name,omitempty"`
	NickName        string `json:"nick_name,omitempty"`
	Email           string `json:"email,omitempty"`
	Phone           string `json:"phone,omitempty"`
	CurrentAddress  string `json:"current_address,omitempty"`
	HomeAddress     string `json:"home_address,omitempty"`
	Telegram        string `json:"telegram,omitempty"`
	IdStatus        uint64 `json:"id_status,omitempty"`
	Gender          string `json:"gender,omitempty"`
	EducationType   string `json:"education_type,omitempty"`
	Major           string `json:"major,omitempty"`
	Graduate        uint32 `json:"graduate,omitempty"`
	Institution     string `json:"institution,omitempty"`
	EmergencyName   string `json:"emergency_name,omitempty"`
	EmergencyPhone  string `json:"emergency_phone,omitempty"`
	EmergencyStatus string `json:"emergency_status,omitempty"`
}

func CoreToGetMenteeResponse(mentee mentee.Core) MenteeResponse {
	return MenteeResponse{
		Id:              mentee.Id,
		IdClass:         mentee.IdClass,
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
