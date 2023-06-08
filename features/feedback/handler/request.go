package handler

import "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/feedback"

type FeedbackRequest struct {
	Notes    string `json:"notes" form:"notes"`
	Proof    string `json:"proof" form:"proof"`
	IdUser   uint64 `json:"id_user" form:"id_user"`
	IdMentee uint64 `json:"id_mentee" form:"id_mentee"`
	IdStatus uint64 `json:"id_status" form:"id_status"`
	Approved bool   `json:"approved" form:"approved"`
}

func FeedbackRequestToCore(feedbackRequest FeedbackRequest) feedback.Core {
	return feedback.Core{
		Notes:    feedbackRequest.Notes,
		Proof:    feedbackRequest.Proof,
		IdUser:   feedbackRequest.IdUser,
		IdMentee: feedbackRequest.IdMentee,
		IdStatus: feedbackRequest.IdStatus,
		Approved: feedbackRequest.Approved,
	}
}
