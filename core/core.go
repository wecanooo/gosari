package core

import (
	"database/sql"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var (
	application        *Application
	databaseConnection *GormConnection
	appConfig          *AppConfig
	redisClient        redis.Cmdable
	appLog             *zap.SugaredLogger
)

func GetApplication() *Application {
	if application == nil {
		panic("application is not valid")
	}
	return application
}

func GetConfig() *AppConfig {
	if appConfig == nil {
		panic("config is not valid")
	}

	return appConfig
}

func GetDatabaseConnection() *sql.DB {
	if databaseConnection == nil || databaseConnection.DB == nil {
		panic("database connection is not valid")
	}
	return databaseConnection.DB
}

func GetRedis() redis.Cmdable {
	if redisClient == nil {
		panic("redis connection is not valid")
	}
	return redisClient
}

func GetLog() *zap.SugaredLogger {
	if appLog == nil {
		panic("log connection is not valid")
	}

	return appLog
}
