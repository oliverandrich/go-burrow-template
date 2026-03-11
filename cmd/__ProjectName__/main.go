// SPDX-License-Identifier: EUPL-1.2

package main

import (
	"context"
	"embed"
	"log"
	"log/slog"
	"os"
	"runtime/debug"

	"github.com/oliverandrich/burrow"
	"github.com/oliverandrich/burrow/contrib/bootstrap"
	"github.com/oliverandrich/burrow/contrib/csrf"
	"github.com/oliverandrich/burrow/contrib/healthcheck"
	"github.com/oliverandrich/burrow/contrib/htmx"
	"github.com/oliverandrich/burrow/contrib/messages"
	"github.com/oliverandrich/burrow/contrib/session"
	"github.com/oliverandrich/burrow/contrib/staticfiles"
	"github.com/oliverandrich/go-burrow-template/internal/pages"
	"github.com/urfave/cli/v3"
)

// version is set via ldflags at build time.
var version = "dev"

// emptyFS is an empty filesystem for staticfiles when the app has
// no user-level static assets. Contrib apps contribute their own via
// HasStaticFiles.
var emptyFS embed.FS

func init() {
	if version == "dev" {
		if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "" && info.Main.Version != "(devel)" {
			version = info.Main.Version
		}
	}
}

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	staticApp, err := staticfiles.New(emptyFS)
	if err != nil {
		log.Fatal(err)
	}

	srv := burrow.NewServer(
		session.New(),
		csrf.New(),
		staticApp,
		healthcheck.New(),
		pages.New(),
		messages.New(),
		htmx.New(),
		bootstrap.New(),
	)

	srv.SetLayout(pages.Layout())

	cmd := &cli.Command{
		Name:     "__ProjectName__",
		Usage:    "__ProjectDescription__",
		Version:  version,
		Flags:    srv.Flags(nil),
		Action:   srv.Run,
		Commands: srv.Registry().AllCLICommands(),
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
