package h

import "os"

func IsWatchMode() bool {
	return os.Getenv("WATCH_MODE") == "true"
}

func IsDevelopment() bool {
	return os.Getenv("ENV") == "development"
}

func IsProduction() bool {
	return os.Getenv("ENV") == "production"
}
