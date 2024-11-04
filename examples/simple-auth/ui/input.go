package ui

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/hx"
)

type InputProps struct {
	Id             string
	Label          string
	Name           string
	Type           string
	DefaultValue   string
	Placeholder    string
	Required       bool
	ValidationPath string
	Error          string
	Children       []h.Ren
}

func Input(props InputProps) *h.Element {
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
		h.Class("border p-2 rounded focus:outline-none focus:ring focus:ring-slate-800"),
		h.If(
			props.Name != "",
			h.Name(props.Name),
		),
		h.If(
			props.Children != nil,
			h.Children(props.Children...),
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
		h.If(
			props.Id != "",
			h.Id(props.Id),
		),
		h.Class("flex flex-col gap-1"),
		h.If(
			props.Label != "",
			h.Label(
				h.Text(props.Label),
			),
		),
		input,
		h.Div(
			h.Id(props.Id+"-error"),
			h.Class("text-red-500"),
		),
	)

	return wrapped
}
