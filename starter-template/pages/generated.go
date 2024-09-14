// Package pages THIS FILE IS GENERATED. DO NOT EDIT.
package pages

import "github.com/gofiber/fiber/v2"
import "github.com/maddalax/htmgo/framework/h"

func RegisterPages(f *fiber.App) {
	f.Get("/", func(ctx *fiber.Ctx) error {
		return h.HtmlView(ctx, IndexPage(ctx))
	})
}
