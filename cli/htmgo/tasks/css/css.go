package css

import (
	"fmt"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/process"
	"log"
	"runtime"
)

func GenerateCss(flags ...process.RunFlag) error {
	return process.RunMany([]string{
		"./assets/css/tailwindcss -i ./assets/css/input.css -o ./assets/dist/main.css -c ./assets/css/tailwind.config.js",
	}, append(flags, process.Silent)...)
}

func GenerateCssWatch(flags ...process.RunFlag) error {
	return process.RunMany([]string{
		"./assets/css/tailwindcss -i ./assets/css/input.css -o ./assets/dist/main.css -c ./assets/css/tailwind.config.js --watch=always",
	}, append(flags, process.KillOnlyOnExit, process.Silent)...)
}

func DownloadTailwindCli() error {
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
	return process.RunMany([]string{
		fmt.Sprintf(`curl -LO %s`, url),
		fmt.Sprintf(`chmod +x %s`, fileName),
		fmt.Sprintf(`mv %s ./assets/css/tailwindcss`, fileName),
	}, process.ExitOnError)
}
