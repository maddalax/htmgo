### Tailwind intellisense

Tailwind's language server allows you to specify custom configuration on what it should match to start giving you tailwind intellisense.


![](/public/tailwind-intellisense.png)

To make this work, you will need to update the tailwind lsp config with the config below:

Main thing to note here is
1. "go" is added to the includeLanguages list
2. classRegex is updated to match the tailwind classes in the go code.

### Jetbrains IDE's (GoLand)
```json
{
  "includeLanguages": {
    "go": "html"
  },
  "experimental": {
    "configFile": null,
    "classRegex": [
      ["Class\\(([^)]*)\\)", "[\"`]([^\"`]*)[\"`]"],
      ["ClassX\\(([^)]*)\\)", "[\"`]([^\"`]*)[\"`]"],
      ["ClassIf\\(([^)]*)\\)", "[\"`]([^\"`]*)[\"`]"],
      ["Classes\\(([^)]*)\\)", "[\"`]([^\"`]*)[\"`]"]
    ]
  }
}
```
To find this configuration in GoLand you can go to `Settings -> Languages & Frameworks -> Style Sheets -> Tailwind CSS` and update the configuration there.
These changes are additive, add these options to your existing tailwind lsp config, instead of replacing the entire file.

See more: https://github.com/tailwindlabs/tailwindcss/issues/7553#issuecomment-735915659

<br>

### Visual Studio Code
For VSCode, you should be able to update your settings.json with the following values:

```json
{
  "tailwindCSS.includeLanguages": {
    "go": "html"
  },
  "tailwindCSS.experimental.classRegex": [
    ["Class\\(([^)]*)\\)", "[\"`]([^\"`]*)[\"`]"],
    ["ClassX\\(([^)]*)\\)", "[\"`]([^\"`]*)[\"`]"],
    ["ClassIf\\(([^)]*)\\)", "[\"`]([^\"`]*)[\"`]"],
    ["Classes\\(([^)]*)\\)", "[\"`]([^\"`]*)[\"`"]
  ]
}
```
