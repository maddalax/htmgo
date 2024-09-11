package ui

import "mhtml/h"

type InputProps struct {
	Id           string
	Label        string
	Name         string
	Type         string
	DefaultValue string
}

func Input(props InputProps) *h.Node {
	input := h.Input(
		props.Type,
		h.Class("border p-2 rounded"),
		h.If(props.Id != "", h.Id(props.Id)),
		h.If(props.Name != "", h.Name(props.Name)),
		h.If(props.DefaultValue != "", h.Attribute("defaultValue", props.DefaultValue)),
	)
	if props.Label != "" {
		return h.Div(
			h.Class("flex flex-col gap-1"),
			h.Label(props.Label),
			input,
		)
	}
	return input
}
