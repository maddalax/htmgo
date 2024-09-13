package news

import (
	"fmt"
	"mhtml/framework/h"
	"mhtml/starter-template/database"
	"time"
)

func StoryList() h.Renderable {

	posts, _ := database.GetOrSet[[]Post]("posts", func() []Post {
		p, _ := List()
		return p
	})

	time.Sleep(200 * time.Millisecond)

	if len(*posts) == 0 {
		return h.Pf("No results found")
	}

	return h.Fragment(
		h.Div(h.List(*posts, func(item Post, index int) h.Renderable {
			return StoryCard(item)
		})),
	)
}

func StoryCard(post Post) h.Renderable {
	url := fmt.Sprintf("/news/%d", post.Id)
	return h.Div(
		h.Class("items-center bg-indigo-200 p-4 rounded"),
		h.A(h.Text(post.Title), h.Href(url)),
	)
}

func StoryFull(id string) h.Renderable {
	post, err := Get(id)
	if err != nil {
		return h.Pf(err.Error())
	}
	return StoryCard(post)
}
