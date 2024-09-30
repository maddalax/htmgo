package main

import (
	"chat/__htmgo"
	"chat/chat"
	"chat/internal/db"
	"chat/ws"
	"embed"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
	"io/fs"
	"net/http"
)

//go:embed assets/dist/*
var StaticAssets embed.FS

func main() {
	locator := service.NewLocator()

	service.Set[db.Queries](locator, service.Singleton, db.Provide)
	service.Set[ws.SocketManager](locator, service.Singleton, func() *ws.SocketManager {
		return ws.NewSocketManager()
	})

	chatManager := chat.NewManager(locator)
	go chatManager.StartListener()

	h.Start(h.AppOpts{
		ServiceLocator: locator,
		LiveReload:     true,
		Register: func(app *h.App) {
			sub, err := fs.Sub(StaticAssets, "assets/dist")

			if err != nil {
				panic(err)
			}

			http.FileServerFS(sub)

			app.Router.Handle("/public/*", http.StripPrefix("/public", http.FileServerFS(sub)))
			app.Router.Handle("/chat", ws.Handle())

			__htmgo.Register(app.Router)
		},
	})
}
