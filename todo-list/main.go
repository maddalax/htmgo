package main

import (
	"github.com/labstack/echo/v4"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
	_ "github.com/mattn/go-sqlite3"
	"todolist/ent"
	"todolist/infrastructure/db"
	"todolist/pages"
	"todolist/partials/load"
)

func main() {
	locator := service.NewLocator()

	service.Set[ent.Client](locator, service.Singleton, func() *ent.Client {
		return db.Provide()
	})

	h.Start(h.AppOpts{
		ServiceLocator: locator,
		LiveReload:     true,
		Register: func(e *echo.Echo) {
			e.Static("/public", "./assets/dist")
			load.RegisterPartials(e)
			pages.RegisterPages(e)
		},
	})
}
