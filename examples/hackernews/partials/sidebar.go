package partials

import (
	"fmt"
	"github.com/maddalax/htmgo/framework/h"
	"hackernews/components"
	"hackernews/internal/news"
	"hackernews/internal/parse"
	"hackernews/internal/timeformat"
	"time"
)

var ScrollJs = `
	const scrollContainer = self;
    let isDown = false;
    let startX;
    let scrollLeft;

    scrollContainer.addEventListener("mousedown", (e) => {
        isDown = true;
        scrollContainer.classList.add("active");
        startX = e.pageX - scrollContainer.offsetLeft;
        scrollLeft = scrollContainer.scrollLeft;
    });

    scrollContainer.addEventListener("mouseleave", () => {
        isDown = false;
        scrollContainer.classList.remove("active");
    });

    scrollContainer.addEventListener("mouseup", () => {
        isDown = false;
        scrollContainer.classList.remove("active");
    });

    scrollContainer.addEventListener("mousemove", (e) => {
        if (!isDown) return;
        e.preventDefault();
        const x = e.pageX - scrollContainer.offsetLeft;
        const walk = (x - startX) * 3; // Adjust scroll speed here
        scrollContainer.scrollLeft = scrollLeft - walk;
    });
`

func StorySidebar(ctx *h.RequestContext) *h.Partial {
	category := h.GetQueryParam(ctx, "category")
	pageRaw := h.GetQueryParam(ctx, "page")
	mode := h.GetQueryParam(ctx, "mode")

	if pageRaw == "" {
		pageRaw = "0"
	}

	if category == "" {
		category = "topstories"
	}

	page := parse.MustParseInt(pageRaw, 0)

	fetchMorePath := h.GetPartialPathWithQs(
		StorySidebar,
		h.NewQs("mode", "infinite", "page", fmt.Sprintf("%d", page+1), "category", category),
	)

	list := CachedStoryList(category, page, 50, fetchMorePath)

	body := h.Aside(
		h.Id("story-sidebar"),
		h.JoinExtensions(
			h.TriggerChildren(),
		),
		h.Class("sticky top-0 h-screen p-1 bg-gray-100 overflow-y-auto max-w-80 min-w-80"),
		h.Div(
			h.Class("flex flex-col gap-1"),
			SidebarTitle(category),
			h.Id("story-list"),
			list,
		),
	)

	if mode == "infinite" {
		return h.NewPartial(
			list,
		)
	}

	if ctx.IsHxRequest() {
		return h.SwapManyPartial(ctx, body)
	}

	return h.NewPartial(body)
}

func SidebarTitle(defaultCategory string) *h.Element {
	today := time.Now().Format("Mon, 02 Jan 2006")
	return h.Div(
		h.Class("flex flex-col px-2 pt-4 pb-2"),
		h.Div(
			h.Class("text-sm text-gray-600"),
			h.Text(today),
		),
		h.Div(
			h.Class("font-bold text-xl"),
			h.Text("Hacker News"),
		),
		h.Div(
			h.OnLoad(
				h.EvalJs(ScrollJs),
			),
			h.Class("scroll-container mt-2 flex gap-1 no-scrollbar overflow-y-hidden whitespace-nowrap overflow-x-auto"),
			h.List(news.Categories, func(item news.Category, index int) *h.Element {
				return CategoryBadge(defaultCategory, item)
			}),
		),
	)
}

func CategoryBadge(defaultCategory string, category news.Category) *h.Element {
	selected := category.Path == defaultCategory
	return components.Badge(
		category.Name,
		selected,
		h.Attribute("hx-swap", "none"),
		h.If(
			!selected,
			h.PostPartialOnClickQs(
				StorySidebar,
				h.NewQs("category", category.Path),
			),
		),
	)
}

var CachedStoryList = h.CachedPerKeyT4(time.Minute*5, func(category string, page int, limit int, fetchMorePath string) (string, h.GetElementFunc) {
	return fmt.Sprintf("%s-stories-%d-%d", category, page, limit), func() *h.Element {
		stories := news.GetStories(category, page, limit)
		return h.List(stories, func(item news.Story, index int) *h.Element {
			return h.Div(
				h.Attribute("hx-swap", "none"),
				h.PostPartialOnClickQs(Story, h.NewQs("item", fmt.Sprintf("%d", item.Id))),
				h.A(h.Href(item.Url)),
				h.Class("block p-2 bg-white rounded-md shadow cursor-pointer"),
				h.Div(
					h.Class("font-bold"),
					h.Text(item.Title),
				),
				h.Div(
					h.Class("text-sm text-gray-600"),
					h.Div(h.TextF("%s ", item.By), h.UnsafeRaw("&bull;"), h.TextF(" %s", timeformat.RelativeTime(item.Time))),
				),
				h.Div(
					h.Class("text-sm text-gray-600"),
					h.UnsafeRaw(fmt.Sprintf("%d upvotes &bull; %d comments", item.Score, item.Descendents)),
				),
				h.If(index == len(stories)-1, h.Div(
					h.Id("load-more"),
					h.Attribute("hx-swap", "beforeend"),
					h.HxTarget("#story-list"),
					h.Get(fetchMorePath, "intersect once"),
				)),
			)
		})
	}
})
