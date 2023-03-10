package main

import (
	"github.com/aluka-7/api-base/app/controller"
	"github.com/aluka-7/api-base/wire"
	"github.com/aluka-7/cache"
	_ "github.com/aluka-7/cache-redis"
	"github.com/aluka-7/configuration"
	"github.com/aluka-7/datasource"
	"github.com/aluka-7/gomongo"
	"github.com/aluka-7/web"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func App(conf configuration.Configuration) {
	web.App(func(eng *echo.Echo) {
		ce := cache.Engine(wire.SystemId, conf)
		orm := datasource.Engine(conf, wire.SystemId).Orm("")
		gmc := gomongo.Engine(conf, wire.SystemId).Connection("")
		cs := wire.InitializeCompanyService(orm, gmc, ce)

		group := eng.Group("/api/v1")
		controller.NewCompanyController(group, "/company", cs)
	}, wire.SystemId, conf)
}

func main() {
	conf := configuration.DefaultEngine()
	App(conf)
}
