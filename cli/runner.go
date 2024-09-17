package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/maddalax/htmgo/cli/tasks/astgen"
	"github.com/maddalax/htmgo/cli/tasks/css"
	"github.com/maddalax/htmgo/cli/tasks/downloadtemplate"
	"github.com/maddalax/htmgo/cli/tasks/process"
	"github.com/maddalax/htmgo/cli/tasks/reloader"
	"github.com/maddalax/htmgo/cli/tasks/run"
	"log/slog"
	"os"
	"strings"
)

func main() {
	done := RegisterSignals()

	commandMap := make(map[string]*flag.FlagSet)
	commands := []string{"template", "run", "watch", "build", "setup", "css", "schema"}

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

	slog.SetLogLoggerLevel(getLogLevel())

	taskName := os.Args[1]

	slog.Debug("Running task: %s", taskName)
	slog.Debug("working dir %s", process.GetWorkingDir())

	if taskName == "watch" {
		os.Setenv("ENV", "development")
		os.Setenv("WATCH_MODE", "true")
		astgen.GenAst(true)
		css.GenerateCss(true)
		run.EntGenerate()
		go func() {
			_ = run.Server(true)
		}()
		startWatcher(reloader.OnFileChange)
	} else {
		if taskName == "schema" {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter entity name:")
			text, _ := reader.ReadString('\n')
			text = strings.TrimSuffix(text, "\n")
			run.EntNewSchema(text)
		}
		if taskName == "generate" {
			run.EntGenerate()
			astgen.GenAst(true)
		}
		if taskName == "setup" {
			run.Setup()
		} else if taskName == "css" {
			_ = css.GenerateCss(true)
		} else if taskName == "ast" {
			_ = astgen.GenAst(true)
		} else if taskName == "run" {
			_ = astgen.GenAst(true)
			_ = css.GenerateCss(true)
			_ = run.Server(true)
		} else if taskName == "template" {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("What would you like to call your new app?: ")
			text, _ := reader.ReadString('\n')
			text = strings.TrimSuffix(text, "\n")
			text = strings.ReplaceAll(text, " ", "-")
			text = strings.ToLower(text)
			downloadtemplate.DownloadTemplate(fmt.Sprintf("./%s", text))
		} else if taskName == "build" {
			run.Build()
		} else {
			fmt.Println(fmt.Sprintf("Usage: htmgo [%s]", strings.Join(commands, " | ")))
		}
		os.Exit(0)
	}

	<-done
}
