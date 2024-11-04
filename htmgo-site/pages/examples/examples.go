package examples

import "htmgo-site/partials/snippets"

var FormWithLoadingStateSnippet = Snippet{
	category:    "Forms",
	name:        "Form",
	description: "A simple form submission example with a loading state",
	sidebarName: "Form With Loading State",
	path:        "/examples/form",
	partial:     snippets.FormExample,
}

var FormWithBlurValidationSnippet = Snippet{
	category:    "Forms",
	name:        "Form",
	description: "A simple form submission example with validation on blur",
	sidebarName: "Form With Blur Validation",
	path:        "/examples/form-with-blur-validation",
	partial:     snippets.FormWithBlurValidation,
}

var UserAuthSnippet = Snippet{
	category:       "Projects",
	name:           "User Authentication",
	description:    "An example showing basic user registration and login with htmgo",
	sidebarName:    "User Authentication",
	path:           "/examples/user-auth",
	externalRoute:  "https://auth-example.htmgo.dev",
	sourceCodePath: "https://github.com/maddalax/htmgo/tree/master/examples/simple-auth",
}

var ChatSnippet = Snippet{
	category:       "Projects",
	name:           "Chat App",
	description:    "A simple chat application built with htmgo using SSE for real-time updates",
	sidebarName:    "Chat App Using SSE",
	path:           "/examples/chat",
	externalRoute:  "https://chat-example.htmgo.dev",
	sourceCodePath: "https://github.com/maddalax/htmgo/tree/master/examples/chat",
}

var HackerNewsSnippet = Snippet{
	category:       "Projects",
	name:           "HackerNews Clone",
	description:    "A hacker news reader clone built with htmgo",
	sidebarName:    "HackerNews Clone",
	path:           "/examples/hackernews",
	externalRoute:  "https://hn.htmgo.dev",
	sourceCodePath: "https://github.com/maddalax/htmgo/tree/master/examples/hackernews",
}

var HtmgoSiteSnippet = Snippet{
	category:       "Projects",
	name:           "Htmgo Doc Site",
	description:    "The htmgo site built with htmgo, recursion am I right?",
	sidebarName:    "Htmgo Doc Site",
	path:           "/examples/htmgo-site",
	externalRoute:  "https://htmgo.dev",
	sourceCodePath: "https://github.com/maddalax/htmgo/tree/master/htmgo-site",
}

var TodoListSnippet = Snippet{
	category:       "Projects",
	name:           "Todo List",
	description:    "A todo list built with htmgo",
	sidebarName:    "Todo List",
	path:           "/examples/todolist",
	externalRoute:  "https://todo-example.htmgo.dev",
	sourceCodePath: "https://github.com/maddalax/htmgo/tree/master/examples/todo-list",
}

var ClickToEditSnippet = Snippet{
	category:    "Forms",
	name:        "Inline Click To Edit",
	description: "List view of items with a click to edit button and persistence",
	sidebarName: "Inline Click To Edit",
	path:        "/examples/click-to-edit",
	partial:     snippets.ClickToEdit,
}

var JsSetTextOnClick = Snippet{
	category:    "Interactivity (JS)",
	name:        "Set Element Text On Click",
	description: "A simple example of how to use htmgo with javascript",
	sidebarName: "Set Text On Click",
	path:        "/examples/js-set-text-on-click",
	partial:     snippets.SetTextOnClick,
}

var JsHideChildrenOnClick = Snippet{
	category:    "Interactivity (JS)",
	name:        "Hide / Show Children On Click",
	description: "Use JS to hide and show children elements on click",
	sidebarName: "Hide / Show Children",
	path:        "/examples/js-hide-children-on-click",
	partial:     snippets.JsHideChildrenOnClick,
}

var InputComponentSnippet = Snippet{
	category:    "Components",
	name:        "Input",
	description: "An example of how you could build a re-usable input component",
	sidebarName: "Text Input",
	path:        "/examples/input-component",
	partial:     snippets.InputComponent,
}

var examples = []Snippet{
	FormWithLoadingStateSnippet,
	FormWithBlurValidationSnippet,
	ClickToEditSnippet,

	JsSetTextOnClick,
	JsHideChildrenOnClick,

	UserAuthSnippet,
	ChatSnippet,
	HackerNewsSnippet,
	TodoListSnippet,
	HtmgoSiteSnippet,

	InputComponentSnippet,
}
