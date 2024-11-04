package pages

import (
	"github.com/maddalax/htmgo/framework/h"
	"simpleauth/partials"
	"simpleauth/ui"
)

func Login(ctx *h.RequestContext) *h.Page {
	return h.NewPage(
		RootPage(
			ui.CenteredForm(ui.CenteredFormProps{
				Title:      "Sign In",
				SubmitText: "Sign In",
				PostUrl:    h.GetPartialPath(partials.LoginUser),
				Children: []h.Ren{
					ui.Input(ui.InputProps{
						Id:       "username",
						Name:     "email",
						Label:    "Email Address",
						Type:     "email",
						Required: true,
						Children: []h.Ren{
							h.Attribute("autocomplete", "off"),
							h.MaxLength(50),
						},
					}),

					ui.Input(ui.InputProps{
						Id:       "password",
						Name:     "password",
						Label:    "Password",
						Type:     "password",
						Required: true,
						Children: []h.Ren{
							h.MinLength(6),
						},
					}),

					h.A(
						h.Href("/register"),
						h.Text("Don't have an account? Register here"),
						h.Class("text-blue-500"),
					),
				},
			}),
		),
	)
}
