package formatter

import (
	"fmt"
	"github.com/maddalax/htmgo/tools/html-to-htmgo/htmltogo"
	"os"
	"path/filepath"
	"strings"
)

func FormatDir(dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("error reading dir: %s\n", err.Error())
		return
	}
	for _, file := range files {
		if file.IsDir() {
			FormatDir(filepath.Join(dir, file.Name()))
		} else {
			FormatFile(filepath.Join(dir, file.Name()))
		}
	}
}

func FormatFile(file string) {
	if !strings.HasSuffix(file, ".go") {
		return
	}

	fmt.Printf("formatting file: %s\n", file)

	source, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("error reading file: %s\n", err.Error())
		return
	}

	str := string(source)

	if !strings.Contains(str, "github.com/maddalax/htmgo/framework/h") {
		return
	}

	parsed := htmltogo.Indent(str)

	os.WriteFile(file, []byte(parsed), 0644)

	return
}
