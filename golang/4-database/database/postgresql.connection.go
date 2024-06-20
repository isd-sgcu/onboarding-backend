package database

import (
	"github.com/isd-sgcu/onboarding-backend/golang/4-database/config"
	"github.com/isd-sgcu/onboarding-backend/golang/4-database/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func InitPostgresDatabase(conf *config.DatabaseConfig, isDebug bool) (db *gorm.DB, err error) {
	gormConf := &gorm.Config{}

	if !isDebug {
		gormConf.Logger = gormLogger.Default.LogMode(gormLogger.Silent)
	}

	db, err = gorm.Open(postgres.Open(conf.Url), gormConf)
	if err != nil {
		return nil, err
	}

	// add Order table with the schema from model.Order
	err = db.AutoMigrate(&model.Image{})
	if err != nil {
		return nil, err
	}

	return
}
