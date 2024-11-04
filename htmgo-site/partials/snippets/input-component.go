package snippets

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/hx"
)

// InputComponent wrapper to make the code snippet work, main code is the Input function
func InputComponent(ctx *h.RequestContext) *h.Partial {

	return h.NewPartial(
		h.Div(
			h.Class("max-w-sm mx-auto flex flex-col gap-4"),
			Input(
				InputProps{
					Id:          "my-input",
					Name:        "my-input",
					Label:       "Input with label",
					Type:        "text",
					Placeholder: "Type something",
					Required:    true,
				},
				h.Attribute("autocomplete", "off"),
				h.MaxLength(50),
			),
			Input(
				InputProps{
					Id:           "my-input",
					Name:         "my-input",
					Label:        "Input with default value",
					Type:         "text",
					DefaultValue: "Default value",
				},
			),
			Input(
				InputProps{
					Id:          "my-input",
					Name:        "my-input",
					Label:       "Input with helper text",
					Type:        "text",
					Placeholder: "Full name",
					HelperText:  "This should be your full legal name",
				},
			),
		),
	)
}

type InputProps struct {
	Id             string
	Label          string
	Name           string
	Type           string
	DefaultValue   string
	Placeholder    string
	Required       bool
	ValidationPath string
	HelperText     string
}

func Input(props InputProps, children ...h.Ren) *h.Element {
	validation := h.If(
		props.ValidationPath != "",
		h.Children(
			h.Post(props.ValidationPath, hx.BlurEvent),
			h.Attribute("hx-swap", "innerHTML transition:true"),
			h.Attribute("hx-target", "next div"),
		),
	)

	if props.Type == "" {
		props.Type = "text"
	}

	input := h.Input(
		props.Type,
		h.Class("border p-2 rounded"),
		h.If(
			props.Id != "",
			h.Id(props.Id),
		),
		h.If(
			props.Name != "",
			h.Name(props.Name),
		),
		h.If(
			children != nil,
			h.Children(children...),
		),
		h.If(
			props.Required,
			h.Required(),
		),
		h.If(
			props.Placeholder != "",
			h.Placeholder(props.Placeholder),
		),
		h.If(
			props.DefaultValue != "",
			h.Attribute("value", props.DefaultValue),
		),
		validation,
	)

	wrapped := h.Div(
		h.Class("flex flex-col gap-1"),
		h.If(
			props.Label != "",
			h.Label(
				h.Text(props.Label),
			),
		),
		input,
		h.If(
			props.HelperText != "",
			h.Div(
				h.Class("text-slate-600 text-sm"),
				h.Text(props.HelperText),
			),
		),
		h.Div(
			h.Id(props.Id+"-error"),
			h.Class("text-red-500"),
		),
	)

	return wrapped
}
