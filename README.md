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

- Go 1.26+
- [gohatch](https://github.com/oliverandrich/gohatch)
- [just](https://github.com/casey/just) (command runner)
- [air](https://github.com/air-verse/air) (live reload during development)
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
just setup            # Check that all required dev tools are installed
just run              # Run the development server with live reload (air)
just run-once         # Run the development server without live reload
just test             # Run tests
just coverage         # Run tests with coverage report
just fmt              # Format code
just lint             # Run linter
just tidy             # Tidy module dependencies
just install          # Install to $GOPATH/bin
```

`just run` starts [air](https://github.com/air-verse/air) which rebuilds and
restarts the server whenever `.go` or `.html` files change. On first run a
`.dev-keys` file is generated with persistent `SESSION_HASH_KEY` and
`CSRF_KEY` values so sessions and CSRF tokens survive air reloads. The file
is gitignored.

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

MIT - see [LICENSE](LICENSE)
