package pages

import "github.com/maddalax/htmgo/framework/h"

func LogoutPage(ctx *h.RequestContext) *h.Page {

	// clear the session cookie
	ctx.Response.Header().Set(
		"Set-Cookie",
		"session_id=; Path=/; Max-Age=0",
	)

	ctx.Response.Header().Set(
		"Location",
		"/login",
	)

	ctx.Response.WriteHeader(
		302,
	)

	return nil
}
