package run

import "github.com/maddalax/htmgo/cli/htmgo/tasks/process"

func Server(flags ...process.RunFlag) error {
	return process.Run("go run .", flags...)
}
