package handler

import (
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/feedback"
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/mentee"
	_menteeHandler "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/mentee/handler"
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

type MenteeFeedbacksResponse struct {
	Mentee    _menteeHandler.MenteeResponse
	Feedbacks []FeedbackResponse `json:"feedbacks,omitempty"`
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
	var feedbacksResponse []FeedbackResponse
	for _, value := range feedbacks {
		feedbacksResponse = append(feedbacksResponse, CoreToGetFeedbackResponse(value))
	}

	return MenteeFeedbacksResponse{
		Mentee:    _menteeHandler.CoreToGetMenteeResponse(mentee),
		Feedbacks: feedbacksResponse,
	}
}
