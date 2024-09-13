# Command to run and watch the Go application using Air
run-app:
    just run-gen && air & just watch-js & just watch-gen

run-gen:
    go run ./tooling/astgen

watch-gen:
    go run ./tooling/watch.go --command 'go run ./tooling/astgen'

watch-js:
    cd assets/js && npm run build && npm run watch

setup-tailwind-cli:
    cd assets/css && curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-arm64
    cd assets/css && chmod +x tailwindcss-macos-arm64
    cd assets/css && mv tailwindcss-macos-arm64 tailwindcss

watch-css:
    cd assets/css && ./tailwindcss -i input.css -o ./../dist/main.css --watch

build-css:
    cd assets/css && ./tailwindcss -i input.css -o ./../dist/main.css --minify