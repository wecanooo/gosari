package core

import (
	"database/sql"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type GormConnection struct {
	Engine *gorm.DB
	DB     *sql.DB
}

func NewDatabaseConnection(e *gorm.DB, d *sql.DB) {
	databaseConnection = &GormConnection{
		Engine: e,
		DB:     d,
	}
}

func NewRedisConnection(r redis.Cmdable) {
	redisClient = r
}
