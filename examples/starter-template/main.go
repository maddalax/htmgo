package main

import (
	"github.com/labstack/echo/v4"
	"github.com/maddalax/htmgo/framework/h"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"starter-template/pages"
	"starter-template/partials/load"
	"time"
)

func main() {
	var startTime = time.Now()
	f := echo.New()

	f.Static("/public", "./assets/dist")

	load.RegisterPartials(f)
	pages.RegisterPages(f)

	log.Printf("main() ready in %s", time.Since(startTime))
	h.Start(f, h.App{
		LiveReload: true,
	})
}
