package formatter

import (
	"fmt"
	"github.com/maddalax/htmgo/tools/html-to-htmgo/htmltogo"
	"os"
	"strings"
)

func FormatFile(file string) {
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
