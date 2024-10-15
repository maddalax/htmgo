package timeformat

import (
	"fmt"
	"time"
)

func ParseUnix(t int64) time.Time {
	return time.UnixMilli(t * 1000)
}

func RelativeTime(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)

	var pluralize = func(s string) string {
		if s[0] == '1' {
			return s[:len(s)-5] + " ago"
		}
		return s
	}

	switch {
	case diff < time.Minute:
		return "just now"
	case diff < time.Hour:
		return pluralize(fmt.Sprintf("%d minutes ago", int(diff.Minutes())))
	case diff < time.Hour*24:
		return pluralize(fmt.Sprintf("%d hours ago", int(diff.Hours())))
	case diff < time.Hour*24*7:
		return pluralize(fmt.Sprintf("%d days ago", int(diff.Hours()/24)))
	case diff < time.Hour*24*30:
		return pluralize(fmt.Sprintf("%d weeks ago", int(diff.Hours()/(24*7))))
	case diff < time.Hour*24*365:
		return pluralize(fmt.Sprintf("%d months ago", int(diff.Hours()/(24*30))))
	default:
		return pluralize(fmt.Sprintf("%d years ago", int(diff.Hours()/(24*365))))
	}
}
