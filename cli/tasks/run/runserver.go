package run

import "github.com/maddalax/htmgo/cli/tasks/process"

func Server(exitOnError bool) error {
	return process.Run("go run .", exitOnError)
}
