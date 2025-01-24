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
	css.GenerateCss(process.ExitOnError)
	astgen.GenAst(process.ExitOnError)
}

func Build() {
	MakeBuildable()

	_ = os.RemoveAll("./dist")

	err := os.Mkdir("./dist", 0755)

	if err != nil {
		fmt.Println("Error creating dist directory", err)
		os.Exit(1)
	}

	if os.Getenv("SKIP_GO_BUILD") != "1" {
		process.RunOrExit(process.NewRawCommand("", fmt.Sprintf("go build -tags prod -o ./dist")))
	}

	fmt.Printf("Executable built at %s\n", process.GetPathRelativeToCwd("dist"))
}
