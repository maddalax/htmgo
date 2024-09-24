## **Introduction**

htmgo is a lightweight pure go way to build interactive websites / web applications using go & htmx.
We give you the utilities to build html using pure go code in a reusable way (go functions are components) while also providing htmx functions to add interactivity to your app.

```go
func DocsPage(ctx *h.RequestContext) *h.Page {
	assets := ctx.Get("embeddedMarkdown").(*embed.FS)
	pages := dirwalk.WalkPages("md/docs", assets)

	return h.NewPage(base.RootPage(
		h.Div(
			h.Class("flex flex-col md:flex-row gap-4 justify-center mb-12"),
			partials.DocSidebar(pages),
			h.Div(
				h.Class("flex flex-col justify-center items-center mt-6"),
				h.List(pages, func(page *dirwalk.Page, index int) *h.Element {
					return h.Div(
						h.Class("border-b border-b-slate-300 w-full pb-8 mb-8"),
						MarkdownContent(ctx, 
                            page.FilePath, 
                            partials.CreateAnchor(page.Parts)),
					)
				}),
			),
		),
	))
}
```

**This site was written with htmgo!**

<br>

**Quick overview**

1. Server side rendered html, deploy as a single binary

2. Built in live reloading

3. Built in support for various libraries such as tailwindcss, htmx

4. Go functions are components, no special syntax neccessary to learn

5. Many composable utility functions to streamline development and reduce boilerplate

   ```go
   func ChangeTab(ctx *h.RequestContext) *h.Partial {
   	service := tasks.NewService(ctx.ServiceLocator())
   	list, _ := service.List()
   
   	tab := ctx.QueryParam("tab")
   
   	return h.SwapManyPartialWithHeaders(ctx,
   		h.PushQsHeader(ctx, h.NewQs("tab", tab)),
   		List(list, tab),
   		Footer(list, tab),
   	)
   }
   ```

   Example: **h.SwapManyPartialWithHeaders** to swap out multiple elements on the page with your response, as well as set a new query string parameter.



<br>

See [#core-concepts](#core-concepts-pages) for more information.
