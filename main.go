package main

import (
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/app/config"
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/app/database"
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/app/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	cfg := config.InitConfig()
	dbMysql := database.InitDBMysql(cfg)

	database.InitialMigration(dbMysql)

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	router.InitRouter(dbMysql, e)

	e.Logger.Fatal(e.Start(":8080"))
}
