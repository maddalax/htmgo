package reloader

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/astgen"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/copyassets"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/css"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/run"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/util"
	"log/slog"
	"strings"
	"sync"
	"time"
)

type Change struct {
	name string
	op   fsnotify.Op
}

func NewChange(event *fsnotify.Event) *Change {
	return &Change{name: event.Name, op: event.Op}
}

func (c *Change) Name() string {
	return c.name
}

func (c *Change) HasAnyPrefix(prefix ...string) bool {
	for _, s := range prefix {
		if strings.HasPrefix(c.name, s) {
			return true
		}
	}
	return false
}

func (c *Change) HasAnySuffix(suffix ...string) bool {
	for _, s := range suffix {
		if strings.HasSuffix(c.name, s) {
			return true
		}
	}
	return false
}

func (c *Change) IsGenerated() bool {
	return c.HasAnySuffix("generated.go")
}

func (c *Change) IsWrite() bool {
	return c.op == fsnotify.Write
}

func (c *Change) IsGo() bool {
	return c.HasAnySuffix(".go")
}

type Tasks struct {
	AstGen bool
	Run    bool
	Ent    bool
	Css    bool
}

func OnFileChange(version string, events []*fsnotify.Event) {
	now := time.Now()

	tasks := Tasks{}
	hasTask := false

	for _, event := range events {
		c := NewChange(event)

		if c.HasAnySuffix(".go~", ".css~") {
			continue
		}

		if c.IsGenerated() {
			continue
		}

		if c.IsGo() && c.HasAnyPrefix("pages/", "partials/") {
			tasks.AstGen = true
			hasTask = true
		}

		if c.IsGo() {
			tasks.Run = true
			tasks.Css = true
			hasTask = true
		}

		if c.HasAnySuffix(".md") {
			tasks.Run = true
			hasTask = true
		}

		if c.HasAnyPrefix("ent/schema") {
			tasks.Ent = true
			hasTask = true
		}

		// framework assets changed
		if c.HasAnySuffix("assets/dist/htmgo.js") {
			copyassets.CopyAssets()
			//tasks.Run = true
		}

		// something in public folder changed
		if c.HasAnyPrefix("assets/public/") {
			copyassets.CopyAssets()
		}

		if hasTask {
			slog.Info("file changed", slog.String("version", version), slog.String("file", c.Name()))
		}
	}

	if !hasTask {
		return
	}

	deps := make([]func() any, 0)

	if tasks.AstGen {
		go func() {
			util.Trace("generate ast", func() any {
				astgen.GenAst()
				return nil
			})
		}()
	}

	if tasks.Css {
		deps = append(deps, func() any {
			return util.Trace("generate css", func() any {
				css.GenerateCss()
				return nil
			})
		})
	}

	if tasks.Ent {
		deps = append(deps, func() any {
			return util.Trace("generate ent", func() any {
				run.EntGenerate()
				return nil
			})
		})
	}

	wg := sync.WaitGroup{}

	for _, dep := range deps {
		wg.Add(1)
		go func(dep func() any) {
			defer wg.Done()
			err := dep()
			if err != nil {
				fmt.Println(err)
			}
		}(dep)
	}

	wg.Wait()

	if tasks.Run {
		go run.Server()
	}

	slog.Info("reloaded in", slog.Duration("duration", time.Since(now)))

}
