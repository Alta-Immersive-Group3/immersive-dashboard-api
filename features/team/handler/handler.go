package handler

import (
	"net/http"

	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/team"
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/helper"
	"github.com/labstack/echo/v4"
)

type TeamHandler struct {
	teamService team.TeamServiceInterface
}

func New(service team.TeamServiceInterface) *TeamHandler {
	return &TeamHandler{
		teamService: service,
	}
}

func (handler *TeamHandler) GetAllTeam(c echo.Context) error {
	results, err := handler.teamService.GetAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	var teamsResponse []TeamResponse
	for _, value := range results {
		teamsResponse = append(teamsResponse, TeamResponse{
			Id:   value.Id,
			Name: value.Name,
		})
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read data", teamsResponse))
}
