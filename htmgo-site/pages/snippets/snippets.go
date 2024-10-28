package snippets

import (
	"htmgo-site/partials/snippets"
)

var FormWithLoadingStateSnippet = Snippet{
	name:        "Form",
	description: "A simple form submission example with a loading state",
	sidebarName: "Form with loading state",
	path:        "/snippets/form",
	partial:     snippets.FormExample,
}

var UserAuthSnippet = Snippet{
	name:           "User Authentication",
	description:    "An example showing basic user registration and login with htmgo",
	sidebarName:    "User Authentication",
	path:           "/snippets/user-auth",
	externalRoute:  "https://auth-example.htmgo.dev",
	sourceCodePath: "https://github.com/maddalax/htmgo/tree/master/examples/simple-auth",
}

var ChatSnippet = Snippet{
	name:           "Chat App",
	description:    "A simple chat application built with htmgo",
	sidebarName:    "Chat App",
	path:           "/snippets/chat",
	externalRoute:  "https://chat-example.htmgo.dev",
	sourceCodePath: "https://github.com/maddalax/htmgo/tree/master/examples/chat",
}

var HackerNewsSnippet = Snippet{
	name:           "Hacker News Clone",
	description:    "A hacker news reader clone built with htmgo",
	sidebarName:    "Hacker News Clone",
	path:           "/snippets/hackernews",
	externalRoute:  "https://hn.htmgo.dev",
	sourceCodePath: "https://github.com/maddalax/htmgo/tree/master/examples/hackernews",
}

var Snippets = []Snippet{
	FormWithLoadingStateSnippet,
	UserAuthSnippet,
	ChatSnippet,
	HackerNewsSnippet,
}
