// setup:feature:demo

package routes

import (
	"context"
	"net/http"

	"catgoose/dothog/internal/demo"
	"catgoose/dothog/internal/routes/handler"
	"catgoose/dothog/internal/routes/hypermedia"
	"catgoose/dothog/internal/routes/middleware"
	"catgoose/dothog/web/views"

	"github.com/labstack/echo/v4"
)

// initFrecency wires the frecency link source and visit recording middleware.
func (ar *appRoutes) initFrecency(db *demo.DB) {
	// Build a set of tracked GET routes for visit recording.
	validRoutes := make(map[string]bool)
	for _, r := range ar.e.Routes() {
		if r.Method == http.MethodGet && r.Path != "" && r.Path != "/*" {
			validRoutes[r.Path] = true
		}
	}

	// Register the frecency middleware (sets session ID on context, records visits).
	ar.e.Use(middleware.FrecencyMiddleware(db, validRoutes))

	// Bridge demo.DB.TopFrecent → hypermedia.LinkRelation slice.
	frecencyFn := func(ctx context.Context, sessionID string, limit int) ([]hypermedia.LinkRelation, error) {
		visits, err := db.TopFrecent(ctx, sessionID, limit)
		if err != nil {
			return nil, err
		}
		links := make([]hypermedia.LinkRelation, len(visits))
		for i, v := range visits {
			links[i] = hypermedia.LinkRelation{
				Rel:   "bookmark",
				Href:  v.Path,
				Title: v.Title,
			}
		}
		return links, nil
	}

	hypermedia.RegisterLinkSource(&hypermedia.FrecencySource{
		Fn:    frecencyFn,
		Limit: 5,
	})
}

// handleDemoIndex renders the /demo page with popular pages from the DB.
func (ar *appRoutes) handleDemoIndex(c echo.Context) error {
	var popular []demo.PageVisit
	if ar.demoDB != nil {
		if pp, err := ar.demoDB.PopularPages(c.Request().Context(), 5); err == nil {
			popular = pp
		}
	}
	return handler.RenderBaseLayout(c, views.DemoIndexPage(popular...))
}
