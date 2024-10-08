//go:build !prod
// +build !prod

package main

import (
	"io/fs"
	"sse-with-state/internal/embedded"
)

func GetStaticAssets() fs.FS {
	return embedded.NewOsFs()
}
