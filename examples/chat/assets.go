//go:build !prod
// +build !prod

package main

import (
	"chat/internal/embedded"
	"io/fs"
)

func GetStaticAssets() fs.FS {
	return embedded.NewOsFs()
}
