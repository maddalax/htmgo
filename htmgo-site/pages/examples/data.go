package examples

import "github.com/maddalax/htmgo/framework/h"

type Snippet struct {
	name           string
	description    string
	sidebarName    string
	path           string
	partial        h.PartialFunc
	externalRoute  string
	sourceCodePath string
	category       string
}

func SetSnippet(ctx *h.RequestContext, snippet *Snippet) {
	ctx.Set("snippet", snippet)
}

func GetSnippet(ctx *h.RequestContext) *Snippet {
	if ctx.Get("snippet") == nil {
		return nil
	}
	return ctx.Get("snippet").(*Snippet)
}
