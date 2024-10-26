//go:build !prod
// +build !prod

package main

import (
	"hackernews/internal/embedded"
	"io/fs"
)

func GetStaticAssets() fs.FS {
	return embedded.NewOsFs()
}
