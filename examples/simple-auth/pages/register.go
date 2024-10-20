package pages

import (
	"github.com/maddalax/htmgo/framework/h"
	"simpleauth/partials"
	"simpleauth/ui"
)

func Register(ctx *h.RequestContext) *h.Page {
	return h.NewPage(
		RootPage(
			ui.CenteredForm(ui.CenteredFormProps{
				PostUrl:    h.GetPartialPath(partials.RegisterUser),
				Title:      "Create an Account",
				SubmitText: "Register",
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
						h.Href("/login"),
						h.Text("Already have an account? Login here"),
						h.Class("text-blue-500"),
					),
				},
			}),
		),
	)
}
