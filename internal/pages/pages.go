// SPDX-License-Identifier: MIT

// Package pages provides the app's static pages (homepage)
// and template overrides for the bootstrap navbar and alerts.
package pages

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oliverandrich/burrow"
	"github.com/oliverandrich/burrow/contrib/bsicons"
	"github.com/oliverandrich/burrow/contrib/messages"
	"github.com/urfave/cli/v3"
)

//go:embed templates
var templateFS embed.FS

// App implements the pages app.
type App struct{}

// New creates a new pages app.
func New() *App { return &App{} }

func (a *App) Name() string { return "pages" }

func (a *App) Configure(cfg *burrow.AppConfig, _ *cli.Command) error {
	cfg.RegisterIconFunc("iconHouse", bsicons.House)
	cfg.RegisterIconFunc("iconPuzzle", bsicons.Puzzle)
	cfg.RegisterIconFunc("iconLightning", bsicons.Lightning)
	return nil
}

// TemplateFS returns the embedded HTML template files.
func (a *App) TemplateFS() fs.FS {
	sub, _ := fs.Sub(templateFS, "templates")
	return sub
}

// FuncMap returns template functions for alert rendering.
func (a *App) FuncMap() template.FuncMap {
	return template.FuncMap{
		"alertClass": func(level messages.Level) string {
			if level == messages.Error {
				return "danger"
			}
			return string(level)
		},
	}
}

// NavItems returns the navigation items for this app.
func (a *App) NavItems() []burrow.NavItem {
	return []burrow.NavItem{
		{Label: "Home", URL: "/", Icon: bsicons.House(), Position: 1},
	}
}

// Routes registers the page routes.
func (a *App) Routes(r chi.Router) {
	r.Get("/", burrow.Handle(home))
}

func home(w http.ResponseWriter, r *http.Request) error {
	return burrow.Render(w, r, http.StatusOK, "pages/home", map[string]any{"Title": "Home"})
}
