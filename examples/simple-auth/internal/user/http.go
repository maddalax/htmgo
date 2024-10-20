package user

import (
	"github.com/maddalax/htmgo/framework/h"
	"simpleauth/internal/db"
)

func GetUserOrRedirect(ctx *h.RequestContext) (db.User, bool) {
	user, err := GetUserFromSession(ctx)

	if err != nil {
		ctx.Response.Header().Set("Location", "/login")
		ctx.Response.WriteHeader(302)
		return db.User{}, false
	}

	return user, true
}
