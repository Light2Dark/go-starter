# GoLang Template

HTMX + Go + TailwindCSS + Templ
- Comes with hot-reload
- Uses slog for logging, tint for colorized logging

## Installation

1. Install wgo (https://github.com/bokwoon95/wgo) with `go install github.com/bokwoon95/wgo@latest`.

2. Run make commands to setup and start the server.
```
make setup          # First time setup

make dev            # Start a server locally
make dev port=9000  # Start a server at a specific port

make build          # Build the binary
```
