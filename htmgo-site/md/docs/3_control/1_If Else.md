**If / Else Statements**

If / else statements are useful when you want to conditionally render attributes or elements / components.

htmgo provides a couple of utilities to do so:

```go
h.If(condition, node)
h.Ternary(condition, node, node2)
h.ElementIf(condition, element) // this is neccessary if a method requires you to pass in *h.element
h.IfElse(condition, node, node2) //essentially an alias to h.Ternary
h.IfElseLazy(condition, func()node, func()node2) // useful for if something should only be called based on the condition
h.AttributeIf(condition, key string, value string) // adds an attribute if condition is true
h.ClassIf(condition, class string) // adds a class if condition is true
h.ClassX(classes, m.ClassMap{}) // allows you to include classes, but also render specific classes conditionally

```

**Examples:**

- Render `border-green-500` or `border-slate-400` conditionally

```go
h.ClassX("w-10 h-10 border rounded-full", map[string]bool {
				"border-green-500": task.CompletedAt != nil,
				"border-slate-400": task.CompletedAt == nil,
})
```



- Render an icon if the task is complete

```go
h.If(task.CompletedAt != nil, CompleteIcon())
```

- Render different elements based on a condition

```go
h.IfElse(editing, EditTaskForm(), ViewTask())
```

Note: This will execute both **EditTaskForm** and **ViewTask**, no matter if the condition is true or false, since a function is being called here.

If you do not want to call the function at all unless the condition is true, use **h.IfElseLazy**

```go
h.IfElseLazy(editing, EditTaskForm, ViewTask)
```

