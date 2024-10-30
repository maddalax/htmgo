package urlhelper

import (
	"github.com/maddalax/htmgo/framework/h"
	"net/url"
)

func ToAbsoluteUrl(ctx *h.RequestContext, path string) string {
	// Define the relative path you want to add
	relativePath := path

	// Parse the current request URL
	currentURL := ctx.Request.URL

	// Set scheme and host from the request to create an absolute URL
	scheme := "http"
	if ctx.Request.TLS != nil {
		scheme = "https"
	}
	currentURL.Host = ctx.Request.Host
	currentURL.Scheme = scheme

	// Combine the base URL with the relative path
	absoluteURL := currentURL.ResolveReference(&url.URL{Path: relativePath})

	// Output the full absolute URL
	return absoluteURL.String()
}
