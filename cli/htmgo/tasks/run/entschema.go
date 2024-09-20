package run

import "github.com/maddalax/htmgo/cli/htmgo/tasks/process"

func EntNewSchema(name string) {
	process.RunOrExit("bash -c GOWORK=off go run -mod=mod entgo.io/ent/cmd/ent new " + name)
}

func EntGenerate() {
	process.RunOrExit("bash -c GOWORK=off go generate ./ent")
}
