// SPDX-License-Identifier: EUPL-1.2

// Package pages provides the app's static pages (homepage),
// layout rendering, and request-path middleware for active nav link highlighting.
package pages

import (
	"context"
	"embed"
	"html/template"
	"io/fs"
	"maps"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/oliverandrich/burrow"
	"github.com/oliverandrich/burrow/contrib/bsicons"
	"github.com/oliverandrich/burrow/contrib/messages"
)

//go:embed templates
var templateFS embed.FS

// ctxKeyRequestPath is used to pass the request path into the template context.
type ctxKeyRequestPath struct{}

// App implements the pages app.
type App struct{}

// New creates a new pages app.
func New() *App { return &App{} }

func (a *App) Name() string                       { return "pages" }
func (a *App) Register(_ *burrow.AppConfig) error { return nil }

// TemplateFS returns the embedded HTML template files.
func (a *App) TemplateFS() fs.FS {
	sub, _ := fs.Sub(templateFS, "templates")
	return sub
}

// FuncMap returns template functions for the layout and home page.
func (a *App) FuncMap() template.FuncMap {
	return template.FuncMap{
		"iconHouse":     func(class ...string) template.HTML { return bsicons.House(class...) },
		"iconPuzzle":    func(class ...string) template.HTML { return bsicons.Puzzle(class...) },
		"iconLightning": func(class ...string) template.HTML { return bsicons.Lightning(class...) },
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

// Middleware injects the request path into context for nav link highlighting.
func (a *App) Middleware() []func(http.Handler) http.Handler {
	return []func(http.Handler) http.Handler{
		func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				ctx := context.WithValue(r.Context(), ctxKeyRequestPath{}, r.URL.Path)
				next.ServeHTTP(w, r.WithContext(ctx))
			})
		},
	}
}

// Routes registers the page routes.
func (a *App) Routes(r chi.Router) {
	r.Get("/", burrow.Handle(home))
}

// navItemData holds a pre-computed nav item for template rendering.
type navItemData struct {
	Label     string
	URL       string
	Icon      template.HTML
	LinkClass string
}

// Layout returns a LayoutFunc that wraps page content in the app layout.
func Layout() burrow.LayoutFunc {
	return func(w http.ResponseWriter, r *http.Request, code int, content template.HTML, data map[string]any) error {
		exec := burrow.TemplateExecutorFromContext(r.Context())
		if exec == nil {
			return burrow.HTML(w, code, string(content))
		}

		ctx := r.Context()
		currentPath, _ := ctx.Value(ctxKeyRequestPath{}).(string)

		// Pre-compute nav items with CSS classes.
		allItems := burrow.NavItems(ctx)
		navItems := make([]navItemData, len(allItems))
		for i, item := range allItems {
			navItems[i] = navItemData{
				Label:     item.Label,
				URL:       item.URL,
				Icon:      item.Icon,
				LinkClass: navLinkClass(currentPath, item.URL),
			}
		}

		layoutData := make(map[string]any, len(data)+3)
		maps.Copy(layoutData, data)
		layoutData["Content"] = content
		layoutData["NavItems"] = navItems
		layoutData["Messages"] = messages.Get(ctx)
		if _, ok := layoutData["Title"]; !ok {
			layoutData["Title"] = ""
		}

		html, err := exec(r, "app/layout", layoutData)
		if err != nil {
			return err
		}
		return burrow.HTML(w, code, string(html))
	}
}

// navLinkClass returns CSS classes for a nav link, marking it active
// when it matches the current path.
func navLinkClass(currentPath, itemURL string) string {
	if currentPath == "" {
		return "nav-link"
	}
	if itemURL == "/" {
		if currentPath == "/" {
			return "nav-link active"
		}
		return "nav-link"
	}
	if strings.HasPrefix(currentPath, itemURL) {
		return "nav-link active"
	}
	return "nav-link"
}

func home(w http.ResponseWriter, r *http.Request) error {
	return burrow.RenderTemplate(w, r, http.StatusOK, "pages/home", map[string]any{"Title": "Home"})
}
