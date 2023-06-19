package context

import (
	"github.com/labstack/echo/v4"
)

type AppContext struct {
	echo.Context
}

type (
	AppHandlerFunc   = func(c *AppContext) error
	EchoRegisterFunc = func(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
)

func NewAppContext(c echo.Context) *AppContext {
	return &AppContext{
		Context: c,
	}
}
