package run

import (
	"fmt"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/astgen"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/copyassets"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/css"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/process"
	"os"
)

func MakeBuildable() {
	copyassets.CopyAssets()
	astgen.GenAst(process.ExitOnError)
	css.GenerateCss(process.ExitOnError)
}

func Build() {
	MakeBuildable()

	process.RunOrExit(process.NewRawCommand("", "mkdir -p ./dist"))

	if os.Getenv("SKIP_GO_BUILD") != "1" {
		process.RunOrExit(process.NewRawCommand("", fmt.Sprintf("go build -tags prod -o ./dist")))
	}

	fmt.Printf("Executable built at %s\n", process.GetPathRelativeToCwd("dist"))
}
