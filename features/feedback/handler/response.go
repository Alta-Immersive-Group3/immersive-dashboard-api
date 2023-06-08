package handler

import (
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/feedback"
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/mentee"
)

type FeedbackResponse struct {
	Id       uint64 `json:"id,omitempty"`
	Notes    string `json:"notes,omitempty"`
	Proof    string `json:"proof,omitempty"`
	IdUser   uint64 `json:"id_user,omitempty"`
	IdMentee uint64 `json:"id_mentee,omitempty"`
	IdStatus uint64 `json:"id_status,omitempty"`
	Approved bool   `json:"approved,omitempty"`
}

type FeedbackMenteeResponse struct {
	Id       uint64 `json:"id,omitempty"`
	IdStatus uint64 `json:"id_status,omitempty"`
	Notes    string `json:"notes,omitempty"`
	Proof    string `json:"proof,omitempty"`
}

type MenteeFeedbacksResponse struct {
	IdMentee  uint64                   `json:"id_mentee,omitempty"`
	Name      string                   `json:"name,omitempty"`
	Feedbacks []FeedbackMenteeResponse `json:"feedbacks"`
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

func MenteeFeedbackCoreToResponse(mentee mentee.Core, feedbacks []feedback.Core) MenteeFeedbacksResponse {
	var feedbacksResponse []FeedbackMenteeResponse
	for _, value := range feedbacks {
		feedbacksResponse = append(feedbacksResponse, FeedbackMenteeResponse{
			Id:       value.Id,
			IdStatus: value.IdStatus,
			Notes:    value.Notes,
			Proof:    value.Proof,
		})
	}

	return MenteeFeedbacksResponse{
		IdMentee:  mentee.Id,
		Name:      mentee.FullName,
		Feedbacks: feedbacksResponse,
	}
}
