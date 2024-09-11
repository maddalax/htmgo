package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"mhtml/h"
	"mhtml/pages"
	"mhtml/partials"
)

func main() {
	f := fiber.New()
	f.Static("/js", "./js")

	f.Use(func(ctx *fiber.Ctx) error {
		if ctx.Cookies("mhtml-session") != "" {
			return ctx.Next()
		}
		id := ctx.IP() + uuid.NewString()
		ctx.Cookie(&fiber.Cookie{
			Name:        "mhtml-session",
			Value:       id,
			SessionOnly: true,
		})
		return ctx.Next()
	})

	f.Get("/mhtml/partials.*", func(ctx *fiber.Ctx) error {
		return h.PartialView(ctx, partials.GetPartialFromContext(ctx))
	})

	pages.RegisterPages(f)

	h.Start(f, h.App{
		LiveReload: true,
	})
}
