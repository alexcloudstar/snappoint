.PHONY: init build cross-compile test test-coverage test-integration lint install clean

BINARY_NAME=snappoint
BUILD_DIR=build
VERSION?=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS=-ldflags "-X github.com/alexcloudstar/snappoint-cli/pkg/version.Version=$(VERSION)"

init:
	@echo "Initializing Go module..."
	go mod tidy
	go mod download

build:
	@echo "Building $(BINARY_NAME) for current platform..."
	@mkdir -p $(BUILD_DIR)
	go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/snappoint

cross-compile:
	@echo "Cross-compiling for all platforms..."
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 ./cmd/snappoint
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 ./cmd/snappoint
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 ./cmd/snappoint
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 ./cmd/snappoint
	@echo "Cross-compilation complete. Binaries in $(BUILD_DIR)/"

test:
	@echo "Running unit tests..."
	go test -v -race ./...

test-coverage:
	@echo "Running tests with coverage..."
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out | tail -n 1

test-integration:
	@echo "Running integration tests..."
	go test -v -tags=integration ./test/integration/...

lint:
	@echo "Running linter..."
	@which golangci-lint > /dev/null || (echo "golangci-lint not installed. Install from https://golangci-lint.run/usage/install/" && exit 1)
	golangci-lint run ./...

install: build
	@echo "Installing $(BINARY_NAME) to /usr/local/bin..."
	@sudo cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)
	@sudo chmod +x /usr/local/bin/$(BINARY_NAME)
	@echo "Installation complete. Run '$(BINARY_NAME) --version' to verify."

clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out
	@rm -f *.test
	@echo "Clean complete."
