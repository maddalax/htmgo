package h

import (
	"net/url"
	"strings"
)

type Qs struct {
	m map[string]string
}

func NewQs(pairs ...string) *Qs {
	q := &Qs{
		m: make(map[string]string),
	}
	if len(pairs)%2 != 0 {
		pairs = append(pairs, "")
	}
	for i := 0; i < len(pairs); i++ {
		q.m[pairs[i]] = pairs[i+1]
		i++
	}
	return q
}

func (q *Qs) Add(key string, value string) *Qs {
	q.m[key] = value
	return q
}

func (q *Qs) Remove(key string) *Qs {
	delete(q.m, key)
	return q
}

func (q *Qs) ToString() string {
	builder := strings.Builder{}
	index := 0
	for k, v := range q.m {
		builder.WriteString(k)
		if v != "" {
			builder.WriteString("=")
			builder.WriteString(v)
		}
		if index < len(q.m)-1 {
			builder.WriteString("&")
		}
		index++
	}
	return builder.String()
}

// GetQueryParam returns the value of the given query parameter from the request URL.
// There are two layers of priority:
// 1. The query parameter in the URL
// 2. The current browser URL
// If the query parameter is not found in the URL from the *RequestContext, it will fall back to the current browser URL if set.
// The URL from the *RequestContext would normally be the url from an XHR request through htmx,
// which is not the current browser url a visitor may be on.
func GetQueryParam(ctx *RequestContext, key string) string {
	value, ok := ctx.Request.URL.Query()[key]
	if value == nil || !ok {
		current := ctx.currentBrowserUrl
		if current != "" {
			u, err := url.Parse(current)
			if err == nil {
				return u.Query().Get(key)
			}
		}
	}
	if len(value) == 0 {
		return ""
	}
	return value[0]
}

// SetQueryParams sets the query parameters of the given URL.
// Given the *Qs passed in, it will set the query parameters of the URL to the given values.
// If the value does not exist in *QS, it will remain untouched.
// If the value is an empty string, it will be removed from the query parameters.
// If the value is not an empty string, it will be set to the given value.
func SetQueryParams(href string, qs *Qs) string {
	u, err := url.Parse(href)
	if err != nil {
		return href
	}
	q := u.Query()
	for key, value := range qs.m {
		if value == "" {
			q.Del(key)
		} else {
			q.Set(key, value)
		}
	}
	u.RawQuery = q.Encode()
	return u.String()
}
