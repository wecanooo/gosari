package core

import (
	"github.com/spf13/viper"
	"time"
)

type AppConfig struct{}

func NewAppConfig() {
	appConfig = &AppConfig{}
}

func (*AppConfig) String(key string) string {
	return viper.GetString(key)
}

func (*AppConfig) DefaultString(key, defaultVal string) string {
	v := viper.GetString(key)
	if v == "" {
		return defaultVal
	}
	return v
}

func (*AppConfig) Int(key string) int {
	return viper.GetInt(key)
}

func (*AppConfig) Bool(key string) bool {
	return viper.GetBool(key)
}

func (*AppConfig) Duration(key string) time.Duration {
	return viper.GetDuration(key)
}

func (c *AppConfig) IsDev() bool {
	return c.AppMode() == "development"
}

func (c *AppConfig) AppMode() string {
	return c.DefaultString("APP.MODE", "development")
}
