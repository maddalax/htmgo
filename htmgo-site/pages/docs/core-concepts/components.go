package core_concepts

import (
	"github.com/maddalax/htmgo/framework/h"
	. "htmgo-site/pages/docs"
	"htmgo-site/ui"
)

const ComponentExampleSnippet = `func Card(ctx *h.RequestContext) *h.Element {
	service := tasks.NewService(ctx.ServiceLocator())
	cards, _ := service.List()

	return h.Div(
		h.Id("task-card"),
		h.Class("bg-white w-full rounded shadow-md"),
		CardBody(cards),
	)
}
`

func Components(ctx *h.RequestContext) *h.Page {
	return DocPage(
		ctx,
		h.Div(
			h.Class("flex flex-col gap-3"),
			Title("Components"),
			Text(`
				Components are re-usable bits of logic to render HTML. Similar to how in React components are Javascript functions, in htmgo, components are pure go functions.
				A component can be pure, or it can have data fetching logic inside of it. Since htmgo uses htmx for interactivity, there is NO re-rendering of your UI automatically from the framework, which means you can safely put data fetching logic inside of components since you can be sure they will only be called by your own code.
			`),
			ComponentExample(),
			NextStep(
				"mt-4",
				PrevBlock("Partials", DocPath("/core-concepts/partials")),
				NextBlock("Tags and Attributes", DocPath("/core-concepts/tags-and-attributes")),
			),
		),
	)
}

func ComponentExample() *h.Element {
	return h.Div(
		Text("Example:"),
		ui.GoCodeSnippet(ComponentExampleSnippet),
		Text(`
			My card component here fetches all my tasks I have on my list, and renders each task.
	        If you are familiar with React, then you would likely place this fetch logic inside of a useEffect or (useQuery library) so it is not constantly re-fetched as the component re-renders.
			With htmgo, the only way to update content on the page is to use htmx to swap out the content from loading a partial. Therefore you control exactly when this card component is called, not the framework behind the scenes.

			You'll learn more about swapping in the next few pages.
		`),
	)
}
