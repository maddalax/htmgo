package control

import . "htmgo-site/pages/docs"
import "htmgo-site/ui"
import "github.com/maddalax/htmgo/framework/h"

func Loops(ctx *h.RequestContext) *h.Page {
	return DocPage(
		ctx,
		h.Div(
			h.Class("flex flex-col gap-3"),
			Title("Loops / Dealing With Lists"),
			Text(`
				Very commonly you will need to render a list or slice of items onto the page. 
				Frameworks generally solve this in different ways, such as React uses regular JS .map function to solve it.
				htmgo provides a couple of utilities to do so:
			`),
			Text("Example: Rendering a list of tasks"),
			ui.GoCodeSnippet(ListExample),
			Text("Example: Rendering a map"),
			ui.GoCodeSnippet(MapExample),
			NextStep(
				"mt-4",
				PrevBlock("Conditionals", DocPath("/control/if-else")),
				NextBlock("Adding Interactivity", DocPath("/interactivity/swapping")),
			),
		),
	)
}

const ListExample = `
var items = []string{"item1", "item2", "item3"}				
h.List(items, func(item string, index int) *h.Element {
	if tab == TabComplete && item.CompletedAt == nil {
       return h.Empty()
    }
	return h.Div(
		h.Text(item),
	)
})
`

const MapExample = `
var values = map[string]string{
	"key1": "value1",
	"key2": "value2",
	"key3": "value3",
}
h.IterMap(values, func(key string, value string) *h.Element {
	return h.Div(
		h.Text(key),
		h.Text(value),
	)
})		
`
