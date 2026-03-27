// setup:feature:demo
package hypermedia

import (
	"context"
	"fmt"
	"strings"
	"sync"
)

// LinkRelation represents a relationship between two resources.
type LinkRelation struct {
	Rel   string // IANA link relation (e.g., "related", "collection", "up")
	Href  string // Target URL
	Title string // Human-readable label
}

// linkRegistry stores registered link relations keyed by source path.
var (
	linksMu  sync.RWMutex
	linksMap = make(map[string][]LinkRelation)
)

// Link registers a relationship from a source path to a target.
// For rel="related", the inverse is automatically registered (symmetric).
func Link(source, rel, target, title string) {
	linksMu.Lock()
	defer linksMu.Unlock()

	linksMap[source] = append(linksMap[source], LinkRelation{
		Rel:   rel,
		Href:  target,
		Title: title,
	})

	// rel="related" is symmetric — auto-create the inverse
	if rel == "related" {
		// Derive the inverse title from the source path
		// e.g., "/demo/inventory" -> "Inventory"
		inverseTitle := TitleFromPath(source)
		linksMap[target] = append(linksMap[target], LinkRelation{
			Rel:   "related",
			Href:  source,
			Title: inverseTitle,
		})
	}
}

// LinksFor returns all registered link relations for a path.
// If rels is provided, only relations matching those types are returned.
func LinksFor(path string, rels ...string) []LinkRelation {
	linksMu.RLock()
	defer linksMu.RUnlock()

	all := linksMap[path]
	if len(rels) == 0 {
		result := make([]LinkRelation, len(all))
		copy(result, all)
		return result
	}

	relSet := make(map[string]bool, len(rels))
	for _, r := range rels {
		relSet[r] = true
	}

	var filtered []LinkRelation
	for _, l := range all {
		if relSet[l.Rel] {
			filtered = append(filtered, l)
		}
	}
	return filtered
}

// RelatedLinksFor returns only rel="related" links for a path,
// excluding the current path itself (for use in context bars).
func RelatedLinksFor(path string) []LinkRelation {
	links := LinksFor(path, "related")
	// Deduplicate by href (symmetric registration can create dupes)
	seen := make(map[string]bool)
	var unique []LinkRelation
	for _, l := range links {
		if l.Href == path || seen[l.Href] {
			continue
		}
		seen[l.Href] = true
		unique = append(unique, l)
	}
	return unique
}

// LinkHeader formats link relations as an RFC 8288 Link header value.
func LinkHeader(links []LinkRelation) string {
	if len(links) == 0 {
		return ""
	}
	parts := make([]string, len(links))
	for i, l := range links {
		parts[i] = fmt.Sprintf("<%s>; rel=\"%s\"; title=\"%s\"", l.Href, l.Rel, l.Title)
	}
	return strings.Join(parts, ", ")
}

// LinkSource provides link relations dynamically for a request.
type LinkSource interface {
	// Links returns link relations relevant to the given path.
	// The context carries request-scoped values such as the session ID.
	Links(ctx context.Context, path string) []LinkRelation
}

var (
	sourcesMu sync.RWMutex
	sources   []LinkSource
)

// RegisterLinkSource adds a dynamic link source to the global set.
func RegisterLinkSource(src LinkSource) {
	sourcesMu.Lock()
	defer sourcesMu.Unlock()
	sources = append(sources, src)
}

// AllSourceLinks collects links from all registered sources for a path.
func AllSourceLinks(ctx context.Context, path string) []LinkRelation {
	sourcesMu.RLock()
	defer sourcesMu.RUnlock()

	var all []LinkRelation
	for _, src := range sources {
		all = append(all, src.Links(ctx, path)...)
	}
	return all
}

// registrySource wraps the static link registry as a LinkSource.
type registrySource struct{}

func (registrySource) Links(_ context.Context, path string) []LinkRelation {
	return LinksFor(path)
}

func init() {
	RegisterLinkSource(registrySource{})
}

// FrecencyFunc returns top frecent pages for a session ID as link relations.
type FrecencyFunc func(ctx context.Context, sessionID string, limit int) ([]LinkRelation, error)

// FrecencySource is a LinkSource backed by session visit history.
type FrecencySource struct {
	Fn    FrecencyFunc
	Limit int
}

// Links returns frecent bookmark links for the current session, excluding
// the page being viewed.
func (f *FrecencySource) Links(ctx context.Context, path string) []LinkRelation {
	sessionID, ok := ctx.Value(sessionIDKey{}).(string)
	if !ok || sessionID == "" {
		return nil
	}
	links, err := f.Fn(ctx, sessionID, f.Limit)
	if err != nil {
		return nil
	}
	var filtered []LinkRelation
	for _, l := range links {
		if l.Href != path {
			filtered = append(filtered, l)
		}
	}
	return filtered
}

// sessionIDKey is the context key for the session ID used by frecency.
type sessionIDKey struct{}

// WithSessionID adds a session ID to the context for frecency tracking.
func WithSessionID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, sessionIDKey{}, id)
}

// TitleFromPath extracts a title from the last segment of a URL path.
// "/demo/inventory" -> "Inventory", "/admin/error-traces" -> "Error Traces"
func TitleFromPath(path string) string {
	path = strings.TrimSuffix(path, "/")
	idx := strings.LastIndex(path, "/")
	if idx < 0 {
		return path
	}
	seg := path[idx+1:]
	seg = strings.ReplaceAll(seg, "-", " ")
	// Title case
	words := strings.Fields(seg)
	for i, w := range words {
		if len(w) > 0 {
			words[i] = strings.ToUpper(w[:1]) + w[1:]
		}
	}
	return strings.Join(words, " ")
}
