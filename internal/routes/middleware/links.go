// setup:feature:demo
package middleware

import (
	"catgoose/dothog/internal/routes/hypermedia"

	"github.com/labstack/echo/v4"
)

// LinkRelationsMiddleware collects link relations from all registered
// LinkSource implementations and stores them on the context for template
// rendering. It also emits an RFC 8288 Link HTTP header.
func LinkRelationsMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			path := c.Request().URL.Path
			ctx := c.Request().Context()
			links := hypermedia.AllSourceLinks(ctx, path)
			if len(links) > 0 {
				// Set RFC 8288 Link header
				c.Response().Header().Set("Link", hypermedia.LinkHeader(links))
				// Store on context for template rendering
				c.Set("link_relations", links)
			}
			return next(c)
		}
	}
}

// GetLinkRelations retrieves link relations from the echo context.
func GetLinkRelations(c echo.Context) []hypermedia.LinkRelation {
	if v := c.Get("link_relations"); v != nil {
		if links, ok := v.([]hypermedia.LinkRelation); ok {
			return links
		}
	}
	return nil
}
