package partials

import (
	"fmt"
	"github.com/maddalax/htmgo/framework/h"
	"hackernews/internal/batch"
	"hackernews/internal/news"
	"hackernews/internal/sanitize"
	"hackernews/internal/timeformat"
	"strings"
	"time"
)

func StoryComments(ctx *h.RequestContext) *h.Partial {
	return h.NewPartial(
		h.Fragment(
			h.OobSwap(
				ctx,
				h.Div(
					h.Id("comments-loader"),
				),
			),
			h.Div(
				h.Class("flex flex-col gap-3 prose max-w-none"),
				CachedStoryComments(news.MustItemId(ctx)),
			),
		),
	)
}

var CachedStoryComments = h.CachedPerKeyT[string, int](time.Minute*3, func(itemId int) (string, h.GetElementFunc) {
	return fmt.Sprintf("story-comments-%d", itemId), func() *h.Element {
		story, err := news.GetStory(itemId)

		if err != nil {
			return h.Div(
				h.Text("Failed to load story"),
			)
		}

		comments := news.GetComments(story.Kids)

		// parallel process because each comment needs to load its children comments
		items := batch.ParallelProcess[news.Comment, *h.Element](comments, 50, func(item news.Comment) *h.Element {
			return Comment(item, 0)
		})

		return h.List(items, func(item *h.Element, index int) *h.Element {
			return item
		})
	}
})

func Comment(item news.Comment, nesting int) *h.Element {
	if item.Text == "" {
		return h.Empty()
	}

	children := news.GetComments(item.Kids)

	return h.Div(
		h.ClassX("block bg-white pb-2 pt-2", h.ClassMap{
			"border-b border-gray-200": nesting == 0,
			"border-l border-gray-200": nesting > 0,
		}),
		h.If(
			nesting > 0,
			h.Attribute("style", fmt.Sprintf("margin-left: %dpx", (nesting-1)*15)),
		),
		h.Div(
			h.If(
				nesting > 0,
				h.Class("pl-4"),
			),
			h.Div(
				h.Class("flex gap-1 items-center"),
				h.Div(
					h.Class("font-bold text-rose-500"),
					h.UnsafeRaw(sanitize.Sanitize(item.By)),
				),
				h.Div(
					h.Class("text-sm text-gray-600"),
					h.UnsafeRaw("&bull;"),
					h.TextF(" %s", timeformat.RelativeTime(item.Time)),
				),
			),
			h.Div(
				h.Class("text-sm text-gray-600"),
				h.UnsafeRaw(sanitize.Sanitize(strings.TrimSpace(item.Text))),
			),
		),
		h.If(
			len(children) > 0,
			h.List(
				children, func(child news.Comment, index int) *h.Element {
					return h.Div(
						Comment(child, nesting+1),
					)
				},
			),
		),
	)
}
