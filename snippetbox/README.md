# Snippetbox

A web application for pasting and sharing snippets of text — similar to Pastebin or GitHub's Gists.

Built following the steps in the book "Let's Go" with Go 1.25.5.

## Project Structure

```
snippetbox/
├── cmd/web/              # Application entry point
│   ├── main.go          # Server setup, database connection, and configuration
│   ├── handlers.go      # HTTP request handlers
│   ├── helpers.go       # Error handling helpers
│   └── routes.go        # Route definitions
├── internal/
│   └── models/          # Data models and database operations
│       └── snippets.go  # Snippet model and database methods
├── ui/                  # User interface assets
│   ├── html/            # HTML templates
│   └── static/          # Static files (CSS, JavaScript, images)
├── bruno/               # API testing configuration
├── go.mod               # Module definition
└── README.md            # This file
```

## Prerequisites

- Go 1.25.5 or later
- MariaDB or MySQL server running
- Go MySQL driver: `github.com/go-sql-driver/mysql`

## Database Setup

### MariaDB/MySQL Installation

Ensure you have MariaDB or MySQL installed and running. The application expects:

- **Database**: `snippetbox`
- **User**: `web`
- **Password**: `passwd` (default in DSN)

### Creating the Database and Table

Connect to your database and execute:

```sql
CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE snippetbox;

CREATE TABLE snippets (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);

CREATE INDEX idx_snippets_created ON snippets(created);
```

### Creating Database User

```sql
CREATE USER 'web'@'localhost' IDENTIFIED BY 'passwd';
GRANT SELECT, INSERT, UPDATE, DELETE ON snippetbox.* TO 'web'@'localhost';
FLUSH PRIVILEGES;
```

## Getting Started

### Running the Application

1. Navigate to the project directory:
   ```bash
   cd snippetbox
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run the application with default configuration:
   ```bash
   go run ./cmd/web
   ```

4. The server will start on `http://localhost:4000` by default

## Configuration Flags

The application supports the following command-line flags:

**Available Flags:**

- `-addr` — HTTP network address to listen on (default: `:4000`)
- `-dsn` — MySQL data source name (default: `web:passwd@/snippetbox?parseTime=true`)

### View All Flags

To see all available flags:

```bash
go run ./cmd/web -help
```

### Examples

**Running on a different port:**
```bash
go run ./cmd/web -addr ":3000"
```

**Connecting to a remote database:**
```bash
go run ./cmd/web -dsn "user:password@tcp(host:3306)/snippetbox?parseTime=true"
```

**Both flags together:**
```bash
go run ./cmd/web -addr ":8080" -dsn "user:password@/snippetbox?parseTime=true"
```

## Routes

- `GET /` — Home page
- `GET /snippet/view/{id}` — View a specific snippet
- `GET /snippet/create` — Show snippet creation form
- `POST /snippet/create` — Submit a new snippet (currently demonstrates insert functionality)
- `GET /static/{file}` — Static assets (CSS, JS, images)

## Data Models

### Snippet Model

Located in `internal/models/snippets.go`, the `SnippetModel` provides database operations for snippets:

```go
type Snippet struct {
    ID      int           // Unique identifier
    Title   string        // Snippet title
    Content string        // Snippet content
    Created time.Time     // Creation timestamp
    Expires time.Time     // Expiration timestamp
}
```

### Available Methods

- **`Insert(title, content string, expires int) (int, error)`** — Creates a new snippet and returns its ID
  - `title` — Snippet title
  - `content` — Snippet text content
  - `expires` — Number of days until expiration
  - Returns the inserted snippet ID or error

- **`Get(id int) (Snippet, error)`** — Retrieves a specific snippet by ID (stub implementation)

- **`Latest() ([]Snippet, error)`** — Returns the 10 most recently created snippets (stub implementation)

## Error Handling

The application includes structured error handling with comprehensive logging:

### Error Helper Functions

- `app.serverError(w, r, err)` — Logs server errors with context (method, URI) and responds with HTTP 500
- `app.clientError(w, status)` — Responds with appropriate HTTP status codes for client errors

### Logging

The application uses structured logging (`log/slog`) with timestamps:

```
time=2024-01-15T10:30:45.123Z level=INFO msg="starting server" addr=:4000
time=2024-01-15T10:30:46.456Z level=ERROR msg="database error" method=POST uri=/snippet/create
```

## Development

### Adding New Handlers

Add new handler functions to `cmd/web/handlers.go` as methods on the `application` struct:

```go
func (app *application) handlerName(w http.ResponseWriter, r *http.Request) {
    // Handler logic
    // Access database via app.snippets
    // Log with app.logger
}
```

Example of using the database:

```go
id, err := app.snippets.Insert(title, content, 7)
if err != nil {
    app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
    app.serverError(w, r, err)
    return
}
```

### Adding New Routes

Edit `cmd/web/routes.go` to add new routes:

```go
mux.HandleFunc("GET /path/{$}", app.handlerName)
mux.HandleFunc("POST /path/{$}", app.handlerName)
```

### Database Access

All database operations go through `app.snippets` (SnippetModel):

```go
// Insert a new snippet
id, err := app.snippets.Insert("Title", "Content", 7)

// Get a snippet (to be implemented)
snippet, err := app.snippets.Get(id)

// Get latest snippets (to be implemented)
snippets, err := app.snippets.Latest()
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

## Building for Production

Create a compiled binary:

```bash
go build -o snippetbox ./cmd/web
```

Run the binary:

```bash
./snippetbox
./snippetbox -addr ":8080" -dsn "user:password@/snippetbox?parseTime=true"
```

## Dependencies

- `github.com/go-sql-driver/mysql` — MySQL/MariaDB driver for Go

Install/update with:

```bash
go get github.com/go-sql-driver/mysql
go mod tidy
```

## Module

Module: `snippetbox.king`

Manage dependencies:

```bash
go mod tidy
go mod download
```

## Notes

- The server listens on port 4000 by default, configurable via the `-addr` flag
- Database connection details are configurable via the `-dsn` flag
- The DSN must include `?parseTime=true` for proper time.Time parsing from the database
- Static files are served from `ui/static/` directory
- Ensure the `ui/` directory structure is accessible when running the application
- Error logs include contextual information for debugging
- Use `-help` flag to view all available configuration options
- The database connection pool is established on startup; if the connection fails, the application will panic