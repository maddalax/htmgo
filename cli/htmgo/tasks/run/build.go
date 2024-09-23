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
	copyassets.CopyAssets()
	astgen.GenAst(process.ExitOnError)
	css.GenerateCss(process.ExitOnError)

	process.RunOrExit("rm -rf ./dist")
	process.RunOrExit("mkdir -p ./dist")

	if os.Getenv("SKIP_GO_BUILD") != "1" {
		process.RunOrExit(fmt.Sprintf("go build -o ./dist"))
	}

	fmt.Printf("Executable built at %s\n", process.GetPathRelativeToCwd("dist"))
}
