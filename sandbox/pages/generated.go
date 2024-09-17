// Package pages THIS FILE IS GENERATED. DO NOT EDIT.
package pages

import "github.com/labstack/echo/v4"
import "github.com/maddalax/htmgo/framework/h"

func RegisterPages(f *echo.Echo) {
	f.Get("/", func(ctx echo.Context) error {
		return h.HtmlView(ctx, IndexPage(ctx))
	})
	f.Get("/news/:id", func(ctx echo.Context) error {
		return h.HtmlView(ctx, Test(ctx))
	})
	f.Get("/news", func(ctx echo.Context) error {
		return h.HtmlView(ctx, ListPage(ctx))
	})
	f.Get("/patients", func(ctx echo.Context) error {
		return h.HtmlView(ctx, PatientsIndex(ctx))
	})
}
