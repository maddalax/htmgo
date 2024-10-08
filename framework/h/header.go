package h

import (
	"github.com/maddalax/htmgo/framework/hx"
	"net/url"
)

type Headers = map[string]string

func ReplaceUrlHeader(url string) *Headers {
	return NewHeaders(hx.ReplaceUrlHeader, url)
}

func PushUrlHeader(url string) *Headers {
	return NewHeaders(hx.PushUrlHeader, url)
}

func PushQsHeader(ctx *RequestContext, qs *Qs) *Headers {
	parsed, err := url.Parse(ctx.CurrentBrowserUrl)
	if err != nil {
		return NewHeaders()
	}
	return NewHeaders(hx.ReplaceUrlHeader, SetQueryParams(parsed.Path, qs))
}

func CombineHeaders(headers ...*Headers) *Headers {
	m := make(Headers)
	for _, h := range headers {
		for k, v := range *h {
			m[k] = v
		}
	}
	return &m
}

func CurrentPath(ctx *RequestContext) string {
	current := ctx.Request.Header.Get(hx.CurrentUrlHeader)
	parsed, err := url.Parse(current)
	if err != nil {
		return ""
	}
	return parsed.Path
}

func NewHeaders(headers ...string) *Headers {
	if len(headers)%2 != 0 {
		return &Headers{}
	}
	m := make(Headers)
	for i := 0; i < len(headers); i++ {
		m[headers[i]] = headers[i+1]
		i++
	}
	return &m
}
