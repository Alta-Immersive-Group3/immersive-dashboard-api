package handler

import (
	"net/http"
	"strings"

	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/user"
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/helper"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.UserServiceInterface
}

func New(service user.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (handler *UserHandler) Login(c echo.Context) error {
	loginInput := AuthRequest{}
	errBind := c.Bind(&loginInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}

	dataLogin, token, err := handler.userService.Login(loginInput.Email, loginInput.Password)
	if err != nil {
		if strings.Contains(err.Error(), "login failed") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error login, internal server error"))
		}
	}

	response := CoreToAuthResponse(dataLogin, token)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("login successful", response))
}

func (handler *UserHandler) CreateUser(c echo.Context) error {
	userInput := UserRequest{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}

	if !isAdmin(userInput.Role) {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("error, unauthorized"))
	}

	userCore := UserRequestToCore(userInput)
	err := handler.userService.Create(userCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error insert data"+err.Error()))
		}
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success insert data"))
}

func isAdmin(role string) bool {
	return role == "admin"
}
