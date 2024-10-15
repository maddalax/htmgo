package main

import (
	"encoding/json"
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
			app.Router.Get("/metrics", func(writer http.ResponseWriter, request *http.Request) {
				writer.Header().Set("Content-Type", "application/json")
				writer.WriteHeader(http.StatusOK)
				metrics := event.GetMetrics()
				serialized, _ := json.Marshal(metrics)
				_, _ = writer.Write(serialized)
			})
			__htmgo.Register(app.Router)
		},
	})
}
