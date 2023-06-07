package router

import (
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/app/middlewares"
	_teamData "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/team/data"
	_teamHandler "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/team/handler"
	_teamService "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/team/service"
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
	e.POST("/users", userHandlerAPI.CreateUser, middlewares.JWTMiddleware())
	e.GET("/users", userHandlerAPI.GetAllUser, middlewares.JWTMiddleware())
	e.PUT("/users", userHandlerAPI.UpdateUserProfile, middlewares.JWTMiddleware())
	e.PUT("/users/:id", userHandlerAPI.UpdateUserById, middlewares.JWTMiddleware())

	teamData := _teamData.New(db)
	teamService := _teamService.New(teamData)
	teamHandlerAPI := _teamHandler.New(teamService)

	e.GET("/teams", teamHandlerAPI.GetAllTeam, middlewares.JWTMiddleware())
}
