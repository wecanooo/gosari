package core

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/wecanooo/gosari/core/context"
	"github.com/wecanooo/gosari/core/errno"
	"net/http"
	"os"
	"strings"
)

type Application struct {
	*echo.Echo
}

func NewApplication(e *echo.Echo) {
	application = &Application{
		Echo: e,
	}
}

func (app *Application) RoutePath(name string, params ...interface{}) string {
	return app.Reverse(name, params...)
}

func (app *Application) PrintRoutes(filename string) {
	routes := make([]*echo.Route, 0)
	for _, item := range app.Routes() {
		if strings.HasPrefix(item.Name, "github.com") || strings.HasSuffix(item.Name, "notFoundHandler") {
			continue
		}
		routes = append(routes, item)
	}

	routesStr, _ := json.MarshalIndent(struct {
		Count  int           `json:"count"`
		Routes []*echo.Route `json:"routes"`
	}{
		Count:  len(routes),
		Routes: routes,
	}, "", " ")

	err := os.WriteFile(filename, routesStr, 0644)
	if err != nil {
		log.Infof("라우트 정보를 파일에 쓰는 중 오류가 발생되었습니다. %s, %v\n", filename, err)
	}
}

func (app *Application) RegisterRoutes(register func(*Application)) {
	app.Use(func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &context.AppContext{Context: c}
			return hf(cc)
		}
	})
	register(app)
}

func (*Application) RegisterHandler(fn context.EchoRegisterFunc, path string, h context.AppHandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	if path != "" && !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	return fn(path, func(c echo.Context) error {
		cc, ok := c.(*context.AppContext)
		if !ok {
			cc = context.NewAppContext(c)
			return h(cc)
		}
		return h(cc)
	}, m...)
}

func (app *Application) RegisterErrorHandler() {
	echo.NotFoundHandler = notFoundHandler
	echo.MethodNotAllowedHandler = notFoundHandler

	app.HTTPErrorHandler = func(err error, c echo.Context) {
		errnoData := transformErrorType(err)

		if !c.Response().Committed {
			if c.Request().Method == http.MethodHead {
				err = c.NoContent(http.StatusOK)
			} else {
				cc := context.NewAppContext(c)
				err = cc.ErrorJSON(errnoData)
			}
			if err != nil {
				log.Printf("routes/error#HTTPErrorHandler: %s", err)
			}
		}
	}
}

func transformErrorType(err error) *errno.Errno {
	switch typed := err.(type) {
	case *errno.Errno:
		return typed
	default:
		return errno.UnknownErr.WithErr(typed).(*errno.Errno)
	}
}

func notFoundHandler(c echo.Context) error {
	return errno.NotFoundErr
}
