package opts

import "github.com/maddalax/htmgo/framework/h"

type ExtensionOpts struct {
	WsPath    string
	RoomName  func(ctx *h.RequestContext) string
	SessionId func(ctx *h.RequestContext) string
}
