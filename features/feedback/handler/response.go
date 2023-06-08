package handler

import (
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/feedback"
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/mentee"
	_menteeHandler "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/mentee/handler"
)

type FeedbackResponse struct {
	Id       uint64 `json:"id" form:"id"`
	Notes    string `json:"notes" form:"notes"`
	Proof    string `json:"proof" form:"proof"`
	IdUser   uint64 `json:"id_user" form:"id_user"`
	IdMentee uint64 `json:"id_mentee" form:"id_mentee"`
	IdStatus uint64 `json:"id_status" form:"id_status"`
	Approved bool   `json:"approved" form:"approved"`
}

type MenteeFeedbacksResponse struct {
	Mentee    _menteeHandler.MenteeResponse
	Feedbacks []FeedbackResponse
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
