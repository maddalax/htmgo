package h

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

type Renderable interface {
	Render() *Node
}

func Ternary[T any](value bool, a T, b T) T {
	if value {
		return a
	}
	return b
}

func Map[T any, U any](items []T, fn func(T) U) []U {
	var result []U
	for _, item := range items {
		result = append(result, fn(item))
	}
	return result
}

func JsonSerialize(data any) string {
	serialized, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(serialized)
}

func GetQueryParam(ctx *fiber.Ctx, key string) string {
	value := ctx.Query(key)
	if value == "" {
		current := ctx.Get("Hx-Current-Url")
		if current != "" {
			u, err := url.Parse(current)
			if err == nil {
				return u.Query().Get(key)
			}
		}
	}
	return value
}
