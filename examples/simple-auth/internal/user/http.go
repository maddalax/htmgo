package user

import (
	"github.com/maddalax/htmgo/framework/h"
	"simpleauth/internal/db"
)

func GetUserOrRedirect(ctx *h.RequestContext) (db.User, bool) {
	user, err := GetUserFromSession(ctx)

	if err != nil {
		ctx.Redirect("/login", 302)
		return db.User{}, false
	}

	return user, true
}
