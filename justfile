dev:
    just run-gen && just build-css && air & just watch-js & just watch-gen & just watch-css

bundle:
    rm -rf dist
    rm -rf assets/dist
    mkdir -p dist/assets/dist
    just run-gen
    GOOS=linux GOARCH=amd64 go build -o dist/mhtml .
    cd assets/js && npm run build
    just build-css
    cp -r assets/dist/* dist/assets/dist
    tar -czvf mhtml-release.tar.gz ./dist
    rm -rf dist
    mkdir -p dist
    mv mhtml-release.tar.gz dist

release-version := '1.0.0'
release:
    just bundle
    gh release create {{release-version}} dist/mhtml-release.tar.gz --title "Release {{release-version}}" --prerelease --notes "new release"

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