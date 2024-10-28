package snippets

import (
	"htmgo-site/partials/snippets"
)

var FormWithLoadingState = Snippet{
	name:        "Form",
	description: "A simple form submission example with a loading state",
	sidebarName: "Form with loading state",
	path:        "/snippets/form",
	partial:     snippets.FormExample,
}

var Snippets = []Snippet{
	FormWithLoadingState,
}
