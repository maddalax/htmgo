package run

import (
	"github.com/maddalax/htmgo/cli/htmgo/tasks/process"
)

func Setup() {
	process.RunOrExit(process.NewRawCommand("", "go mod download"))
	process.RunOrExit(process.NewRawCommand("", "go mod tidy"))
	MakeBuildable()
	EntGenerate()
}
