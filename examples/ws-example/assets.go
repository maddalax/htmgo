//go:build !prod
// +build !prod

package main

import (
	"io/fs"
	"ws-example/internal/embedded"
)

func GetStaticAssets() fs.FS {
	return embedded.NewOsFs()
}
