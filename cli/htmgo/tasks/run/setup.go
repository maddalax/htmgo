package run

import (
	"github.com/maddalax/htmgo/cli/htmgo/tasks/astgen"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/copyassets"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/css"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/process"
)

func Setup() {
	process.RunOrExit("go mod download")
	process.RunOrExit("go mod tidy")

	copyassets.CopyAssets()
	astgen.GenAst(process.ExitOnError)
	css.GenerateCss(process.ExitOnError)

	EntGenerate()
}
