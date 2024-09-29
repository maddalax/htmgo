**Rendering Raw Html**

In some cases, you may want to render raw HTML instead of using htmgo's functions.
This can be done by using the following methods:
```go
h.UnsafeRaw(string)
h.UnsafeRawF(string, ...interface{})
h.UnsafeRawScript(string)
```

Usage:
```go

h.UnsafeRaw("<div>Raw HTML</div>")
h.UnsafeRawF("<div>%s</div>", "Raw HTML")
h.UnsafeRawScript("alert('Hello World')")
```

Important: Be careful when using these methods, these methods do not escape the HTML content 
and should **never** be used with user input unless you have sanitized the input.

Sanitizing input can be done using the `html.EscapeString` function or by using https://github.com/microcosm-cc/bluemonday.