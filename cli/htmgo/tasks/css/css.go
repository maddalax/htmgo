package css

import (
	"fmt"
	"github.com/maddalax/htmgo/cli/htmgo/internal/dirutil"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/copyassets"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/process"
	"log"
	"log/slog"
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

func GenerateCss(flags ...process.RunFlag) error {
	if !Setup() {
		return nil
	}
	return process.RunMany([]string{
		"./__htmgo/tailwind -i ./assets/css/input.css -o ./assets/dist/main.css -c ./tailwind.config.js",
	}, append(flags, process.Silent)...)
}

func GenerateCssWatch(flags ...process.RunFlag) error {
	if !Setup() {
		return nil
	}
	return process.RunMany([]string{
		"./__htmgo/tailwind -i ./assets/css/input.css -o ./assets/dist/main.css -c ./tailwind.config.js --watch=always",
	}, append(flags, process.KillOnlyOnExit, process.Silent)...)
}

func downloadTailwindCli() {

	if dirutil.HasFileFromRoot("__htmgo/tailwind") {
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
	default:
		log.Fatal(fmt.Sprintf("Unsupported OS/ARCH: %s/%s", os, arch))
	}
	fileName := fmt.Sprintf(`tailwindcss-%s`, distro)
	url := fmt.Sprintf(`https://github.com/tailwindlabs/tailwindcss/releases/latest/download/%s`, fileName)
	process.RunMany([]string{
		fmt.Sprintf(`curl -LO %s`, url),
		fmt.Sprintf(`chmod +x %s`, fileName),
		fmt.Sprintf(`mv %s ./__htmgo/tailwind`, fileName),
	}, process.ExitOnError)

	slog.Debug("Successfully downloaded Tailwind CLI", slog.String("url", url))
}
