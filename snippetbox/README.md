# Snippetbox

A web application for pasting and sharing snippets of text — similar to Pastebin or GitHub's Gists.

Built following the steps in the book "Let's Go" with Go 1.25.5.

## Project Structure

```
snippetbox/
├── cmd/web/              # Application entry point
│   ├── main.go          # Server setup, flags, and configuration
│   ├── handlers.go      # HTTP request handlers
│   ├── helpers.go       # Error handling helpers
│   └── routes.go        # Route definitions
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
2. Run the application with default configuration:
   ```bash
   go run ./cmd/web
   ```

3. The server will start on `http://localhost:4000` by default

### Configuration Flags

The application supports the following command-line flags:

```bash
go run ./cmd/web -addr ":8080"
```

**Available Flags:**

- `-addr` — HTTP network address to listen on (default: `:4000`)

To view all available flags:

```bash
go run ./cmd/web -help
```

**Example: Running on a different port**

```bash
go run ./cmd/web -addr ":3000"
```

**Example: Running on a specific IP and port**

```bash
go run ./cmd/web -addr "192.168.1.100:8080"
```

## Routes

- `GET /` — Home page
- `GET /snippet/view/{id}` — View a specific snippet
- `GET /snippet/create` — Show snippet creation form
- `POST /snippet/create` — Submit a new snippet
- `GET /static/{file}` — Static assets (CSS, JS, images)

## Error Handling

The application includes structured error handling with comprehensive logging:

### Error Helper Functions

The application provides two helper methods for error handling:

- `app.serverError(w, r, err)` — Logs server errors with context (method, URI) and responds with a 500 status
- `app.clientError(w, status)` — Responds with appropriate HTTP status codes for client errors

### Logging

The application uses structured logging (`log/slog`) to output request context and error details:

```
time=2024-01-15T10:30:45.123Z level=INFO msg="starting server" addr=:4000
time=2024-01-15T10:30:46.456Z level=ERROR msg="template parse error" method=GET uri=/
```

Error logs include:
- Error message
- HTTP method
- Request URI
- Timestamp

## Development

### Adding New Routes

Routes are defined in `cmd/web/routes.go`. Edit the `routes()` method to add new routes:

```go
mux.HandleFunc("GET /path/{$}", app.handlerName)
```

### Adding New Handlers

Add new handler functions to `cmd/web/handlers.go` as methods on the `application` struct:

```go
func (app *application) handlerName(w http.ResponseWriter, r *http.Request) {
    // Handler logic
}
```

For errors, use the helper functions:

```go
if err != nil {
    app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
    app.serverError(w, r, err)
    return
}
```

### Templates and Static Files

- HTML templates go in `ui/html/`
- Static assets (CSS, JS, images) go in `ui/static/`
- Templates use the base template pattern with partials

### Logging in Handlers

Use structured logging for debugging and monitoring:

```go
app.logger.Info("snippet created", slog.Int("id", snippetID))
app.logger.Error("database error", slog.String("operation", "insert"))
```

## Building

To create a compiled binary:

```bash
go build -o snippetbox ./cmd/web
```

Then run with:

```bash
./snippetbox
./snippetbox -addr ":8080"
```

## Module

Module: `snippetbox.king`

Manage dependencies with:

```bash
go mod tidy
go mod download
```

## Notes

- The server listens on port 4000 by default, configurable via the `-addr` flag
- Static files are served from `ui/static/` directory
- Ensure the `ui/` directory structure is accessible when running the application
- Error logs include contextual information for debugging
- Use `-help` flag to view all available configuration options