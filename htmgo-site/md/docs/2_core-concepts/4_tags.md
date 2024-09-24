**HTML Tags**

htmgo provides many methods to render html tags:

```go
h.Html(children ...Ren) *Element
h.Head(children ...Ren) *Element
h.Div(children ...Ren) *Element
h.Button(children ...Ren) *Element
h.P(children ...Ren) *Element
h.H1(children ...Ren) *Element
h.H2(children ...Ren) *Element
h.Tag(tag string, children ...Ren) *Element
... etc
```

All methods can be found in the `h` package in htmgo/framework

See [#conditionals](#control-if-else) for more information about conditionally rendering tags or attributes.
