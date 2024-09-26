package run

import (
	"github.com/maddalax/htmgo/cli/htmgo/internal/dirutil"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/process"
	"runtime"
)

func EntNewSchema(name string) {
	process.RunOrExit(process.NewRawCommand("", "GOWORK=off go run -mod=mod entgo.io/ent/cmd/ent new "+name))
}

func EntGenerate() {
	if dirutil.HasFileFromRoot("ent/schema") {
		if runtime.GOOS == "windows" {
			process.RunOrExit(process.NewRawCommand("ent-generate", "go generate ./ent"))
		} else {
			process.RunOrExit(process.NewRawCommand("ent-generate", "bash -c GOWORK=off go generate ./ent"))
		}
	}
}
