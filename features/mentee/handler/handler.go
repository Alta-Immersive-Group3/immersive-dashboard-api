package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/mentee"
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/helper"
	"github.com/labstack/echo/v4"
)

type menteeHandler struct {
	menteeService mentee.MenteeServiceInterface
}

func New(service mentee.MenteeServiceInterface) *menteeHandler {
	return &menteeHandler{
		menteeService: service,
	}
}

func (handler *menteeHandler) CreateMentee(c echo.Context) error {
	menteeInput := MenteeRequest{}
	errBind := c.Bind(&menteeInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}

	menteeCore := MenteeRequestToCore(menteeInput)
	result, err := handler.menteeService.Create(menteeCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error insert data, "+err.Error()))
		}
	}

	menteeResponse := CoreToGetMenteeResponse(result)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read data", menteeResponse))
}

func (handler *menteeHandler) GetAllMentee(c echo.Context) error {
	results, err := handler.menteeService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data, "+err.Error()))
	}

	var menteesResponse []MenteeResponse
	for _, value := range results {
		menteesResponse = append(menteesResponse, CoreToGetMenteeResponse(value))
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read data", menteesResponse))
}

func (handler *menteeHandler) GetMenteeById(c echo.Context) error {
	paramId := c.Param("id")
	menteeId, _ := strconv.ParseUint(paramId, 10, 64)

	result, err := handler.menteeService.GetById(menteeId)
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.FailedResponse("error read data, "+err.Error()))
	}

	menteeResponse := CoreToGetMenteeResponse(result)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read data", menteeResponse))
}

func (handler *menteeHandler) UpdateMenteeById(c echo.Context) error {
	paramId := c.Param("id")
	menteeId, _ := strconv.ParseUint(paramId, 10, 64)

	menteeInput := MenteeRequest{}
	errBind := c.Bind(&menteeInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}

	menteeCore := MenteeRequestToCore(menteeInput)
	result, err := handler.menteeService.UpdateById(menteeId, menteeCore)
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.FailedResponse("error update data, "+err.Error()))
	}

	menteeResponse := CoreToGetMenteeResponse(result)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success update data", menteeResponse))
}

func (handler *menteeHandler) DeleteMenteeById(c echo.Context) error {
	paramId := c.Param("id")
	menteeId, _ := strconv.ParseUint(paramId, 10, 64)

	err := handler.menteeService.DeleteById(menteeId)
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.FailedResponse("error delete data, "+err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success delete data"))
}
