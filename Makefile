# Variables
APP_NAME := bin/main
MAIN_FILE := main.go
PKG := ./...
TEST_OUTPUT := test-report.out
LINT_OUTPUT := lint-report.out

# Targets
.PHONY: all build run test clean lint fmt vet

# Default target
all: build

# Build the application
build:
	@echo "Building the application..."
	go build -o $(APP_NAME) $(MAIN_FILE)

# Run the application
run: build
	@echo "Running the application..."
	./$(APP_NAME)

# Test the application
test:
	@echo "Running tests..."
	go test -v $(PKG) | tee $(TEST_OUTPUT)

# Lint the code (requires golangci-lint)
lint:
	@echo "Running linter..."
	golangci-lint run $(PKG) | tee $(LINT_OUTPUT)

# Format the code
fmt:
	@echo "Formatting the code..."
	go fmt $(PKG)

# Check for issues in the code
vet:
	@echo "Running go vet..."
	go vet $(PKG)

# Clean up built files
clean:
	@echo "Cleaning up..."
	rm -rf $(APP_NAME) $(TEST_OUTPUT) $(LINT_OUTPUT)

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod tidy