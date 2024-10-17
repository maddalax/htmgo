package opts

import "github.com/maddalax/htmgo/framework/h"

type ExtensionOpts struct {
	WsPath    string
	SessionId func(ctx *h.RequestContext) string
}
