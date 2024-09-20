package css

import "github.com/maddalax/htmgo/cli/htmgo/tasks/process"

func GenerateCss(flags ...process.RunFlag) error {
	return process.RunMany([]string{
		"chmod +x ./assets/css/tailwindcss",
		"./assets/css/tailwindcss -i ./assets/css/input.css -o ./assets/dist/main.css -c ./assets/css/tailwind.config.js",
	}, append(flags, process.Silent)...)
}

func GenerateCssWatch(flags ...process.RunFlag) error {
	return process.RunMany([]string{
		"./assets/css/tailwindcss -i ./assets/css/input.css -o ./assets/dist/main.css -c ./assets/css/tailwind.config.js --watch=always",
	}, append(flags, process.KillOnlyOnExit, process.Silent)...)
}
