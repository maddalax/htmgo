package ui

import (
	"github.com/maddalax/htmgo/framework/h"
)

type CenteredFormProps struct {
	Title      string
	Children   []h.Ren
	SubmitText string
	PostUrl    string
}

func CenteredForm(props CenteredFormProps) *h.Element {
	return h.Div(
		h.Class("flex flex-col items-center justify-center min-h-screen bg-neutral-100"),
		h.Div(
			h.Class("bg-white p-8 rounded-lg shadow-lg w-full max-w-md"),
			h.H2F(
				props.Title,
				h.Class("text-3xl font-bold text-center mb-6"),
			),
			h.Form(
				h.TriggerChildren(),
				h.Post(props.PostUrl),
				h.Attribute("hx-swap", "none"),
				h.Class("flex flex-col gap-4"),
				h.Children(props.Children...),
				// Error message
				FormError(""),
				// Submit button at the bottom
				SubmitButton(props.SubmitText),
			),
		),
	)
}
