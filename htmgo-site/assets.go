//go:build !prod
// +build !prod

package main

import (
	"htmgo-site/internal/embedded"
	"io/fs"
)

func GetStaticAssets() fs.FS {
	return embedded.NewOsFs()
}

func GetMarkdownAssets() fs.FS {
	return embedded.NewOsFs()
}
