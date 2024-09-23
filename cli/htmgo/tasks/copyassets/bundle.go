package copyassets

import (
	"fmt"
	"github.com/maddalax/htmgo/cli/htmgo/internal/dirutil"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/module"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/process"
	"golang.org/x/mod/modfile"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func getModuleVersion(modulePath string) (string, error) {
	// Read the go.mod file
	data, err := os.ReadFile(process.GetPathRelativeToCwd("go.mod"))
	if err != nil {
		return "", fmt.Errorf("error reading go.mod: %v", err)
	}

	// Parse the go.mod file
	modFile, err := modfile.Parse(process.GetPathRelativeToCwd("go.mod"), data, nil)
	if err != nil {
		return "", fmt.Errorf("error parsing go.mod: %v", err)
	}

	// Find the module version
	for _, req := range modFile.Require {
		if req.Mod.Path == modulePath {
			return req.Mod.Version, nil
		}
	}

	return "", fmt.Errorf("module %s not found in go.mod", modulePath)
}

func CopyAssets() {
	moduleName := "github.com/maddalax/htmgo/framework"
	modulePath := module.GetDependencyPath(moduleName)

	assetDir := ""
	// Is hosted version and not local version from .work file
	if strings.HasPrefix(modulePath, "github.com/") {
		version, err := getModuleVersion(modulePath)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		dirname, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		assetDir = fmt.Sprintf("%s/go/pkg/mod/%s@%s/assets", dirname, modulePath, version)
	} else {
		assetDir = fmt.Sprintf("%s/assets", modulePath)
	}

	assetDistDir := fmt.Sprintf("%s/dist", assetDir)
	assetCssDir := fmt.Sprintf("%s/css", assetDir)

	cwd := process.GetWorkingDir()

	destDir := fmt.Sprintf("%s/assets", cwd)
	destDirDist := fmt.Sprintf("%s/dist", destDir)
	destDirCss := fmt.Sprintf("%s/css", destDir)

	err := dirutil.CopyDir(assetDistDir, destDirDist, func(path string, exists bool) bool {
		return true
	})

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	err = dirutil.CopyDir(assetCssDir, destDirCss, func(path string, exists bool) bool {
		if strings.HasSuffix(path, "tailwind.config.js") {
			return false
		}
		return true
	})

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if !dirutil.HasFileFromRoot("tailwind.config.js") {
		err = dirutil.CopyFile(
			filepath.Join(assetCssDir, "tailwind.config.js"),
			filepath.Join(process.GetWorkingDir(), "tailwind.config.js"),
		)
	}

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	process.Run(fmt.Sprintf("cd %s && git add .", destDirCss))
}
