package run

import "github.com/maddalax/htmgo/cli/htmgo/tasks/process"

func Server(flags ...process.RunFlag) error {
	return process.Run(process.NewRawCommand("run-server", "go run .", flags...))
}
