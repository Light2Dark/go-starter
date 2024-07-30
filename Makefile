
setup:
	curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-arm64
	chmod +x tailwindcss-macos-arm64
	mv tailwindcss-macos-arm64 static/css/tailwindcss
	./static/css/tailwindcss init

	curl -sL https://unpkg.com/htmx.org/dist/htmx.min.js > ./static/htmx.min.js

dev:
	make -j 3 run_tailwind run_templ run_go port?=8080

run_tailwind:
	./static/css/tailwindcss -i static/css/input.css -o static/css/output.css --watch

run_go:
	wgo run ./cmd/api -port=$(port)

run_templ:
	templ generate --watch --proxy="http://localhost:$(port)"

build:
	go mod tidy && \
   	templ generate && \
	./static/css/tailwindcss -i static/css/input.css -o static/css/output.css --minify && \
	go build -ldflags="-w -s" -o bin/go-starter ./cmd/api

update_packages:
	go get -d -u ./...
	go mod tidy