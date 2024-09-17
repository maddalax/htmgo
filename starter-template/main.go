package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/maddalax/htmgo/framework/h"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"starter-template/ent"
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

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed schema resources: %v", err)
	}

	log.Printf("main() ready in %s", time.Since(startTime))
	h.Start(f, h.App{
		LiveReload: true,
	})
}
