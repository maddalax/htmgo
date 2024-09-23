package run

import (
	"fmt"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/astgen"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/copyassets"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/css"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/process"
	"os"
)

func Build() {
	css.DownloadTailwindCli()

	astgen.GenAst(process.ExitOnError)
	css.GenerateCss(process.ExitOnError)

	copyassets.CopyAssets()

	process.RunOrExit("rm -rf ./dist")
	process.RunOrExit("mkdir -p ./dist")

	flags := ""
	if os.Getenv("CGO_ENABLED") == "1" {
		flags = `-a -ldflags '-linkmode external -extldflags "-static"' `
	}

	process.RunOrExit(fmt.Sprintf("go build -o ./dist/app %s.", flags))
	process.RunOrExit(fmt.Sprintf("go build -o ./dist/app %s.", flags))

	fmt.Printf("Executable built at %s\n", process.GetPathRelativeToCwd("dist"))
}
