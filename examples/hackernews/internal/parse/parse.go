package parse

import "strconv"

func MustParseInt(s string, fallback int) int {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return fallback
	}
	return int(v)
}
