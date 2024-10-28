package examples

import "htmgo-site/partials/snippets"

var FormWithLoadingStateSnippet = Snippet{
	name:        "Form",
	description: "A simple form submission example with a loading state",
	sidebarName: "Form with loading state",
	path:        "/examples/form",
	partial:     snippets.FormExample,
}

var UserAuthSnippet = Snippet{
	name:           "User Authentication",
	description:    "An example showing basic user registration and login with htmgo",
	sidebarName:    "User Authentication",
	path:           "/examples/user-auth",
	externalRoute:  "https://auth-example.htmgo.dev",
	sourceCodePath: "https://github.com/maddalax/htmgo/tree/master/examples/simple-auth",
}

var ChatSnippet = Snippet{
	name:           "Chat App",
	description:    "A simple chat application built with htmgo using SSE for real-time updates",
	sidebarName:    "Chat App",
	path:           "/examples/chat",
	externalRoute:  "https://chat-example.htmgo.dev",
	sourceCodePath: "https://github.com/maddalax/htmgo/tree/master/examples/chat",
}

var HackerNewsSnippet = Snippet{
	name:           "Hacker News Clone",
	description:    "A hacker news reader clone built with htmgo",
	sidebarName:    "Hacker News Clone",
	path:           "/examples/hackernews",
	externalRoute:  "https://hn.htmgo.dev",
	sourceCodePath: "https://github.com/maddalax/htmgo/tree/master/examples/hackernews",
}

var HtmgoSiteSnippet = Snippet{
	name:           "Htmgo Doc Site",
	description:    "The htmgo site built with htmgo, recursion am I right?",
	sidebarName:    "Htmgo Doc Site",
	path:           "/examples/htmgo-site",
	externalRoute:  "https://htmgo.dev",
	sourceCodePath: "https://github.com/maddalax/htmgo/tree/master/htmgo-site",
}

var TodoListSnippet = Snippet{
	name:           "Todo List",
	description:    "A todo list built with htmgo",
	sidebarName:    "Todo List",
	path:           "/examples/todolist",
	externalRoute:  "https://todo-example.htmgo.dev",
	sourceCodePath: "https://github.com/maddalax/htmgo/tree/master/examples/todo-list",
}

var examples = []Snippet{
	FormWithLoadingStateSnippet,
	UserAuthSnippet,
	ChatSnippet,
	HackerNewsSnippet,
	TodoListSnippet,
	HtmgoSiteSnippet,
}
