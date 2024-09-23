package run

import (
	"github.com/maddalax/htmgo/cli/htmgo/internal/dirutil"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/process"
)

func EntNewSchema(name string) {
	process.RunOrExit("GOWORK=off go run -mod=mod entgo.io/ent/cmd/ent new " + name)
}

func EntGenerate() {
	if dirutil.HasFileFromRoot("ent/schema") {
		process.RunOrExit("GOWORK=off go generate ./ent")
	}
}
