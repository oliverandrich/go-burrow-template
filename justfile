# __ProjectName__ - Task Runner

# Version from git
version := `git describe --tags --always --dirty 2>/dev/null || echo "dev"`

# Default: show available commands
default:
    @just --list

# Setup project after checkout
setup:
    go mod download
    @command -v pre-commit >/dev/null && pre-commit install || echo "pre-commit not found, skipping hook installation"

# Build binary
build:
    go build -ldflags="-s -w -X 'main.version={{version}}'" -trimpath -o build/__ProjectName__ ./cmd/__ProjectName__

# Run the development server
run:
    go run -ldflags="-X 'main.version={{version}}'" ./cmd/__ProjectName__

# Run all tests
# Requires: go install github.com/mfridman/tparse@latest
test:
    set -o pipefail && go test -json ./... | tparse -progress

# Run tests with coverage report
# Requires: go install github.com/mfridman/tparse@latest
cover:
    set -o pipefail && go test -json -coverprofile=coverage.out ./... | tparse -progress

# Open coverage report in browser
cover-report:
    go tool cover -html=coverage.out

# Format code
fmt:
    go fmt ./...

# Lint code
lint:
    golangci-lint run

# Run all checks (format, lint, test)
check:
    just fmt
    just lint
    just test

# Clean build artifacts
clean:
    rm -rf build/

# Install binary to $GOPATH/bin
install:
    go install -ldflags="-s -w -X 'main.version={{version}}'" ./cmd/__ProjectName__

# Tidy dependencies
tidy:
    go mod tidy

# Create and publish a release (requires git tag)
release *ARGS:
    goreleaser release --clean {{ARGS}}
