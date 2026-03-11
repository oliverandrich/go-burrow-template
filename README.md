# __ProjectName__

<a href="https://github.com/__GitUser__/__ProjectName__/actions/workflows/ci.yml"><img src="https://img.shields.io/github/actions/workflow/status/__GitUser__/__ProjectName__/ci.yml?branch=main&label=CI&style=for-the-badge" alt="CI"></a>
<a href="https://github.com/__GitUser__/__ProjectName__/releases"><img src="https://img.shields.io/github/v/release/__GitUser__/__ProjectName__?style=for-the-badge" alt="Release"></a>
<a href="https://go.dev/"><img src="https://img.shields.io/github/go-mod/go-version/__GitUser__/__ProjectName__?style=for-the-badge" alt="Go Version"></a>
<a href="https://goreportcard.com/report/github.com/__GitUser__/__ProjectName__"><img src="https://goreportcard.com/badge/github.com/__GitUser__/__ProjectName__?style=for-the-badge" alt="Go Report Card"></a>
<a href="/LICENSE"><img src="https://img.shields.io/github/license/__GitUser__/__ProjectName__?style=for-the-badge" alt="License"></a>

__ProjectDescription__

## Stack

- **[Burrow](https://github.com/oliverandrich/burrow)** - Django-inspired Go web framework
- **Bootstrap 5** with dark/light theme switcher
- **htmx** for server-driven interactivity
- **SQLite** via Burrow's built-in database support
- **just** task runner
- **golangci-lint** for code quality
- **goreleaser** for releases

## Quick Start

```bash
# Create new project from template
gohatch github.com/oliverandrich/go-burrow-template github.com/you/your-app

# Run the development server
cd your-app
just run
```

The server starts at [http://localhost:8080](http://localhost:8080).

## Requirements

- Go 1.24+
- [gohatch](https://github.com/oliverandrich/gohatch)
- [just](https://github.com/casey/just) (command runner)
- [golangci-lint](https://golangci-lint.run/) (linting)
- [tparse](https://github.com/mfridman/tparse) (test output formatting)

## Template Variables

The template uses placeholders that gohatch replaces automatically:

| Placeholder              | Replaced with                        |
| ------------------------ | ------------------------------------ |
| `__ProjectName__`        | Binary name (last path segment)      |
| `__ProjectDescription__` | Project description (from `-d` flag) |
| `__GitUser__`            | GitHub user/org (second path segment)|

## Development

```bash
just setup            # Setup project (download deps, install pre-commit hooks)
just run              # Run the development server
just build            # Build binary to build/<name>
just test             # Run tests
just cover            # Run tests with coverage
just cover-report     # Open coverage report in browser
just fmt              # Format code
just lint             # Run linter
just check            # Run fmt, lint, and test
just clean            # Remove build artifacts
just install          # Install to $GOPATH/bin
just release          # Create release with goreleaser
```

## Project Structure

```
├── cmd/
│   └── <name>/             # Server entry point
│       └── main.go
├── internal/
│   └── pages/              # Pages app (homepage, layout)
│       ├── pages.go
│       └── templates/
│           ├── app/
│           │   └── layout.html
│           └── pages/
│               └── home.html
├── go.mod
├── justfile                # Task runner
├── .golangci.yml           # Linter config
└── .goreleaser.yaml        # Release config
```

## License

EUPL-1.2 - see [LICENSE](LICENSE)
