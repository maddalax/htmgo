package snippets

import (
	"bytes"
	"fmt"
	"github.com/maddalax/htmgo/framework/h"
	"htmgo-site/ui"
	"io"
	"log/slog"
	"net/http"
	"time"
)

func GetGithubPath(path string) string {
	return fmt.Sprintf("https://github.com/maddalax/htmgo/tree/master/htmgo-site/partials%s.go", path)
}

func GetGithubRawPath(path string) string {
	return fmt.Sprintf("https://raw.githubusercontent.com/maddalax/htmgo/master/htmgo-site/partials%s.go", path)
}

var RenderCodeToStringCached = h.CachedPerKeyT(time.Minute*30, func(snippet *Snippet) (string, h.GetElementFunc) {
	return snippet.path, func() *h.Element {
		return renderCodeToString(snippet)
	}
})

func renderCodeToString(snippet *Snippet) *h.Element {
	url := GetGithubRawPath(snippet.path)
	slog.Info("getting snippet source code", slog.String("url", url))
	resp, err := http.Get(url)
	if err != nil {
		return h.Empty()
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return h.Empty()
	}
	out := bytes.NewBuffer(nil)
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return h.Empty()
	}
	return ui.CodeSnippet(out.String(), "border-radius: 0.5rem;")
}
