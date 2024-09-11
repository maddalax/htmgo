package h

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

var Version = time.Now().Nanosecond()

func LiveReloadHandler(c *fiber.Ctx) error {
	v := strconv.FormatInt(int64(Version), 10)
	current := c.Cookies("version", v)

	if current != v {
		c.Set("HX-Refresh", "true")
	}

	c.Cookie(&fiber.Cookie{
		Name:  "version",
		Value: v,
	})

	return c.SendString("")
}

func LiveReload() *Node {
	return Div(Get("/livereload"), Trigger("every 200ms"))
}

func AddLiveReloadHandler(path string, app *fiber.App) {
	app.Get(path, LiveReloadHandler)
}
