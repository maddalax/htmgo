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
						h.Class("flex flex-col gap-6"),

						// Username input at the top
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

						// Single box for Create or Join a Chat Room
						h.Div(
							h.Class("p-4 border border-gray-300 rounded-md flex flex-col gap-6"),

							// Create New Chat Room input
							components.Input(components.InputProps{
								Name:        "new-chat-room",
								Label:       "Create a new chat room",
								Placeholder: "Enter chat room name",
								Children: []h.Ren{
									h.Attribute("autocomplete", "off"),
									h.MaxLength(20),
								},
							}),

							// OR divider
							h.Div(
								h.Class("flex items-center justify-center gap-4"),
								h.Div(h.Class("border-t border-gray-300 flex-grow")),
								h.P(h.Text("OR"), h.Class("text-gray-500")),
								h.Div(h.Class("border-t border-gray-300 flex-grow")),
							),

							// Join Chat Room input
							components.Input(components.InputProps{
								Id:           "join-chat-room",
								Name:         "join-chat-room",
								Label:        "Join an existing chat room",
								Placeholder:  "Enter chat room ID",
								DefaultValue: ctx.QueryParam("roomId"),
								Children: []h.Ren{
									h.Attribute("autocomplete", "off"),
									h.MaxLength(100),
								},
							}),
						),

						// Error message
						components.FormError(""),

						// Submit button at the bottom
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
