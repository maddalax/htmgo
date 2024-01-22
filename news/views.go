package news

import (
	"fmt"
	"mhtml/h"
)

func StoryList() *h.Node {
	posts, err := List()

	if err != nil {
		return h.P(err.Error())
	}

	if len(posts) == 0 {
		return h.P("No results found")
	}

	return h.Fragment(
		h.VStack(h.List(posts, func(item Post) *h.Node {
			return StoryCard(item)
		})),
	)
}

func StoryCard(post Post) *h.Node {
	url := fmt.Sprintf("/news/%d", post.Id)
	return h.VStack(
		h.Class("items-center bg-red-200 p-4 rounded"),
		h.A(post.Title, h.Href(url)),
	)
}

func StoryFull(id string) *h.Node {
	post, err := Get(id)
	if err != nil {
		return h.P(err.Error())
	}
	return StoryCard(post)
}
