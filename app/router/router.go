package router

import (
	_userData "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/user/data"
	_userHandler "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/user/handler"
	_userService "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userData := _userData.New(db)
	userService := _userService.New(userData)
	userHandlerAPI := _userHandler.New(userService)

	e.POST("/login", userHandlerAPI.Login)
}
