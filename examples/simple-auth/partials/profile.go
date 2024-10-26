package partials

import (
	"github.com/maddalax/htmgo/framework/h"
	"log/slog"
	"simpleauth/internal/user"
	"simpleauth/ui"
)

func UpdateProfile(ctx *h.RequestContext) *h.Partial {
	if !ctx.IsHttpPost() {
		return nil
	}

	patch := map[string]any{
		"birthDate":     ctx.FormValue("birth-date"),
		"favoriteColor": ctx.FormValue("favorite-color"),
		"occupation":    ctx.FormValue("occupation"),
	}

	u, ok := user.GetUserOrRedirect(ctx)

	if !ok {
		return nil
	}

	err := user.SetMeta(ctx, u.ID, patch)

	if err != nil {
		slog.Error("failed to update user profile", slog.String("error", err.Error()))
		ctx.Response.WriteHeader(400)
		return ui.SwapFormError(ctx, "something went wrong")
	}

	return h.RedirectPartial("/")
}
