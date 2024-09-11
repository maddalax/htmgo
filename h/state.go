package h

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"mhtml/database"
)

func SessionSet(ctx *fiber.Ctx, key string, value string) error {
	sessionId := getSessionId(ctx)
	if sessionId == "" {
		return nil
	}
	return database.HSet(fmt.Sprintf("session:%s", sessionId), key, value)
}

func SessionIncr(ctx *fiber.Ctx, key string) int64 {
	sessionId := getSessionId(ctx)
	if sessionId == "" {
		return 0
	}
	return database.HIncr(fmt.Sprintf("session:%s", sessionId), key)
}

func SessionGet[T any](ctx *fiber.Ctx, key string) *T {
	sessionId := getSessionId(ctx)
	if sessionId == "" {
		return nil
	}
	return database.HGet[T](fmt.Sprintf("session:%s", sessionId), key)
}

func getSessionId(ctx *fiber.Ctx) string {
	return ctx.Cookies("mhtml-session")
}
