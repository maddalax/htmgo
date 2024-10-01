package pages

import (
	"chat/components"
	"chat/partials"
	"github.com/maddalax/htmgo/framework/h"
)

func ChatAppFirstScreen(ctx *h.RequestContext) *h.Page {
	return h.NewPage(
		RootPage(
			h.Div(
				h.Class("flex flex-col items-center justify-center min-h-screen bg-neutral-100"),
				h.Div(
					h.Class("bg-white p-8 rounded-lg shadow-lg w-full max-w-md"),
					h.H2F("htmgo chat", h.Class("text-3xl font-bold text-center mb-6")),
					h.Form(
						h.Attribute("hx-swap", "none"),
						h.PostPartial(partials.CreateOrJoinRoom),
						h.Class("flex flex-col gap-3"),

						components.Input(components.InputProps{
							Id:       "username",
							Name:     "username",
							Label:    "Username",
							Required: true,
							Children: []h.Ren{
								h.Attribute("autocomplete", "off"),
								h.MaxLength(15),
							},
						}),

						h.Div(
							h.Class("mt-6 flex flex-col gap-3"),

							components.Input(components.InputProps{
								Name:        "new-chat-room",
								Label:       "Create a New Chat Room",
								Placeholder: "Chat Room Name",
								Children: []h.Ren{
									h.Attribute("autocomplete", "off"),
									h.MaxLength(20),
								},
							}),

							h.Div(
								h.Class("flex items-center justify-center gap-4"),
								h.Div(h.Class("border-t border-gray-300 flex-grow")),
								h.P(h.Text("OR"), h.Class("text-gray-500")),
								h.Div(h.Class("border-t border-gray-300 flex-grow")),
							),

							components.Input(components.InputProps{
								Id:          "join-chat-room",
								Name:        "join-chat-room",
								Label:       "Join a Chat Room",
								Placeholder: "Chat Room Id",
								Children: []h.Ren{
									h.Attribute("autocomplete", "off"),
									h.MaxLength(100),
								},
							}),
						),

						components.FormError(""),
						components.PrimaryButton(components.ButtonProps{
							Type: "submit",
							Text: "Submit",
						}),
					),
				),
			),
		),
	)
}
