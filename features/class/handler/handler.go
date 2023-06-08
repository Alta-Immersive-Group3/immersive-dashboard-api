package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/class"
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/helper"
	"github.com/labstack/echo/v4"
)

type ClassHandler struct {
	classService class.ClassServiceInterface
}

func New(service class.ClassServiceInterface) *ClassHandler {
	return &ClassHandler{
		classService: service,
	}
}

func (handler *ClassHandler) CreateClass(c echo.Context) error {
	classInput := ClassRequest{}
	errBind := c.Bind(&classInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}

	classCore := classRequestToCore(classInput)
	result, err := handler.classService.Create(classCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error insert data, "+err.Error()))
		}
	}

	classResponse := CoreToGetClassResponse(result)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read data", classResponse))
}

func (handler *ClassHandler) GetAllClass(c echo.Context) error {
	results, err := handler.classService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data, "+err.Error()))
	}

	var classsResponse []ClassResponse
	for _, value := range results {
		classsResponse = append(classsResponse, CoreToGetClassResponse(value))
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read data", classsResponse))
}

func (handler *ClassHandler) GetClassById(c echo.Context) error {
	paramId := c.Param("id")
	classId, _ := strconv.ParseUint(paramId, 10, 64)

	result, err := handler.classService.GetById(classId)
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.FailedResponse("error read data, "+err.Error()))
	}

	classResponse := CoreToGetClassResponse(result)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read data", classResponse))
}

func (handler *ClassHandler) UpdateClassById(c echo.Context) error {
	paramId := c.Param("id")
	classId, _ := strconv.ParseUint(paramId, 10, 64)

	classInput := ClassRequest{}
	errBind := c.Bind(&classInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}

	classCore := classRequestToCore(classInput)
	result, err := handler.classService.UpdateById(classId, classCore)
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.FailedResponse("error update data, "+err.Error()))
	}

	classResponse := CoreToGetClassResponse(result)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success update data", classResponse))
}

func (handler *ClassHandler) DeleteClassById(c echo.Context) error {
	paramId := c.Param("id")
	classId, _ := strconv.ParseUint(paramId, 10, 64)

	err := handler.classService.DeleteById(classId)
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.FailedResponse("error delete data, "+err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success delete data"))
}
