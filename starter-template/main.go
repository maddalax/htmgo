package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/maddalax/htmgo/framework/h"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"starter-template/ent"
	"starter-template/pages"
	"starter-template/partials/load"
	"time"
)

func main() {
	f := fiber.New()

	f.Static("/public", "./assets/dist")

	f.Use(func(ctx *fiber.Ctx) error {
		if ctx.Path() == "/livereload" {
			return ctx.Next()
		}
		now := time.Now()
		err := ctx.Next()
		duration := time.Since(now)
		ctx.Set("X-Response-Times", duration.String())
		// Log or print the request method, URL, and duration
		log.Printf("Requests: %s %s took %dms", ctx.Method(), ctx.OriginalURL(), duration.Milliseconds())
		return err
	})

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

	h.Start(f, h.App{
		LiveReload: true,
	})
}
