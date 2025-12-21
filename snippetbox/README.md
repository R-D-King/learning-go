# Snippetbox

A web application for pasting and sharing snippets of text — similar to Pastebin or GitHub's Gists.

Built following the steps in the book "Let's Go" with Go 1.25.5.

## Project Structure

```
snippetbox/
├── cmd/web/              # Application entry point
│   ├── main.go          # Server setup and routing
│   └── handlers.go      # HTTP request handlers
├── internal/            # Internal packages (reserved for future use)
├── ui/                  # User interface assets
│   ├── html/            # HTML templates
│   └── static/          # Static files (CSS, JavaScript, images)
├── bruno/               # API testing configuration
├── go.mod               # Module definition
└── README.md            # This file
```

## Getting Started

### Prerequisites

- Go 1.25.5 or later

### Running the Application

1. Clone or navigate to the project directory
2. Run the application:
   ```bash
   go run ./cmd/web
   ```

3. The server will start on `http://localhost:4000`

## Routes

- `GET /` — Home page
- `GET /snippet/view/{id}` — View a specific snippet
- `GET /snippet/create` — Show snippet creation form
- `POST /snippet/create` — Submit a new snippet
- `GET /static/{file}` — Static assets (CSS, JS, images)

## Development

### Adding New Routes

Edit `cmd/web/main.go` to add new routes to the `mux`:
```go
mux.HandleFunc("GET /path/{$}", handlerFunction)
```

### Adding New Handlers

Add new handler functions to `cmd/web/handlers.go` following this pattern:
```go
func handlerName(w http.ResponseWriter, r *http.Request) {
    // Handler logic
}
```

### Templates and Static Files

- HTML templates go in `ui/html/`
- Static assets (CSS, JS, images) go in `ui/static/`

## Building

To create a compiled binary:

```bash
go build -o snippetbox ./cmd/web
```

Then run with:
```bash
./snippetbox
```

## Module

Module: `snippetbox.king`

Manage dependencies with:
```bash
go mod tidy
go mod download
```

## Notes

- The server listens on port 4000
- Static files are served from `ui/static/` directory
- Ensure the `ui/` directory is accessible when running the application