package handler

import (
	"net/http"

	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/status"
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/helper"
	"github.com/labstack/echo/v4"
)

type StatusHandler struct {
	statusService status.StatusServiceInterface
}

func New(service status.StatusServiceInterface) *StatusHandler {
	return &StatusHandler{
		statusService: service,
	}
}

func (handler *StatusHandler) GetAllStatus(c echo.Context) error {
	results, err := handler.statusService.GetAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data "+err.Error()))
	}

	var statussResponse []StatusResponse
	for _, value := range results {
		statussResponse = append(statussResponse, StatusResponse{
			Id:   value.Id,
			Name: value.Name,
		})
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read data", statussResponse))
}
