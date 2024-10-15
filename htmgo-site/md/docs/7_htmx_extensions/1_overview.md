## HTMX Extensions

htmgo provides a few extra htmx extensions to make common tasks easier.
Some of these extensions are optional, and some of these are required for htmgo to work correctly.

The following extensions are provided by htmgo:
- [Trigger Children](#htmx-extensions-trigger-children)
- [Mutation Error](#htmx-extensions-mutation-error)
- [SSE](#pushing-data-server-sent-events)
- [Path Deps](https://github.com/bigskysoftware/htmx-extensions/blob/main/src/path-deps/README.md)

Default extensions should be included in your project by adding the following attribute to your html tag.
```go
h.Html(
    h.HxExtension(h.BaseExtensions())
)
```

If you need to combine multiple extensions, you can use:

```go
h.HxExtensions(h.BaseExtensions(), "my-extension"),
```
or
```go
h.JoinExtensions(
    h.HxExtension("sse"),
    h.HxExtension("my-extension"),
),
```


**Important**: h.BaseExtensions will add the the 'htmgo' extension, which is a required extension for inline scripts to work properly, please always include it in your project.

