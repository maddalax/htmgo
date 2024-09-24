**Attributes**

Attributes are one of the main ways we can add interactivity to the pages with [htmx](http://htmx.org). If you have not read over the htmx documentation, please do so before continuing. 

htmgo provides many methods to add attributes

```go
h.Class(string)
h.ClassX(string, h.ClassMap)
h.Href(string)
h.Attribute(key, value)
h.AttributeIf(condition, key, value)
h.AttributePairs(values...string) // set multiple attributes, must be an even number of parameters
h.Attributes(h.AttributeMap) // set multiple attributes as key/value pairs
h.Id(string)
h.Trigger(hx.Trigger) //htmx trigger using additional functions to construct the trigger
h.TriggerString(string) // htmx trigger in pure htmx string form

```



