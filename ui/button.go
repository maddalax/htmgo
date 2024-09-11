package ui

import (
	"mhtml/h"
)

type ButtonProps struct {
	Text     string
	Target   string
	Get      string
	Class    string
	Children *h.Node
}

func PrimaryButton(props ButtonProps) *h.Node {
	props.Class = h.MergeClasses(props.Class, "border-blue-700 bg-blue-700 text-white")
	return Button(props)
}

func SecondaryButton(props ButtonProps) *h.Node {
	props.Class = h.MergeClasses(props.Class, "border-gray-700 bg-gray-700 text-white")
	return Button(props)
}

func Button(props ButtonProps) *h.Node {

	text := h.P(props.Text)

	button := h.Button(
		h.If(props.Children != nil, props.Children),
		h.Class("flex gap-1 items-center border p-4 rounded cursor-hover", props.Class),
		h.If(props.Get != "", h.Get(props.Get)),
		h.If(props.Target != "", h.Target(props.Target)),
		//h.BeforeRequestSetHtml(
		//	h.Div(
		//		h.Class("flex gap-1"),
		//		h.Text("Loading..."),
		//	),
		//),
		//h.AfterRequestSetHtml(h.Text(props.Text)),
		// Note, i really like this idea of being able to reference elements just by the instance,
		//and automatically adding id
		text,
	)

	return button
}
