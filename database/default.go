package database

import (
	"database/sql"
	"fmt"
	"github.com/wecanooo/gosari/core"
	"github.com/wecanooo/gosari/core/pkg/db"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDefaultDatabase() (*gorm.DB, *sql.DB) {
	adapter := core.GetConfig().DefaultString("DB.MASTER.ADAPTER", "mysql")
	dsn := db.BuildDatabaseDSN(adapter, db.DatabaseConfig{
		UserName: core.GetConfig().String("DB.MASTER.USERNAME"),
		Password: core.GetConfig().String("DB.MASTER.PASSWORD"),
		Host:     core.GetConfig().String("DB.MASTER.HOST"),
		Port:     core.GetConfig().String("DB.MASTER.PORT"),
		DBName:   core.GetConfig().String("DB.MASTER.DATABASE"),
		Options:  core.GetConfig().String("DB.MASTER.OPTIONS"),
	}, func(s string) string {
		return s + "_" + string(core.GetConfig().AppMode())
	})

	gormDB, err := gorm.Open(getDialect(adapter, dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(getGormLoggerLevel()),
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic("fail to connect database: " + err.Error() + " " + dsn)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		panic("fail to connect database: " + err.Error() + " " + dsn)
	}

	sqlDB.SetMaxOpenConns(core.GetConfig().Int("DB.MASTER.MAX_OPEN_CONNECTIONS"))
	sqlDB.SetMaxIdleConns(core.GetConfig().Int("DB.MASTER.MAX_IDLE_CONNECTIONS"))

	fmt.Println("database connected successfully")
	return gormDB, sqlDB
}

func getDialect(adapter, dsn string) gorm.Dialector {
	if adapter == "postgres" {
		return postgres.Open(dsn)
	}

	if adapter == "sqlite" {
		return sqlite.Open(dsn)
	}

	return mysql.Open(dsn)
}

func getGormLoggerLevel() logger.LogLevel {
	if core.GetConfig().IsDev() {
		return logger.Info
	}

	return logger.Silent
}
