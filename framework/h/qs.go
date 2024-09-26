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
		return q
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
		builder.WriteString("=")
		builder.WriteString(v)
		if index < len(q.m)-1 {
			builder.WriteString("&")
		}
		index++
	}
	return builder.String()
}

func GetQueryParam(ctx *RequestContext, key string) string {
	value := ctx.QueryParam(key)
	if value == "" {
		current := ctx.currentBrowserUrl
		if current != "" {
			u, err := url.Parse(current)
			if err == nil {
				return u.Query().Get(key)
			}
		}
	}
	return value
}

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
