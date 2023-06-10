package database

import (
	"fmt"

	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/app/config"
	_classData "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/class/data"
	_feedbackData "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/feedback/data"
	_menteeData "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/mentee/data"
	_statusData "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/status/data"
	_teamData "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/team/data"
	_userData "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/user/data"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBMysql(cfg *config.AppConfig) *gorm.DB {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOSTNAME, cfg.DB_PORT, cfg.DB_NAME)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&_teamData.Team{}, &_userData.User{}, &_classData.Class{}, &_statusData.Status{}, &_menteeData.Mentee{}, &_feedbackData.Feedback{})
}
