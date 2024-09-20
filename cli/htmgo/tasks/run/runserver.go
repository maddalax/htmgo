package run

import "github.com/maddalax/htmgo/cli/tasks/process"

func Server(flags ...process.RunFlag) error {
	return process.Run("go run .", flags...)
}
