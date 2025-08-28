# Variables
APP_NAME = meli-products-api
MAIN_PATH = cmd/api/main.go
BINARY_NAME = meli-products-api
DOCS_PATH = docs
DATA_PATH = data
PORT = 8080

# Colors for output
GREEN = \033[0;32m
YELLOW = \033[0;33m
BLUE = \033[0;34m
RED = \033[0;31m
NC = \033[0m # No Color

.PHONY: help build run clean test swagger dev deps format lint docker-build docker-run

# Default target
all: deps swagger build

## help: Show this help message
help:
	@echo "$(BLUE)$(APP_NAME) - Available commands:$(NC)"
	@echo ""
	@echo "$(GREEN)Development:$(NC)"
	@echo "  make run          - Run the application"
	@echo "  make dev          - Run in development mode with auto-reload"
	@echo "  make build        - Build the application binary"
	@echo "  make test         - Run tests"
	@echo "  make swagger      - Generate Swagger documentation"
	@echo ""
	@echo "$(GREEN)Dependencies:$(NC)"
	@echo "  make deps         - Install Go dependencies"
	@echo "  make deps-tools   - Install development tools"
	@echo ""
	@echo "$(GREEN)Code Quality:$(NC)"
	@echo "  make format       - Format Go code"
	@echo "  make lint         - Run linter"
	@echo "  make vet          - Run go vet"
	@echo ""
	@echo "$(GREEN)Docker:$(NC)"
	@echo "  make docker-build - Build Docker image"
	@echo "  make docker-run   - Run application in Docker"
	@echo ""
	@echo "$(GREEN)Utilities:$(NC)"
	@echo "  make clean        - Clean build artifacts"
	@echo "  make check-data   - Validate JSON data files"
	@echo "  make help         - Show this help message"

## deps: Install Go dependencies
deps:
	@echo "$(YELLOW)Installing Go dependencies...$(NC)"
	go mod download
	go mod tidy
	@echo "$(GREEN)Dependencies installed successfully!$(NC)"

## deps-tools: Install development tools
deps-tools:
	@echo "$(YELLOW)Installing development tools...$(NC)"
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/cosmtrek/air@latest
	@echo "$(GREEN)Development tools installed successfully!$(NC)"

## swagger: Generate Swagger documentation
swagger:
	@echo "$(YELLOW)Generating Swagger documentation...$(NC)"
	@if ! command -v swag &> /dev/null; then \
		echo "$(RED)Error: swag is not installed. Run 'make deps-tools' first.$(NC)"; \
		exit 1; \
	fi
	swag init -g $(MAIN_PATH) -o $(DOCS_PATH) --parseDependency --parseInternal
	@echo "$(GREEN)Swagger documentation generated successfully!$(NC)"

## build: Build the application
build: swagger
	@echo "$(YELLOW)Building $(APP_NAME)...$(NC)"
	go build -o bin/$(BINARY_NAME) $(MAIN_PATH)
	@echo "$(GREEN)Build completed: bin/$(BINARY_NAME)$(NC)"

## run: Run the application
run: swagger
	@echo "$(YELLOW)Starting $(APP_NAME) on port $(PORT)...$(NC)"
	@echo "$(BLUE)API will be available at: http://localhost:$(PORT)$(NC)"
	@echo "$(BLUE)Swagger docs at: http://localhost:$(PORT)/swagger/index.html$(NC)"
	go run $(MAIN_PATH)

## dev: Run in development mode with auto-reload
dev:
	@echo "$(YELLOW)Starting $(APP_NAME) in development mode...$(NC)"
	@if ! command -v air &> /dev/null; then \
		echo "$(RED)Error: air is not installed. Run 'make deps-tools' first.$(NC)"; \
		exit 1; \
	fi
	air -c .air.toml

## test: Run tests
test:
	@echo "$(YELLOW)Running tests...$(NC)"
	go test -v ./...
	@echo "$(GREEN)Tests completed!$(NC)"

## test-coverage: Run tests with coverage
test-coverage:
	@echo "$(YELLOW)Running tests with coverage...$(NC)"
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "$(GREEN)Coverage report generated: coverage.html$(NC)"

## format: Format Go code
format:
	@echo "$(YELLOW)Formatting Go code...$(NC)"
	go fmt ./...
	@echo "$(GREEN)Code formatted successfully!$(NC)"

## lint: Run linter
lint:
	@echo "$(YELLOW)Running linter...$(NC)"
	@if ! command -v golangci-lint &> /dev/null; then \
		echo "$(RED)Error: golangci-lint is not installed. Run 'make deps-tools' first.$(NC)"; \
		exit 1; \
	fi
	golangci-lint run
	@echo "$(GREEN)Linting completed!$(NC)"

## vet: Run go vet
vet:
	@echo "$(YELLOW)Running go vet...$(NC)"
	go vet ./...
	@echo "$(GREEN)go vet completed!$(NC)"

## check-data: Validate JSON data files
check-data:
	@echo "$(YELLOW)Validating JSON data files...$(NC)"
	@for file in $(DATA_PATH)/*.json; do \
		echo "Checking $$file..."; \
		if ! python -m json.tool $$file > /dev/null 2>&1; then \
			echo "$(RED)Error: $$file contains invalid JSON$(NC)"; \
			exit 1; \
		fi; \
	done
	@echo "$(GREEN)All JSON files are valid!$(NC)"

## clean: Clean build artifacts
clean:
	@echo "$(YELLOW)Cleaning build artifacts...$(NC)"
	rm -rf bin/
	rm -f coverage.out coverage.html
	rm -rf tmp/
	@echo "$(GREEN)Clean completed!$(NC)"

## docker-build: Build Docker image
docker-build:
	@echo "$(YELLOW)Building Docker image...$(NC)"
	docker build -t $(APP_NAME):latest .
	@echo "$(GREEN)Docker image built successfully!$(NC)"

## docker-run: Run application in Docker
docker-run:
	@echo "$(YELLOW)Running $(APP_NAME) in Docker...$(NC)"
	docker run -p $(PORT):$(PORT) --name $(APP_NAME) --rm $(APP_NAME):latest

## install: Install the application binary
install: build
	@echo "$(YELLOW)Installing $(APP_NAME)...$(NC)"
	go install $(MAIN_PATH)
	@echo "$(GREEN)$(APP_NAME) installed successfully!$(NC)"

## mod-tidy: Clean up go.mod and go.sum
mod-tidy:
	@echo "$(YELLOW)Tidying go modules...$(NC)"
	go mod tidy
	@echo "$(GREEN)Modules tidied!$(NC)"

## security-check: Run security checks
security-check:
	@echo "$(YELLOW)Running security checks...$(NC)"
	@if command -v gosec &> /dev/null; then \
		gosec ./...; \
	else \
		echo "$(RED)gosec not installed. Install with: go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest$(NC)"; \
	fi

## benchmark: Run benchmarks
benchmark:
	@echo "$(YELLOW)Running benchmarks...$(NC)"
	go test -bench=. -benchmem ./...

# Create air configuration for development
.air.toml:
	@echo "$(YELLOW)Creating air configuration...$(NC)"
	@cat > .air.toml << 'EOF'
root = "."
testdata_dir = "testdata"
tmp_dir = "tmp"

[build]
args_bin = []
bin = "./tmp/main"
cmd = "go build -o ./tmp/main $(MAIN_PATH)"
delay = 1000
exclude_dir = ["assets", "tmp", "vendor", "testdata", "docs"]
exclude_file = []
exclude_regex = ["_test.go"]
exclude_unchanged = false
follow_symlink = false
full_bin = ""
include_dir = []
include_ext = ["go", "tpl", "tmpl", "html"]
kill_delay = "0s"
log = "build-errors.log"
send_interrupt = false
stop_on_root = false

[color]
app = ""
build = "yellow"
main = "magenta"
runner = "green"
watcher = "cyan"

[log]
time = false

[misc]
clean_on_exit = false

[screen]
clear_on_rebuild = false
EOF
	@echo "$(GREEN)Air configuration created!$(NC)"

# Setup development environment
setup-dev: deps deps-tools .air.toml swagger
	@echo "$(GREEN)Development environment setup completed!$(NC)"
	@echo "$(BLUE)You can now run 'make dev' to start development mode$(NC)"

# Quick development workflow
quick: format vet test run

# Production build
prod-build: clean deps swagger test build
	@echo "$(GREEN)Production build completed!$(NC)"
