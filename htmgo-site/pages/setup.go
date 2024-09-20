package pages

import (
	"github.com/labstack/echo/v4"
	"htmgo-site/internal/dirwalk"
)

func RegisterMarkdown(app *echo.Echo, dir string, handler func(ctx echo.Context, path string) error) {
	for _, page := range dirwalk.WalkPages(dir) {
		app.GET(page.RoutePath, func(ctx echo.Context) error {
			return handler(ctx, page.FilePath)
		})
	}
}
