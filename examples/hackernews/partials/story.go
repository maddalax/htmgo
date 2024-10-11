package partials

import (
	"fmt"
	"github.com/maddalax/htmgo/framework/h"
	"hackernews/internal/news"
	"hackernews/internal/timeformat"
	"time"
)

func Story(ctx *h.RequestContext) *h.Partial {
	storyId := news.MustItemId(ctx)
	if storyId == 0 {
		return h.NewPartial(
			h.Div(
				h.Class("flex justify-center bg-neutral-300"),
				h.Id("story-body"),
			),
		)
	}

	if ctx.IsHxRequest() {
		return h.SwapManyPartialWithHeaders(
			ctx,
			h.PushUrlHeader(fmt.Sprintf("/?item=%d", storyId)),
			h.Div(
				h.Class("w-full"),
				h.Id("story-body"),
				CachedStoryBody(storyId),
			),
		)
	}

	return h.NewPartial(
		CachedStoryBody(storyId),
	)
}

var CachedStoryBody = h.CachedPerKeyT[string, int](time.Minute*3, func(itemId int) (string, h.GetElementFunc) {
	return fmt.Sprintf("story-%d", itemId), func() *h.Element {
		story, err := news.GetStory(itemId)
		if err != nil {
			return h.Div(
				h.Id("story-body"),
				h.Text("Failed to load story"),
			)
		}
		return StoryBody(story)
	}
})

func StoryBody(story *news.Story) *h.Element {
	return h.Div(
		h.Class("w-full"),
		h.Id("story-body"),
		h.Div(
			h.Class("prose prose-2xl border-b border-gray-200 pb-3 max-w-none w-full"),
			h.H5(
				h.Class("flex gap-2 items-left font-bold"),
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
			h.Id("comments-loader"),
			h.Class("flex justify-center items-center h-24"),
			h.Div(
				h.Class("animate-spin rounded-full h-8 w-8 border-t-2 border-b-2 border-rose-500"),
			),
		),
		h.Div(
			h.Class("mt-2 min-w-3xl max-w-3xl"),
			h.GetPartial(StoryComments, "load"),
		),
	)
}
