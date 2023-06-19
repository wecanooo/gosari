package bootstrap

import (
	"github.com/labstack/gommon/log"
	"github.com/wecanooo/gosari/config"
	"github.com/wecanooo/gosari/core"
	"github.com/wecanooo/gosari/core/pkg/jwt"
)

func SetupConfig(configFilePath, configFileType string) {
	config.Setup(configFilePath, configFileType)
	core.NewAppConfig()
	err := config.WriteConfig(core.GetConfig().String("APP.TEMP_DIR") + "/config.json")
	if err != nil {
		log.Infof("설정파일 정보를 저장하는 중 오류가 발생하였습니다. %v\n", err)
	}

	jwt.Setup(&jwt.Config{
		SecretKey:            core.GetConfig().String("APP.KEY"),
		AccessTokenLifeTime:  core.GetConfig().Duration("TOKEN.ACCESS_TOKEN_LIFETIME"),
		RefreshTokenLifeTime: core.GetConfig().Duration("TOKEN.REFRESH_TOKEN_LIFETIME"),
	})
}
