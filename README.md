# __ProjectName__

<a href="https://github.com/__GitUser__/__ProjectName__/actions/workflows/ci.yml"><img src="https://img.shields.io/github/actions/workflow/status/__GitUser__/__ProjectName__/ci.yml?branch=main&label=CI&style=for-the-badge" alt="CI"></a>
<a href="https://github.com/__GitUser__/__ProjectName__/releases"><img src="https://img.shields.io/github/v/release/__GitUser__/__ProjectName__?style=for-the-badge" alt="Release"></a>
<a href="https://go.dev/"><img src="https://img.shields.io/github/go-mod/go-version/__GitUser__/__ProjectName__?style=for-the-badge" alt="Go Version"></a>
<a href="https://goreportcard.com/report/github.com/__GitUser__/__ProjectName__"><img src="https://goreportcard.com/badge/github.com/__GitUser__/__ProjectName__?style=for-the-badge" alt="Go Report Card"></a>
<a href="/LICENSE"><img src="https://img.shields.io/github/license/__GitUser__/__ProjectName__?style=for-the-badge" alt="License"></a>

__ProjectDescription__

## Stack

- **[Burrow](https://github.com/oliverandrich/burrow)** v0.20+ — Django-inspired Go web framework
- **Tailwind v4** via the standalone CLI (`cmd/burrow-tailwind`), with dark mode that follows `prefers-color-scheme`
- **htmx** for server-driven interactivity
- **SQLite** via Den (pure Go, no CGO)
- **mise** for task running and tool pinning
- **golangci-lint** for code quality
- **goreleaser** for releases

## Quick Start

```bash
# Create a new project from the template
gohatch github.com/oliverandrich/go-burrow-template github.com/you/your-app
cd your-app

# Install pinned tools (Go, Tailwind, linters, ...) via mise
mise install

# Run the development server with live reload
mise run dev
```

The server starts at [http://localhost:8080](http://localhost:8080).

## Requirements

- [mise](https://mise.jdx.dev/) (handles every other dependency on `mise install`)
- [gohatch](https://github.com/oliverandrich/gohatch) for instantiating the template

`mise install` pins and installs the Go toolchain, the Tailwind v4 CLI, `air` for live reload, `golangci-lint`, `tparse`, `goimports`, `govulncheck`, and `pre-commit`. No manual `go install` lines.

## Template Variables

The template uses placeholders that gohatch replaces automatically:

| Placeholder              | Replaced with                        |
| ------------------------ | ------------------------------------ |
| `__ProjectName__`        | Binary name (last path segment)      |
| `__ProjectDescription__` | Project description (from `-d` flag) |
| `__GitUser__`            | GitHub user/org (second path segment)|

## Development

```bash
mise tasks            # List every available task
mise run setup        # Verify dev tools + remind to install pre-commit hooks
mise run dev          # Live-reload dev server (builds CSS first, then air)
mise run run-once     # Same but without live reload
mise run css          # Build the Tailwind CSS bundle once
mise run css-watch    # Rebuild the CSS on every template change
mise run test         # Run tests
mise run coverage     # Run tests with HTML coverage report
mise run fmt          # Format code
mise run lint         # Run golangci-lint
mise run vuln         # Run govulncheck
mise run tidy         # Tidy module dependencies
mise run install      # Install the binary to $GOPATH/bin
```

`mise run dev` builds the Tailwind CSS bundle once, then starts [air](https://github.com/air-verse/air) which rebuilds and restarts the Go binary whenever `.go` or `.html` files change. For CSS changes during development, run `mise run css-watch` in a second terminal — air doesn't touch the stylesheet.

On first run a `.dev-keys` file is generated with persistent `SESSION_HASH_KEY` and `CSRF_KEY` so sessions and CSRF tokens survive reloads. The file is gitignored.

## Project Structure

```
├── cmd/
│   └── <name>/                  # Server entry point
│       └── main.go              # NewServer + cli wiring
├── internal/
│   └── app/                     # Shell app: layout, homepage, CSS bundle (Pattern B)
│       ├── app.go               # HasTemplates, HasStaticFiles, NavItems, Routes
│       ├── static/              # Tailwind output (app.min.css, gitignored)
│       └── templates/
│           ├── app/
│           │   ├── layout.html  # Site layout with navbar, alerts, htmx
│           │   └── icons.html   # Inline-SVG icon defines
│           ├── error/           # Tailwind-styled overrides for 403/404/500/...
│           │   └── errors.html
│           └── pages/
│               └── home.html
├── tailwind.css                 # Tailwind entrypoint (imports source list)
├── go.mod                       # Pinned to burrow v0.20+
├── .mise.toml                   # Tool pins + task runner
├── .golangci.yml                # Linter config
└── .goreleaser.yaml             # Release config
```

`internal/app/` follows the Pattern B project layout from Burrow's [Tailwind guide](https://burrow.readthedocs.io/en/latest/guide/tailwind/) — the shell app owns the layout templates, the compiled Tailwind CSS, and the project-level icon defines, all served under the `/static/app/` URL prefix.

## License

MIT - see [LICENSE](LICENSE)
