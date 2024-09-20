package pages

import (
	"github.com/labstack/echo/v4"
	"htmgo-site/internal/dirwalk"
	"io/fs"
)

func RegisterMarkdown(app *echo.Echo, dir string, system fs.FS, handler func(ctx echo.Context, path string) error) {
	for _, page := range dirwalk.WalkPages(dir, system) {
		app.GET(page.RoutePath, func(ctx echo.Context) error {
			return handler(ctx, page.FilePath)
		})
	}
}
