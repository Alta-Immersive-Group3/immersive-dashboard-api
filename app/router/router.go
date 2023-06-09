package router

import (
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/app/middlewares"
	_classData "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/class/data"
	_classHandler "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/class/handler"
	_classService "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/class/service"
	_feedbackData "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/feedback/data"
	_feedbackHandler "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/feedback/handler"
	_feedbackService "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/feedback/service"
	_menteeData "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/mentee/data"
	_menteeHandler "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/mentee/handler"
	_menteeService "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/mentee/service"
	_statusData "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/status/data"
	_statusHandler "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/status/handler"
	_statusService "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/status/service"
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
	e.GET("/users/:id", userHandlerAPI.GetUserById, middlewares.JWTMiddleware())
	e.PUT("/users", userHandlerAPI.UpdateUserProfile, middlewares.JWTMiddleware())
	e.PUT("/users/:id", userHandlerAPI.UpdateUserById, middlewares.JWTMiddleware())
	e.DELETE("/users/:id", userHandlerAPI.DeleteUserById, middlewares.JWTMiddleware())

	teamData := _teamData.New(db)
	teamService := _teamService.New(teamData)
	teamHandlerAPI := _teamHandler.New(teamService)

	e.GET("/teams", teamHandlerAPI.GetAllTeam, middlewares.JWTMiddleware())

	classData := _classData.New(db)
	classService := _classService.New(classData)
	classHandlerAPI := _classHandler.New(classService)

	e.POST("/classes", classHandlerAPI.CreateClass, middlewares.JWTMiddleware())
	e.GET("/classes", classHandlerAPI.GetAllClass, middlewares.JWTMiddleware())
	e.GET("/classes/:id", classHandlerAPI.GetClassById, middlewares.JWTMiddleware())
	e.PUT("/classes/:id", classHandlerAPI.UpdateClassById, middlewares.JWTMiddleware())
	e.DELETE("/classes/:id", classHandlerAPI.DeleteClassById, middlewares.JWTMiddleware())

	statusData := _statusData.New(db)
	statusService := _statusService.New(statusData)
	statusHandlerAPI := _statusHandler.New(statusService)

	e.GET("/statuses", statusHandlerAPI.GetAllStatus, middlewares.JWTMiddleware())

	menteeData := _menteeData.New(db)
	menteeService := _menteeService.New(menteeData)
	menteeHandlerAPI := _menteeHandler.New(menteeService)

	e.POST("/mentees", menteeHandlerAPI.CreateMentee, middlewares.JWTMiddleware())
	e.GET("/mentees", menteeHandlerAPI.GetAllMentee, middlewares.JWTMiddleware())
	e.GET("/mentees/:id", menteeHandlerAPI.GetMenteeById, middlewares.JWTMiddleware())
	e.PUT("/mentees/:id", menteeHandlerAPI.UpdateMenteeById, middlewares.JWTMiddleware())
	e.DELETE("/mentees/:id", menteeHandlerAPI.DeleteMenteeById, middlewares.JWTMiddleware())

	feedbackData := _feedbackData.New(db)
	feedbackService := _feedbackService.New(feedbackData)
	feedbackHandlerAPI := _feedbackHandler.New(feedbackService)

	e.POST("/feedbacks", feedbackHandlerAPI.CreateFeedback, middlewares.JWTMiddleware())
	e.GET("/feedbacks", feedbackHandlerAPI.GetAllFeedback, middlewares.JWTMiddleware())
	e.GET("/feedbacks/:id", feedbackHandlerAPI.GetFeedbackById, middlewares.JWTMiddleware())
	e.PUT("/feedbacks/:id", feedbackHandlerAPI.UpdateFeedbackById, middlewares.JWTMiddleware())
	e.DELETE("/feedbacks/:id", feedbackHandlerAPI.DeleteFeedbackById, middlewares.JWTMiddleware())
	e.GET("/mentees/:id/feedback", feedbackHandlerAPI.GetAllMenteeFeedback, middlewares.JWTMiddleware())
}
