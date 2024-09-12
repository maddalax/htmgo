# Command to run and watch the Go application using Air
run-app:
    just run-gen && air & just watch-js & just watch-gen

run-gen:
    go run ./tooling/astgen

watch-gen:
    go run ./tooling/watch.go --command 'go run ./tooling/astgen'

watch-js:
    cd js && npm run build && npm run watch