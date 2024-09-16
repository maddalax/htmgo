package main

import (
	"flag"
	"fmt"
	"github.com/maddalax/htmgo/cli/tasks/astgen"
	"github.com/maddalax/htmgo/cli/tasks/copyassets"
	"github.com/maddalax/htmgo/cli/tasks/css"
	"github.com/maddalax/htmgo/cli/tasks/downloadtemplate"
	"github.com/maddalax/htmgo/cli/tasks/process"
	"github.com/maddalax/htmgo/cli/tasks/reloader"
	"github.com/maddalax/htmgo/cli/tasks/run"
	"os"
	"strings"
)

func main() {
	done := RegisterSignals()

	commandMap := make(map[string]*flag.FlagSet)
	commands := []string{"template", "run", "watch", "build", "setup", "css"}

	for _, command := range commands {
		commandMap[command] = flag.NewFlagSet(command, flag.ExitOnError)
	}

	if len(os.Args) < 2 {
		fmt.Println(fmt.Sprintf("Usage: htmgo [%s]", strings.Join(commands, " | ")))
		os.Exit(1)
	}

	c := commandMap[os.Args[1]]

	if c == nil {
		fmt.Println(fmt.Sprintf("Usage: htmgo [%s]", strings.Join(commands, " | ")))
		os.Exit(1)
		return
	}

	err := c.Parse(os.Args[2:])

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
		return
	}

	taskName := os.Args[1]

	if taskName == "watch" {
		astgen.GenAst(true)
		css.GenerateCss(true)
		go func() {
			_ = run.Server(true)
		}()
		startWatcher(reloader.OnFileChange)
	} else {
		if taskName == "setup" {
			process.RunOrExit("go mod download")
			process.RunOrExit("go mod tidy")
			copyassets.CopyAssets()
			_ = astgen.GenAst(true)
			_ = css.GenerateCss(true)
		}
		if taskName == "css" {
			_ = css.GenerateCss(true)
		}
		if taskName == "ast" {
			_ = astgen.GenAst(true)
		}
		if taskName == "run" {
			_ = astgen.GenAst(true)
			_ = css.GenerateCss(true)
			_ = run.Server(true)
		}
		if taskName == "template" {
			downloadtemplate.DownloadTemplate("./my-app")
		}
	}

	<-done
	fmt.Println("Cleanup complete. Exiting.")
}
