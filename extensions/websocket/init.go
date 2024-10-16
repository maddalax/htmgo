package websocket

import (
	"github.com/maddalax/htmgo/extensions/websocket/internal/wsutil"
	"github.com/maddalax/htmgo/extensions/websocket/ws"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
)

type WsExtensionOpts struct {
	WsPath string
}

func EnableExtension(app *h.App, opts WsExtensionOpts) {
	if app.Opts.ServiceLocator == nil {
		app.Opts.ServiceLocator = service.NewLocator()
	}
	service.Set[wsutil.SocketManager](app.Opts.ServiceLocator, service.Singleton, func() *wsutil.SocketManager {
		return wsutil.NewSocketManager()
	})
	ws.StartListener(app.Opts.ServiceLocator)
	app.Router.Handle(opts.WsPath, wsutil.WsHttpHandler())
}
