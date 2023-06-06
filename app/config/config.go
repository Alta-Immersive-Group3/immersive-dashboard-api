package config

import (
	"fmt"
	"os"
	"strconv"
)

var (
	SECRET_JWT = ""
)

type AppConfig struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOSTNAME string
	DB_PORT     int
	DB_NAME     string
	jwtKey      string
}

func InitConfig() *AppConfig {
	return ReadEnv()
}

func ReadEnv() *AppConfig {
	app := AppConfig{}
	isRead := true

	if val, found := os.LookupEnv("JWT_KEY"); found {
		app.jwtKey = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBUSER"); found {
		app.DB_USERNAME = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPASS"); found {
		app.DB_PASSWORD = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBHOST"); found {
		app.DB_HOSTNAME = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPORT"); found {
		cnv, _ := strconv.Atoi(val)
		app.DB_PORT = cnv
		isRead = false
	}
	if val, found := os.LookupEnv("DBNAME"); found {
		app.DB_NAME = val
		isRead = false
	}

	if isRead {
		app.jwtKey = os.Getenv("JWT_KEY")
		app.DB_USERNAME = os.Getenv("DBUSER")
		app.DB_PASSWORD = os.Getenv("DBPASS")
		app.DB_HOSTNAME = os.Getenv("DBHOST")
		app.DB_PORT, _ = strconv.Atoi(os.Getenv("DBPORT"))
		app.DB_NAME = os.Getenv("DBNAME")
	}

	SECRET_JWT = app.jwtKey
	fmt.Println("check", app.jwtKey)
	return &app
}
