## Htmgo Format

htmgo has a built-in formatter that can be used to format htmgo element blocks.

It is available through the 'htmgo' cli tool that is installed with htmgo.

**Note:** if you have previously installed htmgo, you will need to run `GOPROXY=direct go install github.com/maddalax/htmgo/cli/htmgo@latest` to update the cli tool.

<br>
To use it, run the following command:

```bash
// format all .go files in the current directory recursively
htmgo format .

// format the file specified
htmgo format ./my-file.go
```

This will format all htmgo element blocks in your project.

**Example:**

```go
h.Div(
	h.Class("flex gap-2"), h.Text("hello"), h.Text("world"),
)
```

**Output:**

```go
h.Div(
	h.Class("flex gap-2"),
	h.Text("hello"),
	h.Text("world"),
)
```

## Running htmgo format on save

### Jetbrains IDE's

1. Go to Settings -> Tools -> File Watchers -> + custom

2. Set the following values:

```yaml
Name: htmgo format
File Type: Go
Scope: Current File
Program: htmgo
Arguments: format $FilePath$
Output paths to refresh: $FilePath$
Working directory: $ProjectFileDir$
```

3. Save the file watcher and ensure it is enabled

4. Go to `Settings -> Tools -> Actions On Save` and ensure the `htmgo format` action is enabled
