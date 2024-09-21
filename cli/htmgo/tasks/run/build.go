package run

import (
	"github.com/maddalax/htmgo/cli/htmgo/tasks/astgen"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/copyassets"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/css"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/process"
)

func Build() {
	css.DownloadTailwindCli()

	astgen.GenAst(process.ExitOnError)
	css.GenerateCss(process.ExitOnError)

	copyassets.CopyAssets()

	process.RunOrExit("rm -rf ./dist")
	process.RunOrExit("mkdir -p ./dist")

	process.RunOrExit("env GOOS=linux GOARCH=amd64 go build -o ./dist/app-linux-amd64 .")
	process.RunOrExit("go build -o ./dist/app .")

	process.RunOrExit("echo \"Build successful\"")
}
