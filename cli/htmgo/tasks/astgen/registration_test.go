package astgen

import (
	"fmt"
	"github.com/maddalax/htmgo/cli/htmgo/internal/dirutil"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/process"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"
)

func TestAstGen(t *testing.T) {
	t.Parallel()

	workingDir, err := filepath.Abs("./project-sample")

	assert.NoError(t, err)
	process.SetWorkingDir(workingDir)
	assert.NoError(t, os.Chdir(workingDir))

	err = dirutil.DeleteDir(filepath.Join(process.GetWorkingDir(), "__htmgo"))
	assert.NoError(t, err)
	err = process.Run(process.NewRawCommand("", "go build ."))
	assert.Error(t, err)
	err = GenAst()
	assert.NoError(t, err)

	go func() {
		// project was buildable after astgen, confirmed working
		err = process.Run(process.NewRawCommand("server", "go run ."))
		assert.NoError(t, err)
	}()

	time.Sleep(time.Second * 1)

	urls := []string{
		"/astgen-project-sample/partials.CountersPartial",
		"/",
		"/astgen-project-sample/pages.TestPartial",
	}

	defer func() {
		serverProcess := process.GetProcessByName("server")
		assert.NotNil(t, serverProcess)
		process.KillProcess(*serverProcess)
	}()

	wg := sync.WaitGroup{}

	for _, url := range urls {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// ensure we can get a 200 response on the partials
			resp, e := http.Get(fmt.Sprintf("http://localhost:3000%s", url))
			assert.NoError(t, e)
			assert.Equal(t, http.StatusOK, resp.StatusCode, fmt.Sprintf("%s was not a 200 response", url))
		}()
	}

	wg.Wait()
}
