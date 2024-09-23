package css

import (
	"fmt"
	"github.com/maddalax/htmgo/cli/htmgo/internal/dirutil"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/copyassets"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/process"
	"log"
	"log/slog"
	"path/filepath"
	"runtime"
)

func IsTailwindEnabled() bool {
	return dirutil.HasFileFromRoot("tailwind.config.js")
}

func Setup() bool {
	if !IsTailwindEnabled() {
		slog.Debug("Tailwind is not enabled. Skipping CSS generation.")
		return false
	}
	downloadTailwindCli()

	if dirutil.HasFileFromRoot("assets/css/input.css") {
		copyassets.CopyAssets()
	}

	return true
}

func GetTailwindExecutableName() string {
	if runtime.GOOS == "windows" {
		return "./__htmgo/tailwind.exe"
	} else {
		return "./__htmgo/tailwind"
	}
}

func GenerateCss(flags ...process.RunFlag) error {
	if !Setup() {
		return nil
	}
	exec := GetTailwindExecutableName()
	return process.RunMany([]string{
		fmt.Sprintf("%s -i ./assets/css/input.css -o ./assets/dist/main.css -c ./tailwind.config.js", exec),
	}, append(flags, process.Silent)...)
}

func GenerateCssWatch(flags ...process.RunFlag) error {
	if !Setup() {
		return nil
	}
	exec := GetTailwindExecutableName()
	return process.RunMany([]string{
		fmt.Sprintf("%s -i ./assets/css/input.css -o ./assets/dist/main.css -c ./tailwind.config.js --watch=always", exec),
	}, append(flags, process.KillOnlyOnExit, process.Silent)...)
}

func downloadTailwindCli() {

	if dirutil.HasFileFromRoot(GetTailwindExecutableName()) {
		slog.Debug("Tailwind CLI already exists. Skipping download.")
		return
	}

	if !IsTailwindEnabled() {
		slog.Debug("Tailwind is not enabled. Skipping tailwind cli download.")
		return
	}
	distro := ""
	os := runtime.GOOS
	arch := runtime.GOARCH
	switch {
	case os == "darwin" && arch == "arm64":
		distro = "macos-arm64"
	case os == "darwin" && arch == "amd64":
		distro = "macos-x64"
	case os == "linux" && arch == "arm64":
		distro = "linux-arm64"
	case os == "linux" && arch == "amd64":
		distro = "linux-x64"
	case os == "windows" && arch == "amd64":
		distro = "windows-x64.exe"
	case os == "windows" && arch == "arm64":
		distro = "windows-arm64.exe"

	default:
		log.Fatal(fmt.Sprintf("Unsupported OS/ARCH: %s/%s", os, arch))
	}
	fileName := fmt.Sprintf(`tailwindcss-%s`, distro)
	url := fmt.Sprintf(`https://github.com/tailwindlabs/tailwindcss/releases/latest/download/%s`, fileName)
	if os == "windows" {
		process.RunMany([]string{
			fmt.Sprintf(`curl -LO %s`, url),
		}, process.ExitOnError)
	} else {
		process.RunMany([]string{
			fmt.Sprintf(`curl -LO %s`, url),
			fmt.Sprintf(`chmod +x %s`, fileName),
		}, process.ExitOnError)
	}

	outputFileName := GetTailwindExecutableName()

	err := dirutil.MoveFile(
		filepath.Join(process.GetWorkingDir(), fileName),
		filepath.Join(process.GetWorkingDir(), outputFileName))

	if err != nil {
		log.Fatalf("Error moving file: %s\n", err.Error())
	}

	slog.Debug("Successfully downloaded Tailwind CLI", slog.String("url", url))
}
