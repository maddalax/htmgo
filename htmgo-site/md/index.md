**introduction:**

htmgo is a lightweight pure go way to build interactive websites / web applications using go & htmx.

By combining the speed & simplicity of go + hypermedia attributes ([htmx](https://htmx.org)) to add interactivity to websites, all conveniently wrapped in pure go, you can build simple, fast, interactive websites without touching javascript. All compiled to a **single deployable binary**.

```go
func IndexPage(ctx *h.RequestContext) *h.Page {
  now := time.Now()
  return h.NewPage(
    h.Div(
      h.Class("flex gap-2"),
      h.TextF("the current time is %s", now.String())
    )
  )
}
```

**core features:**

1. deployable single binary
2. live reload (rebuilds css, go, ent schema, and routes upon change)
3. automatic page and partial registration based on file path
4. built in tailwindcss support, no need to configure anything by default
5. custom [htmx extensions](https://github.com/maddalax/htmgo/tree/master/framework/assets/js/htmxextensions) to reduce boilerplate with common tasks

------

**what can be built with htmgo?**

Most web applications can be built with htmgo, including but not limited to:

- Traditional business CRUD applications
- blogs 
- documentation sites
- consumer facing websites 
- internal tools
- and more

<br>

For a more detailed overview of when you should use hypermedia to build web applications, see [when-to-use-hypermedia](https://htmx.org/essays/when-to-use-hypermedia/) from htmx.org.

Interested in some examples? Check out [examples](/examples).
