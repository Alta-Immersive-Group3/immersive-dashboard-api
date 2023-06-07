package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/app/middlewares"
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
		} else if strings.Contains(err.Error(), "validation") {
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

func (handler *UserHandler) GetAllUser(c echo.Context) error {
	results, err := handler.userService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data "+err.Error()))
	}

	var usersResponse []UserResponse
	for _, value := range results {
		usersResponse = append(usersResponse, CoreToGetUserResponse(value))
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read data", usersResponse))
}

func (handler *UserHandler) GetUserById(c echo.Context) error {
	paramId := c.Param("id")
	userId, _ := strconv.ParseUint(paramId, 10, 64)

	result, err := handler.userService.GetById(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data "+err.Error()))
	}

	userResponse := CoreToGetUserResponse(result)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read data", userResponse))
}

func (handler *UserHandler) UpdateUserProfile(c echo.Context) error {
	paramId := middlewares.ExtractTokenUserId(c)
	userId := uint64(paramId)

	userInput := UserRequest{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}

	userCore := UserRequestToCore(userInput)
	err := handler.userService.UpdateById(userId, userCore)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error update data "+err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success update data"))
}

func (handler *UserHandler) UpdateUserById(c echo.Context) error {
	paramId := c.Param("id")
	userId, _ := strconv.ParseUint(paramId, 10, 64)

	userInput := UserRequest{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}

	userCore := UserRequestToCore(userInput)
	err := handler.userService.UpdateById(userId, userCore)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error update data "+err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success update data"))
}

func (handler *UserHandler) DeleteUserById(c echo.Context) error {
	paramId := c.Param("id")
	userId, _ := strconv.ParseUint(paramId, 10, 64)

	err := handler.userService.DeleteById(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error delete data "+err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success delete data"))
}
