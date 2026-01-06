# Project variables
PROJECT_NAME := snippetbox
BINARY_NAME  := $(PROJECT_NAME)
SOURCES      := ./cmd/web
OUTPUT_DIR   := ./bin

# -- Default Target --
.PHONY: all
all: run

# -- Main Development Tasks --

.PHONY: run
run: ## Run the application (Dev mode)
	@echo "🚀 Running application..."
	go run $(SOURCES)

.PHONY: run-hot
run-hot: ## Run with hot reloading (requires air)
	@echo "🔥 Running with hot-reloading..."
	air

.PHONY: build
build: ## Build the binary for production
	@echo "🔨 Building binary..."
	@mkdir -p $(OUTPUT_DIR)
	go build -o $(OUTPUT_DIR)/$(BINARY_NAME) $(SOURCES)
	@echo "✅ Build complete: $(OUTPUT_DIR)/$(BINARY_NAME)"

# -- Code Quality --

.PHONY: test
test: ## Run tests
	@echo "🧪 Running tests..."
	go test -v ./...

.PHONY: vet
vet: ## Run go vet
	@echo "🧐 Running go vet..."
	go vet ./...

.PHONY: lint
lint: ## Lint the code (requires golangci-lint)
	@echo "🧹 Linting code..."
	golangci-lint run

.PHONY: fmt
fmt: ## Format code
	@echo "📝 Formatting code..."
	go fmt ./...

.PHONY: tidy
tidy: ## Tidy up go.mod dependencies
	@echo "📦 Tidying modules..."
	go mod tidy

# -- Utilities --

.PHONY: clean
clean: ## Clean build artifacts
	@echo "🗑️  Cleaning build artifacts..."
	@rm -rf $(OUTPUT_DIR)
	@echo "✨ Clean complete."

.PHONY: help
help: ## Show this help message
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
