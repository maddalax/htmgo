package partials

import (
	"fmt"
	"github.com/maddalax/htmgo/framework/h"
	"hackernews/internal/news"
	"hackernews/internal/timeformat"
	"strconv"
)

func Story(ctx *h.RequestContext) *h.Partial {
	storyId, err := strconv.ParseInt(ctx.QueryParam("item"), 10, 64)

	if err != nil {
		return h.SwapManyPartial(
			ctx,
			h.Div(
				h.Text("Invalid story id"),
			),
		)
	}

	story, err := news.GetStory(int(storyId))

	if err != nil {
		return h.SwapManyPartial(
			ctx,
			h.Div(
				h.Text("Failed to load story"),
			),
		)
	}

	if ctx.IsHxRequest() {
		return h.SwapManyPartialWithHeaders(
			ctx,
			h.PushUrlHeader(fmt.Sprintf("/?item=%d", storyId)),
			StoryBody(story),
		)
	}

	return h.NewPartial(
		StoryBody(story),
	)
}

func StoryBody(story *news.Story) *h.Element {
	return h.Div(
		h.Id("story-body"),
		h.Div(
			h.Class("prose prose-2xl bg-white border-b border-gray-200 pb-3 max-w-3xl"),
			h.H5(
				h.Class("flex gap-2 items-center font-bold"),
				h.UnsafeRaw(story.Title),
			),
			h.A(
				h.Href(story.Url),
				h.Class("text-sm text-rose-400 no-underline"),
				h.Text(story.Url),
			),
			h.Div(
				h.Class("text-sm text-gray-600"),
				h.UnsafeRaw(story.Text),
			),
			h.Div(
				h.Class("text-sm text-gray-600 mt-2"),
				h.TextF("%d upvotes ", story.Score),
				h.UnsafeRaw("&bull;"),
				h.TextF(" %s ", story.By),
				h.UnsafeRaw("&bull;"),
				h.TextF(" %s", timeformat.RelativeTime(story.Time)),
			),
		),
		h.Div(
			h.Class("mt-2 max-w-3xl"),
			h.GetPartial(StoryComments, "load"),
		),
	)
}
