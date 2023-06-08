package handler

import "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/feedback"

type FeedbackResponse struct {
	Id       uint64 `json:"id" form:"id"`
	Notes    string `json:"notes" form:"notes"`
	Proof    string `json:"proof" form:"proof"`
	IdUser   uint64 `json:"id_user" form:"id_user"`
	IdMentee uint64 `json:"id_mentee" form:"id_mentee"`
	IdStatus uint64 `json:"id_status" form:"id_status"`
	Approved bool   `json:"approved" form:"approved"`
}

func CoreToGetFeedbackResponse(feedback feedback.Core) FeedbackResponse {
	return FeedbackResponse{
		Id:       feedback.Id,
		Notes:    feedback.Notes,
		Proof:    feedback.Proof,
		IdUser:   feedback.IdUser,
		IdMentee: feedback.IdMentee,
		IdStatus: feedback.IdStatus,
		Approved: feedback.Approved,
	}
}
