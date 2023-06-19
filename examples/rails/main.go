package main

import (
	"github.com/wecanooo/gosari/bootstrap"
	"github.com/wecanooo/gosari/examples/rails/routes"
)

func main() {
	bootstrap.SetupConfig("examples/config/development.yaml", "yaml")
	bootstrap.SetupDB()
	bootstrap.SetupServer(routes.Register)
	bootstrap.RunServer()
}
