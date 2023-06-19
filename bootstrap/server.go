package bootstrap

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wecanooo/gosari/core"
	"net/http"
	"runtime"
	"strings"
)

func setupServer(fn func(router *core.Application)) {
	core.SetupLog()

	e := echo.New()
	e.Debug = core.GetConfig().IsDev()
	e.HideBanner = true

	core.NewApplication(e)

	defaultMiddleware(e)

	core.GetApplication().RegisterRoutes(fn)
	core.GetApplication().PrintRoutes(core.GetConfig().String("APP.TEMP_DIR") + "/routes.json")

	fmt.Printf("서버를 준비하는 중입니다: %s, %s\n\n", core.GetConfig().AppMode(), core.GetConfig().String("APP.SERVER"))
}

func defaultMiddleware(app *echo.Echo) {
	staticURL := core.GetConfig().String("APP.STATIC_URL")

	if !core.GetConfig().IsDev() {
		app.Use(middleware.Recover())
	}

	if core.GetConfig().IsDev() {
		app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "${status}   ${method}   ${latency_human}               ${uri}\n",
		}))
	} else {
	}

	app.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
		Getter: middleware.MethodFromForm("_method"),
	}))

	app.Pre(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))

	if core.GetConfig().Bool("APP.GZIP") {
		app.Use(middleware.GzipWithConfig(middleware.GzipConfig{
			Skipper: func(c echo.Context) bool {
				return !strings.HasPrefix(c.Request().URL.Path, staticURL)
			},
		}))
	}
}

func RunServer(fn func(router *core.Application)) {
	setupServer(fn)

	if core.GetConfig().Bool("APP.PROFILE") {
		runtime.SetBlockProfileRate(1)
		go func() {
			profileListen := fmt.Sprintf("0.0.0.0:%d", core.GetConfig().Int("APP.PROFILE_PORT"))
			err := http.ListenAndServe(profileListen, nil)
			if err != nil {
				return
			}
		}()
	}

	core.GetApplication().Echo.Logger.Fatal(core.GetApplication().Start(core.GetConfig().String("APP.ADDR")))
}
