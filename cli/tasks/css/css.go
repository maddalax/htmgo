package css

import "github.com/maddalax/htmgo/cli/tasks/process"

func GenerateCss(exitOnError bool) error {
	return process.RunMany([]string{
		"chmod +x ./assets/css/tailwindcss",
		"./assets/css/tailwindcss -i ./assets/css/input.css -o ./assets/dist/main.css -c ./assets/css/tailwind.config.js",
	}, exitOnError)
}
