package h

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/url"
)

type Ren interface {
	Render() string
}

func Ternary[T any](value bool, a T, b T) T {
	if value {
		return a
	}
	return b
}

func JsonSerialize(data any) string {
	serialized, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(serialized)
}

func GetQueryParam(ctx echo.Context, key string) string {
	value := ctx.QueryParam(key)
	if value == "" {
		current := ctx.Request().Header.Get("Hx-Current-Url")
		if current != "" {
			u, err := url.Parse(current)
			if err == nil {
				return u.Query().Get(key)
			}
		}
	}
	return value
}

func SetQueryParams(href string, qs map[string]string) string {
	u, err := url.Parse(href)
	if err != nil {
		return href
	}
	q := u.Query()
	for key, value := range qs {
		if value == "" {
			q.Del(key)
		} else {
			q.Set(key, value)
		}
	}
	u.RawQuery = q.Encode()
	return u.String()
}
