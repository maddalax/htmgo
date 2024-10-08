package main

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
	"io/fs"
	"net/http"
	"sse-with-state/__htmgo"
	"sse-with-state/event"
	"sse-with-state/sse"
)

func main() {
	locator := service.NewLocator()

	service.Set[sse.SocketManager](locator, service.Singleton, func() *sse.SocketManager {
		return sse.NewSocketManager()
	})

	event.StartListener(locator)

	h.Start(h.AppOpts{
		ServiceLocator: locator,
		LiveReload:     true,
		Register: func(app *h.App) {
			sub, err := fs.Sub(GetStaticAssets(), "assets/dist")

			if err != nil {
				panic(err)
			}

			http.FileServerFS(sub)

			app.Router.Handle("/public/*", http.StripPrefix("/public", http.FileServerFS(sub)))
			app.Router.Handle("/ws/test", sse.HandleWs())
			__htmgo.Register(app.Router)
		},
	})
}
