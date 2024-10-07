//go:build !prod
// +build !prod

package main

import (
	"io/fs"
	"starter-template/internal/embedded"
)

func GetStaticAssets() fs.FS {
	return embedded.NewOsFs()
}
