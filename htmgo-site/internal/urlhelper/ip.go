package urlhelper

import (
	"net/http"
	"strings"
)

func GetClientIp(r *http.Request) string {
	// Try to get the real client IP from the 'CF-Connecting-IP' header
	if ip := r.Header.Get("CF-Connecting-IP"); ip != "" {
		return ip
	}

	// If not available, fall back to 'X-Forwarded-For'
	if ip := r.Header.Get("X-Forwarded-For"); ip != "" {
		return ip
	}

	// Otherwise, use the default remote address (this will be Cloudflare's IP)
	remote := r.RemoteAddr

	if strings.HasPrefix(remote, "[::1]") {
		return "localhost"
	}

	return remote
}
