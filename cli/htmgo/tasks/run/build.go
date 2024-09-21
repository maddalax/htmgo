package run

import (
	"fmt"
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

	process.RunOrExit("env go build -o ./dist .")
	process.RunOrExit("go build -o ./dist/app .")

	fmt.Printf("Executable built at %s\n", process.GetPathRelativeToCwd("dist"))
}
