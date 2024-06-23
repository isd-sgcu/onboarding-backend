package database

import (
	"github.com/isd-sgcu/onboarding-backend/golang/10-unit-test/config"
	"github.com/isd-sgcu/onboarding-backend/golang/10-unit-test/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func InitDatabase(conf *config.DatabaseConfig, isDebug bool) (db *gorm.DB, err error) {
	gormConf := &gorm.Config{}

	if !isDebug {
		gormConf.Logger = gormLogger.Default.LogMode(gormLogger.Silent)
	}

	db, err = gorm.Open(postgres.Open(conf.Url), gormConf)
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return nil, err
	}

	return
}
