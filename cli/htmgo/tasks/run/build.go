package run

import (
	"github.com/maddalax/htmgo/cli/htmgo/tasks/astgen"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/css"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/process"
)

func Build() {
	astgen.GenAst(process.ExitOnError)
	css.GenerateCss(process.ExitOnError)
	process.RunOrExit("rm -rf ./dist")
	process.RunOrExit("mkdir -p ./dist")

	//process.RunOrExit("mkdir -p ./dist/assets/dist")

	//dirutil.CopyDir(
	//	"./assets/dist",
	//	"./dist/assets/dist",
	//	func(path string, exists bool) bool {
	//		return true
	//	},
	//)

	process.RunOrExit("env GOOS=linux GOARCH=amd64 go build -o ./dist .")
	process.RunOrExit("echo \"Build successful\"")
}
