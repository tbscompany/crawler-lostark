package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	l "gorm.io/gorm/logger"

	"crawler-lostark/pkg/config"
	"crawler-lostark/pkg/logger"
)

type DbConnection struct {
	DB *gorm.DB
}

var (
	connection DbConnection
	DB         *gorm.DB
	err        error
)

func InitDatabase() {
	switch engine := config.DatabaseEngine; engine {
	case "postgres":
		connection.DB, err = gorm.Open(postgres.Open(config.Database), &gorm.Config{Logger: l.Default.LogMode(l.Silent)})
	default:
		connection.DB, err = gorm.Open(sqlite.Open(config.Database), &gorm.Config{Logger: l.Default.LogMode(l.Silent)})
	}

	DB = connection.DB
	if err != nil {
		logger.Logger.Error().Err(err).Msg("")
	}
}

func GetDbConnection() DbConnection {
	return connection
}
