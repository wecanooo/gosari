# gosari :seedling:

:seedling: [gosari](https://github.com/wecanooo/gosari) 프로젝트는 [go](https://golang.org/) 언어를 이용하여 웹서비스를 제작함에 있어 DRY(Don't Repeat Yourself) 원칙에 따라 반복되는 프로젝트 구성 및 설정을 간편하게 제공하고자 제작된 라이브러리입니다.

## Requirements

- Go (1.19 or later)
- Editor : Goland (recommended), Visual Studio Code

## 기술구성

- Web Framework: [Echo](https://echo.labstack.com/)
- ORM: [Gorm](https://gorm.io/index.html)
- Logging: [Zap](https://github.com/uber-go/zap)
- Config: [Viper](https://github.com/spf13/viper)
- Redis: [Go-Redis](github.com/redis/go-redis/v9)

## 사용방법

### 환경구성

[gosari](https://github.com/wecanooo/gosari) 를 이용하기 위해서는 Go 1.19.x 버전이 설치되어 있어야 합니다.

```shell
$ go version

go version go1.19.3 darwin/amd64
```

### 프로젝트 생성

편의상 [GoLand](https://www.jetbrains.com/go/) 를 이용하는 것으로 합니다. Goland 를 실행 후 `New Project` 를 실행한 화면에서 `Location` 을 지정합니다. `Location` 에 입력되는 project 명을 편의상 `godori` 라고 입력하도록 하겠습니다.

이제, [gosari](https://github.com/wecanooo/gosari) 라이브러리를 다운로드 받습니다. (원하는 버전이 별도로 있을 경우 버전명을 붙여서 다운로드 받습니다.)

```shell
$ go get github.com/wecanooo/gosari@v1.1.4

go: downloading github.com/wecanooo/gosari v1.1.4
go get: added github.com/wecanooo/gosari v1.1.4
```

여기에서는 Cli 어플리케이션을 편리하게 만들기 위해 [cobra](https://github.com/spf13/cobra) 를 이용하도록 하겠습니다. 먼저 Cobra Cli 를 사용하기 위해 다음과 같이 Cli Tool 을 다운로드 받습니다.

```shell
$ go get -u https://github.com/spf13/cobra/cobra
```

정상적으로 cobra cli 가 설치되었다면, repository root 로 이동하여 아래의 명령을 실행하여 cli 초기 구성을 진행합니다.

> 주의! : 실행하는 위치가 project root 가 아닌 repository root 입니다. (ex, github.com/wecanooo)

```shell
$ cobra init godori --pkg-name cmd

Your Cobra application is ready at
/Users/dante/go/src/github.com/wecanooo/godori
```

위의 명령을 수행하면 project root 내 `cmd` 폴더와 `main.go` 파일이 생성됩니다.
`cmd` 폴더 내에 생성되는 파일과 `main.go` 파일은 [cobra document](https://github.com/spf13/cobra) 를 참고하시기 바랍니다.

`main.go` 파일을 열고 다음과 같이 변경합니다.

```go
package main

import "github.com/wecanooo/godori/cmd"

func main() {
	cmd.Execute()
}
```

위와 같이 변경한 뒤 `cmd/root.go` 파일을 다음과 같이 변경합니다.

```go
package cmd

import (
	"fmt"
	"os"

	"github.com/wecanooo/gosari/bootstrap"
	"github.com/spf13/cobra"
)

const (
	defaultConfigFilePath = "config/development.yaml"
	configFileType        = "yaml"
)

var configFilePath string

var rootCmd = &cobra.Command{
	Use: "gosari",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&configFilePath, "config", "c", defaultConfigFilePath, "config file (default is config/development.yaml)")
}

func initConfig() {
	if configFilePath == "" {
		configFilePath = defaultConfigFilePath
	}
	bootstrap.SetupConfig(configFilePath, configFileType)
}
```

이제 마지막 설정으로 gosari 를 이용한 web server 를 실행하기 위해 `cmd/server.go` 파일을 생성하고 아래의 내용을 넣습니다.

```go
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/wecanooo/godori/routes"
	"github.com/wecanooo/gosari/bootstrap"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "run echo server",
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.SetupDB()
		//bootstrap.SetupRedis()        # redis 를 사용하려면 주석을 해제하세요.

		bootstrap.SetupServer(routes.Register)
		bootstrap.RunServer()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
```

여기까지 완료했다면 초기 구성은 완성되었습니다. 이제 application 구조를 잡아가는 것을 설명하도록 하겠습니다.

### Config 구성

gosari 라이브러리가 가지고 있는 기본 설정은 다음과 같습니다.

```go
// gosari/config/default.go

const (
  defaultTempDir = "tmp"
  defaultAppPort = "3000"
  defaultAppName = "gosari"
)

// application 기본설정
var defaultConfigMap = map[string]interface{}{
"APP.NAME":       defaultAppName,
"APP.VERSION":    "1.0.0",
"APP.RUNMODE":    "production",
"APP.ADDR":       ":" + defaultAppPort,
"APP.URL":        "http://localhost:" + defaultAppPort,
"APP.KEY":        "USER_DEFINED_KEY",
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
"DB.MASTER.AUTO_MIGRATE":         true,

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

```

위 설정들은 default 설정이며, 만약 설정항목을 변경하고 싶다면 `config` 폴더 내에 `development.yaml` 또는 `production.yaml`, `stage.yaml` 파일을 넣어서 항목을 변경할 수 있습니다. 아래에 예시를 참고하세요.


```yaml
APP:
  NAME: godori
  RUNMODE: development
  SERVER: :3000
  KEY: USER_DEFINED_KEY
  PROFILE: false
  PROFILE_PORT: 4000
  TEMP_DIR: tmp

DB:
  MASTER:
    ADAPTER: mysql
    HOST: 127.0.0.1
    PORT: 3306
    DATABASE: gosari
    USERNAME: root
    PASSWORD: top_secret

REDIS:
  HOST: localhost:6379
  PASSWORD:
  DATABASE:

LOG:
  FOLDER:
  PREFIX:
  MAXSIZE:
  MAXAGES:
```

### Route 등록

`cmd/server.go` 파일을 다시 한번 살펴보겠습니다.

```go
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/wecanooo/godori/routes"
	"github.com/wecanooo/gosari/bootstrap"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "run echo server",
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.SetupDB()
		//bootstrap.SetupRedis()

		bootstrap.RunServer(routes.Register)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

```

위에서 눈여겨 볼 부분은 `bootstrap.RunServer(routes.Register)` 구문입니다. Ruby on Rails 의 `routes.rb` 파일과 같은 역할을 할 수 있는 파일이 필요하며, `routes/routes.go` 파일에 생성하도록 하겠습니다.

```go
package routes

import (
	"github.com/wecanooo/godori/app/controllers/api"
	"github.com/wecanooo/godori/app/repositories"
	"github.com/wecanooo/godori/app/services"
	"github.com/wecanooo/godori/routes/wrapper"
	"github.com/wecanooo/gosari/core"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)


const (
	APIPrefix = "/api"
)

func Register(app *core.Application) {
	if core.GetConfig().IsDev() {
		app.GET("/api-doc/*", echoSwagger.WrapHandler).Name = "api-doc"
	}

	// 여기 아래부터는 직접 코드를 작성하는 구간입니다.
	ur := repositories.NewUserRepository(core.GetDatabaseEngine())
	us := services.NewUserServices(ur)

	e := app.Group(APIPrefix, middleware.CORS())

	auth := e.Group("/auth")
	{
		tc := api.NewAuthController(us)
		app.RegisterHandler(auth.POST, "/token", tc.CreateToken).Name = "token.create"
		app.RegisterHandler(auth.PUT, "/refresh", wrapper.TokenAuth(tc.RefreshToken)).Name = "token.refresh"
	}

	users := e.Group("/users")
	{
		uc := api.NewUserController(us)
		app.RegisterHandler(users.POST, "", uc.Create).Name = "users.create"
		app.RegisterHandler(users.GET, "", uc.Index).Name = "users.index"
	}
}

```

위와 같이 route 를 지정한 뒤 `bootstrap.RunServer(routes.Register)` 를 하면 서버가 정상적으로 실행이 됩니다.