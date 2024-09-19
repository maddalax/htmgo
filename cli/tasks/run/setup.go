package run

import (
	"github.com/maddalax/htmgo/cli/tasks/astgen"
	"github.com/maddalax/htmgo/cli/tasks/copyassets"
	"github.com/maddalax/htmgo/cli/tasks/css"
	"github.com/maddalax/htmgo/cli/tasks/process"
)

func Setup() {
	process.RunOrExit("go mod download")
	process.RunOrExit("go mod tidy")
	copyassets.CopyAssets()
	_ = astgen.GenAst(process.ExitOnError)
	_ = css.GenerateCss(process.ExitOnError)
	EntGenerate()
}
