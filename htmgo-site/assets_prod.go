//go:build prod
// +build prod

package main

import (
	"embed"
	"io/fs"
)

//go:embed assets/dist/*
var staticAssets embed.FS

//go:embed md/*
var markdownAssets embed.FS

func GetStaticAssets() fs.FS {
	return staticAssets
}

func GetMarkdownAssets() fs.FS {
	return markdownAssets
}
