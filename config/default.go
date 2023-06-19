package config

import (
	"time"

	"github.com/spf13/viper"
)

const (
	defaultTempDir = "tmp"
	defaultAppPort = "3000"
	defaultAppName = "gosari"
)

var defaultConfigMap = map[string]interface{}{
	// application 기본설정
	"APP.NAME":       defaultAppName,
	"APP.VERSION":    "1.0.0",
	"APP.RUNMODE":    "production",
	"APP.ADDR":       ":" + defaultAppPort,
	"APP.URL":        "http://localhost:" + defaultAppPort,
	"APP.KEY":        "Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5",
	"APP.TEMP_DIR":   defaultTempDir,
	"APP.UPLOAD_DIR": "public/uploads",

	// database 기본설정
	"DB.MASTER.ADAPTER":              "mysql",
	"DB.MASTER.HOST":                 "127.0.0.1",
	"DB.MASTER.PORT":                 "3306",
	"DB.MASTER.DATABASE":             defaultAppName,
	"DB.MASTER.USERNAME":             "root",
	"DB.MASTER.PASSWORD":             "secret",
	"DB.MASTER.OPTIONS":              "charset=utf8mb4&parseTime=True&loc=Local",
	"DB.MASTER.MAX_OPEN_CONNECTIONS": 100,
	"DB.MASTER.MAX_IDLE_CONNECTIONS": 90,
	"DB.MASTER.AUTO_MIGRATE":         false,

	// redis 기본설정
	"REDIS.HOST":     "127.0.0.1",
	"REDIS.PASSWORD": "",
	"REDIS.DATABASE": "0",

	// jwt-token 기본설정
	"TOKEN.ACCESS_TOKEN_LIFETIME":  60 * time.Minute,
	"TOKEN.REFRESH_TOKEN_LIFETIME": 60 * 24 * time.Minute,

	// log 기본설정
	"LOG.PREFIX":     "[ZAP_LOGGER]",
	"LOG.FOLDER":     defaultTempDir + "/logs/zap",
	"LOG.LEVEL":      "debug", // debug, info, warn, error, dpanic, panic, fatal
	"LOG.MAXSIZE":    10,
	"LOG.MAXBACKUPS": 5,
	"LOG.MAXAGES":    30,
}

func setupDefaultConfig() {
	for k, v := range defaultConfigMap {
		viper.SetDefault(k, v)
	}
}
