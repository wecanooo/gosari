package routes

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/wecanooo/gosari/core"
	"github.com/wecanooo/gosari/examples/rails/app/controllers"
)

const (
	APIPrefix = "/api"
)

func registerApi(router *core.Application) {
	e := router.Group(APIPrefix, middleware.CORS())
	user := e.Group("/user")
	{
		uc := controllers.NewUserController()
		router.RegisterHandler(user.GET, "", uc.Index).Name = "user.index"
		router.RegisterHandler(user.GET, ":id", uc.Show).Name = "user.show"
		router.RegisterHandler(user.POST, "", uc.Create).Name = "user.create"
	}
}
