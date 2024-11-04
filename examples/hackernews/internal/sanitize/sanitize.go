package sanitize

import "github.com/microcosm-cc/bluemonday"

var p = bluemonday.UGCPolicy()

func Sanitize(text string) string {
	return p.Sanitize(text)
}
