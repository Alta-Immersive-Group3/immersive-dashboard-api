package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/feedback"
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/helper"
	"github.com/labstack/echo/v4"
)

type feedbackHandler struct {
	feedbackService feedback.FeedbackServiceInterface
}

func New(service feedback.FeedbackServiceInterface) *feedbackHandler {
	return &feedbackHandler{
		feedbackService: service,
	}
}

func (handler *feedbackHandler) CreateFeedback(c echo.Context) error {
	FeedbackInput := FeedbackRequest{}
	errBind := c.Bind(&FeedbackInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}

	FeedbackCore := FeedbackRequestToCore(FeedbackInput)
	result, err := handler.feedbackService.Create(FeedbackCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error insert data, "+err.Error()))
		}
	}

	FeedbackResponse := CoreToGetFeedbackResponse(result)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read data", FeedbackResponse))
}

func (handler *feedbackHandler) GetAllFeedback(c echo.Context) error {
	results, err := handler.feedbackService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data, "+err.Error()))
	}

	var FeedbacksResponse []FeedbackResponse
	for _, value := range results {
		FeedbacksResponse = append(FeedbacksResponse, CoreToGetFeedbackResponse(value))
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read data", FeedbacksResponse))
}

func (handler *feedbackHandler) GetAllMenteeFeedback(c echo.Context) error {
	paramId := c.Param("id")
	menteeId, _ := strconv.ParseUint(paramId, 10, 64)

	feedbacks, mentee, err := handler.feedbackService.GetAllByMenteeId(menteeId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data, "+err.Error()))
	}

	FeedbacksResponse := MenteeFeedbackCoreToResponse(mentee, feedbacks)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read data", FeedbacksResponse))
}

func (handler *feedbackHandler) GetFeedbackById(c echo.Context) error {
	paramId := c.Param("id")
	feedbackId, _ := strconv.ParseUint(paramId, 10, 64)

	result, err := handler.feedbackService.GetById(feedbackId)
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.FailedResponse("error read data, "+err.Error()))
	}

	FeedbackResponse := CoreToGetFeedbackResponse(result)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read data", FeedbackResponse))
}

func (handler *feedbackHandler) UpdateFeedbackById(c echo.Context) error {
	paramId := c.Param("id")
	FeedbackId, _ := strconv.ParseUint(paramId, 10, 64)

	FeedbackInput := FeedbackRequest{}
	errBind := c.Bind(&FeedbackInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}

	FeedbackCore := FeedbackRequestToCore(FeedbackInput)
	result, err := handler.feedbackService.UpdateById(FeedbackId, FeedbackCore)
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.FailedResponse("error update data, "+err.Error()))
	}

	FeedbackResponse := CoreToGetFeedbackResponse(result)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success update data", FeedbackResponse))
}

func (handler *feedbackHandler) DeleteFeedbackById(c echo.Context) error {
	paramId := c.Param("id")
	FeedbackId, _ := strconv.ParseUint(paramId, 10, 64)

	err := handler.feedbackService.DeleteById(FeedbackId)
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.FailedResponse("error delete data, "+err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success delete data"))
}
