# CoinMarketCap Go SDK Makefile

.PHONY: help test test-unit test-integration test-coverage build clean lint fmt

# Default target
help: ## Show this help message
	@echo "CoinMarketCap Go SDK"
	@echo "===================="
	@echo ""
	@echo "Available targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

test: test-unit ## Run all tests (unit tests only by default)

test-unit: ## Run unit tests
	@echo "Running unit tests..."
	go test -v

test-integration: ## Run integration tests (requires CMC_API_KEY)
	@echo "Running integration tests..."
	@if [ -z "$(CMC_API_KEY)" ]; then \
		echo "❌ Error: CMC_API_KEY environment variable is required for integration tests"; \
		echo "Set it with: export CMC_API_KEY=your-api-key"; \
		exit 1; \
	fi
	go test -tags=integration -v

test-coverage: ## Run unit tests with coverage report
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

test-all: test-unit test-integration ## Run both unit and integration tests

build: ## Build the package
	@echo "Building package..."
	go build

clean: ## Clean build artifacts and coverage reports
	@echo "Cleaning up..."
	rm -f coverage.out coverage.html

lint: ## Run linter
	@echo "Running linter..."
	go vet ./...
	go fmt ./...

fmt: ## Format code
	@echo "Formatting code..."
	go fmt ./...

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

.PHONY: examples
examples: ## Run example programs
	@echo "Running basic example..."
	@if [ -z "$(CMC_API_KEY)" ]; then \
		echo "❌ Error: CMC_API_KEY environment variable is required for examples"; \
		echo "Set it with: export CMC_API_KEY=your-api-key"; \
		exit 1; \
	fi
	go run examples/basic/main.go

example-advanced: ## Run advanced example
	@echo "Running advanced example..."
	@if [ -z "$(CMC_API_KEY)" ]; then \
		echo "❌ Error: CMC_API_KEY environment variable is required for examples"; \
		echo "Set it with: export CMC_API_KEY=your-api-key"; \
		exit 1; \
	fi
	go run examples/advanced/main.go

# Development targets
dev-setup: ## Set up development environment
	@echo "Setting up development environment..."
	go mod download
	go install golang.org/x/tools/cmd/cover@latest

# CI/CD targets  
ci-test: test-unit ## Run tests for CI (unit tests only)

ci-test-integration: ## Run integration tests for CI
	@echo "Running integration tests in CI..."
	go test -tags=integration -v -timeout=5m