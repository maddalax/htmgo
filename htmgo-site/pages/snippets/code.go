package snippets

import (
	"fmt"
	"github.com/maddalax/htmgo/framework/h"
	"htmgo-site/ui"
	"os"
	"reflect"
	"runtime"
)

func getFunctionFilePath(fn interface{}) string {
	// Get the function pointer using reflection
	ptr := reflect.ValueOf(fn).Pointer()
	// Get the file path and line number using runtime
	fnInfo := runtime.FuncForPC(ptr)
	if fnInfo == nil {
		return ""
	}
	file, _ := fnInfo.FileLine(ptr)
	return file
}

func GetGithubPath(path string) string {
	return fmt.Sprintf("https://github.com/maddalax/htmgo/tree/master/htmgo-site/partials%s.go", path)
}

func RenderCodeToString(partial h.PartialFunc) *h.Element {
	path := getFunctionFilePath(partial)
	if path == "" {
		return h.Empty()
	}
	bytes, err := os.ReadFile(path)
	if err != nil {
		return h.Empty()
	}
	return ui.CodeSnippet(string(bytes), "border-radius: 0.5rem;")
}
