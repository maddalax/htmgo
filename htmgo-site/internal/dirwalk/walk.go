package dirwalk

import (
	"github.com/maddalax/htmgo/framework/h"
	"io/fs"
	"os"
	"strings"
)

type Page struct {
	RoutePath string
	FilePath  string
	Parts     []string
}

func WalkPages(dir string, system fs.FS) []Page {
	pages := make([]Page, 0)
	fs.WalkDir(system, dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		name := d.Name()
		if !d.IsDir() && (strings.HasSuffix(name, ".md") || strings.HasSuffix(name, ".go")) {
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
