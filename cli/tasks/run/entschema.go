package run

import "github.com/maddalax/htmgo/cli/tasks/process"

func EntNewSchema(name string) {
	process.RunOrExit("GOWORK=off go run -mod=mod entgo.io/ent/cmd/ent new " + name)
}

func EntGenerate() {
	process.RunOrExit("GOWORK=off go generate ./ent")
}
