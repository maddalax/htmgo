> [!WARNING]
> htmgo is in alpha release and active development. API's may have breaking changes between versions. Please report any issues on GitHub.

## **htmgo**

### build simple and scalable systems with go + htmx

-------
[![Go Report Card](https://goreportcard.com/badge/github.com/maddalax/htmgo)](https://goreportcard.com/report/github.com/maddalax/htmgo)
![Build](https://github.com/maddalax/htmgo/actions/workflows/run-framework-tests.yml/badge.svg)


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
5. plugin architecture to include optional plugins to streamline development, such as http://entgo.io
6. custom [htmx extensions](https://github.com/maddalax/htmgo/tree/b610aefa36e648b98a13823a6f8d87566120cfcc/framework/assets/js/htmxextensions) to reduce boilerplate with common tasks

**get started:**

View documentation on [htmgo.dev](https://htmgo.dev/docs).

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=maddalax/htmgo&type=Date)](https://star-history.com/#maddalax/htmgo&Date)
