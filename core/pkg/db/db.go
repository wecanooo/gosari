package db

import (
	"fmt"
	"log"
)

type DatabaseConfig struct {
	UserName string
	Password string
	Host     string
	Port     string
	DBName   string
	Options  string
}

func BuildDatabaseDSN(adapter string, config DatabaseConfig, buildDBName func(string) string) string {
	if buildDBName != nil {
		config.DBName = buildDBName(config.DBName)
	}

	switch adapter {
	case "mysql":
		return fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?%s",
			config.UserName,
			config.Password,
			config.Host,
			config.Port,
			config.DBName,
			config.Options,
		)
	case "postgres":
		return fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s options='%s'",
			config.Host,
			config.Port,
			config.UserName,
			config.DBName,
			config.Password,
			config.Options,
		)
	case "sqlite3":
		return config.DBName
	}

	log.Panicf("지원하지 않는 데이터베이스 타입입니다: %s", adapter)
	return ""
}
