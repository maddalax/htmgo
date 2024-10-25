package pages

import (
	"github.com/maddalax/htmgo/framework/h"
	"simpleauth/internal/db"
	"simpleauth/internal/user"
	"simpleauth/partials"
	"simpleauth/ui"
)

func IndexPage(ctx *h.RequestContext) *h.Page {
	u, ok := user.GetUserOrRedirect(ctx)
	if !ok {
		return nil
	}
	return h.NewPage(
		RootPage(UserProfilePage(u)),
	)
}

func UserProfilePage(u db.User) *h.Element {

	meta := user.ParseMeta(u.Metadata)

	return h.Div(
		h.Class("flex flex-col gap-6 items-center pt-10 min-h-screen bg-neutral-100"),
		h.H3F(
			"User Profile",
			h.Class("text-2xl font-bold"),
		),
		h.Pf("Welcome, %s!", u.Email),
		h.Form(
			h.Attribute("hx-swap", "none"),
			h.PostPartial(partials.UpdateProfile),
			h.TriggerChildren(),
			h.Class("flex flex-col gap-4 w-full max-w-md p-6 bg-white rounded-md shadow-md"),
			ui.Input(ui.InputProps{
				Id:           "email",
				Name:         "email",
				Label:        "Email Address",
				Type:         "email",
				DefaultValue: u.Email,
				Children: []h.Ren{
					h.Disabled(),
				},
			}),
			ui.Input(ui.InputProps{
				Name:         "birth-date",
				Label:        "Birth Date",
				DefaultValue: user.GetMetaKey(meta, "birthDate"),
				Type:         "date",
			}),
			ui.Input(ui.InputProps{
				Name:         "favorite-color",
				Label:        "Favorite Color",
				DefaultValue: user.GetMetaKey(meta, "favoriteColor"),
			}),
			ui.Input(ui.InputProps{
				Name:         "occupation",
				Label:        "Occupation",
				DefaultValue: user.GetMetaKey(meta, "occupation"),
			}),
			ui.FormError(""),
			ui.SubmitButton("Save Changes"),
		),
		h.A(
			h.Text("Log out"),
			h.Href("/logout"),
			h.Class("text-blue-400"),
		),
	)
}
