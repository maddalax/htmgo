package dirwalk

import (
	"github.com/maddalax/htmgo/framework/h"
	"os"
	"path/filepath"
	"strings"
)

type Page struct {
	RoutePath string
	FilePath  string
	Parts     []string
}

func WalkPages(dir string) []Page {
	pages := make([]Page, 0)
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		name := info.Name()
		if !info.IsDir() && (strings.HasSuffix(name, ".md") || strings.HasSuffix(name, ".go")) {
			fullPath := strings.Replace(path, dir, "", 1)
			fullPath = strings.TrimSuffix(fullPath, ".md")
			pages = append(pages, Page{
				RoutePath: fullPath,
				FilePath:  path,
				Parts: h.Filter(strings.Split(fullPath, string(os.PathSeparator)), func(item string) bool {
					return item != ""
				}),
			})
		}
		return nil
	})
	return pages
}
