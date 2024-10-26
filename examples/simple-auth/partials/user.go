package partials

import (
	"github.com/maddalax/htmgo/framework/h"
	"simpleauth/internal/user"
	"simpleauth/ui"
)

func RegisterUser(ctx *h.RequestContext) *h.Partial {
	if !ctx.IsHttpPost() {
		return nil
	}

	payload := user.CreateUserRequest{
		Email:    ctx.FormValue("email"),
		Password: ctx.FormValue("password"),
	}

	id, err := user.Create(
		ctx,
		payload,
	)

	if err != nil {
		ctx.Response.WriteHeader(400)
		return ui.SwapFormError(ctx, err.Error())
	}

	session, err := user.CreateSession(ctx, id)

	if err != nil {
		ctx.Response.WriteHeader(500)
		return ui.SwapFormError(ctx, "something went wrong")
	}

	user.WriteSessionCookie(ctx, session)

	return h.RedirectPartial("/")
}

func LoginUser(ctx *h.RequestContext) *h.Partial {
	if !ctx.IsHttpPost() {
		return nil
	}

	payload := user.LoginUserRequest{
		Email:    ctx.FormValue("email"),
		Password: ctx.FormValue("password"),
	}

	_, err := user.Login(
		ctx,
		payload,
	)

	if err != nil {
		ctx.Response.WriteHeader(400)
		return ui.SwapFormError(ctx, err.Error())
	}

	return h.RedirectPartial("/")
}
