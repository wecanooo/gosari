package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"strings"
)

func Setup(configFilePath, configFileType string) {
	viper.SetConfigFile(configFilePath)
	viper.SetConfigType(configFileType)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("설정파일을 찾지 못했습니다. 설정파일이 존재하는지 확인해 주세요. %v\n", err))
	}

	setupDefaultConfig()

	viper.AutomaticEnv()
	viper.SetEnvPrefix(viper.GetString("APP.NAME"))
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	watchConfig()

	fmt.Printf("설정파일이 로드되었습니다: %s\n", configFilePath)
}

func WriteConfig(filename string) error {
	return viper.WriteConfigAs(filename)
}

func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(event fsnotify.Event) {
		log.Infof("설정파일이 변경되었습니다: %s\n", event.Name)
	})
}
