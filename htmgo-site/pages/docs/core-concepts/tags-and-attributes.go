package core_concepts

import "htmgo-site/ui"
import "github.com/maddalax/htmgo/framework/h"
import . "htmgo-site/pages/docs"

func TagsAndAttributes(ctx *h.RequestContext) *h.Page {
	return DocPage(
		ctx,
		h.Div(
			h.Class("flex flex-col gap-3"),
			Title("Tags and Attributes"),
			Text(`
				In htmgo, html is built using a set of functions that return <b>*h.Element</b>. 
				These functions are all defined in the <b>'h'</b> package in htmgo/framework
				htmgo provides methods to render most if not all html tags and attributes.
			`),
			Text(`Example:`),
			ui.GoCodeSnippet(TagExample),
			Text(`
				All methods can be found in the 'h' package in htmgo/framework			
			`),
			Text("<b>h.Tag</b> and <b>h.Attribute</b> are available to use when you need to render a tag or attribute that htmgo does not provide a method for."),
			ui.GoCodeSnippet(TagExampleUsingTagFunc),
			Text(`
			<b>Attributes</b> are one of the main ways we can add interactivity to the pages with htmx.
			htmgo provides various methods to add attributes to elements, as well as adding attributes based on a condition.
			`),
			ui.GoCodeSnippet(AttributeExample),
			HelpText("In this example we are conditionally adding an attribute based on if there is an error on not, you'll learn more about conditionals in the next few pages."),
			Text("Example using htmx attributes:"),
			ui.GoCodeSnippet(HxAttributeExample),
			NextStep(
				"mt-4",
				PrevBlock("Components", DocPath("/core-concepts/components")),
				NextBlock("Raw HTML", DocPath("/core-concepts/raw-html")),
			),
		),
	)
}

const TagExample = `h.Div(
	h.Class("flex gap-2"),
	h.Button(
		h.Text("Submit"),
	),
)
`

const TagExampleUsingTagFunc = `h.Tag("my-custom-tag",
	h.Class("flex gap-2"),
	h.Button(
        h.Attribute("x-custom-attr", "my-value"),
		h.Text("Submit"),
	),
)
`

const AttributeExample = `h.Div(
	h.Class("flex gap-2"),
	h.Id("my-div"),
    h.If(
      error != "",
	  h.Class("p-4 bg-rose-400 text-white rounded"),
	)
)
`

const HxAttributeExample = `h.Tr(
	h.Class("flex gap-2"),
	h.HxInclude("input")
	h.Td(
		h.Input("text",
			h.Class("p-4 rounded"),
			h.Placeholder("Type something"),
			h.Name("my-input"),
		)
	),
   	h.Td(
		h.Button(
			h.Text("Submit"),
		)
	),
)`
