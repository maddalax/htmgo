package h

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
	"time"
)

var Version = uuid.NewString()

func handler(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		_ = websocket.Message.Send(ws, Version)
		// keep ws alive
		for {
			err := websocket.Message.Send(ws, Version)
			if err != nil {
				return
			}
			time.Sleep(500 * time.Millisecond)
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func AddLiveReloadHandler(path string, app *echo.Echo) {
	app.GET(path, handler)
}
