# PHILOSOPHY.md Compliance Audit -- Dothog Demo App

**Date:** 2026-03-29
**Scope:** Demo application (`internal/demo/`, `internal/routes/`, `web/views/`, `web/components/`)
**Audited against:** `PHILOSOPHY.md` (22 principles)

## Executive Summary

The dothog demo application demonstrates **strong overall compliance** with its PHILOSOPHY.md. The architecture is genuinely hypermedia-driven, the `Control` struct uniform interface is well-implemented, error handling is exemplary, SQL is explicit, and traits/composable helpers replace ORM abstractions effectively. **9 violations** were found, mostly around strict POST/Redirect/GET enforcement and a handful of verb-based URLs.

### Compliance Scorecard

| Status | Count |
|--------|-------|
| Compliant | 15 principles |
| Minor violations | 5 principles |
| Medium violations | 2 principles |
| Not applicable | 0 |

---

## Violations Found

### V1: `panic` in `must()` helper (Low)

- **Principle:** #20 -- No panic for runtime conditions
- **File:** `main.go:47-52`
- **Description:** The `must()` helper panics if the embedded filesystem fails to initialize. While this occurs at init-time (before `main()` runs), the philosophy draws a hard line on panic.
- **Fix:** Replace with `log.Fatal("failed to sub static FS: ", err)` or document as an acceptable init-time invariant.

### V2: POST handlers return components instead of redirecting (Medium)

- **Principle:** #8 -- Mutations redirect (POST/Redirect/GET)
- **Files:** `internal/routes/routes_inventory.go:68-75` (`handleCreateItem`), `internal/routes/routes_repository.go` (`handleCreateTask`), `internal/routes/routes_hypermedia.go` (`handleCRUDCreate`)
- **Description:** POST handlers respond directly with rendered components (200 OK) instead of issuing `HX-Redirect` with 303. This is a common HTMX inline table update pattern but contradicts the documented PRG principle.
- **Fix:** Either align code (return HX-Redirect + 303) or refine the philosophy to clarify that inline HTMX partial updates are an acceptable exception to PRG for non-navigation mutations.

### V3: PUT handlers return components instead of redirecting (Medium)

- **Principle:** #8 -- Mutations redirect
- **Files:** `internal/routes/routes_people.go:104` (`handlePersonUpdate`), `internal/routes/routes_inventory.go` (`handleUpdateItem`), `internal/routes/routes_vendors_contacts.go` (`handleContactUpdate`)
- **Description:** PUT handlers return rendered components directly (200) instead of redirecting with 303/HX-Redirect.
- **Fix:** Same as V2.

### V4: Verb-based URL `/demo/kanban/tasks/:id/move` (Low)

- **Principle:** #4 -- Resource identification (no verb-based URLs)
- **File:** `internal/routes/routes_kanban.go:31`
- **Description:** URL contains the verb "move."
- **Fix:** Use `PATCH /demo/kanban/tasks/:id` with target status in request body.

### V5: Verb-based URL `/demo/approvals/:id/:action` (Low)

- **Principle:** #4 -- Resource identification
- **File:** `internal/routes/routes_approvals.go:31`
- **Description:** URL captures an `:action` parameter (approve/reject), embedding the verb.
- **Fix:** Use `PATCH /demo/approvals/:id` with action in body.

### V6: Verb-based URLs for restore/archive/unarchive (Low)

- **Principle:** #4 -- Resource identification
- **File:** `internal/routes/routes_repository.go:37-39`
- **URLs:** `POST .../restore`, `POST .../archive`, `POST .../unarchive`
- **Fix:** Use `PATCH /demo/repository/tasks/:id` with desired state in body.

### V7: POST used for state modifications (not creation) (Medium)

- **Principle:** #3 -- Self-descriptive HTTP methods (POST creates)
- **File:** `internal/routes/routes_repository.go:37-39`
- **Description:** `POST` for restore/archive/unarchive -- these are partial modifications of existing resources, not creations.
- **Fix:** Use `PATCH` instead of `POST`.

### V8: Alpine.js used where `_hyperscript` preferred (Low)

- **Principle:** #13 -- Locality of behavior
- **File:** `web/components/core/controls.templ` (backButton, homeButton, dismissButton)
- **Description:** Uses Alpine.js `x-on:click` instead of `_hyperscript`.
- **Fix:** Migrate to `_hyperscript` or document Alpine.js as acceptable for simple event handlers.

### V9: DELETE returns 200 instead of 204 (Very Low)

- **Principle:** #3 -- Self-descriptive HTTP methods
- **File:** `internal/routes/routes_hypermedia.go:172` (`handleCRUDDelete`)
- **Description:** Returns `c.NoContent(200)` instead of `c.NoContent(http.StatusNoContent)` (204).
- **Fix:** Return 204 (verify HTMX handles it correctly with `hx-swap="delete"`).

---

## Areas of Strong Compliance

### Hypermedia-Driven Architecture (Principle #1) ✓
Zero client-side routing or SPA framework. HTML is the API contract throughout. HTMX for partial updates.

### Uniform Interface via `hypermedia.Control` (Principle #2) ✓
Textbook implementation. Factory functions: `FormActions()`, `RowActions()`, `TableRowActions()`, `ResourceActions()`, `BulkActions()`, `NewRowFormActions()`, `EmptyStateAction()`, `CatalogRowAction()`. Controls rendered by `web/components/core/controls.templ`.

### Parent Routes Are Documents (Principle #5) ✓
`/demo` renders `DemoIndexPage()`, `/hypermedia` renders `PatternsIndexPage()`, `/dashboard` renders aggregate stats.

### Server-Side State (Principle #6) ✓
All state in SQLite or in-memory stores. State changes broadcast via SSE. No client-side state management.

### Content Negotiation (Principle #7) ✓
`Vary: HX-Request` set globally. `RenderBaseLayout()` vs `RenderComponent()` based on `HX-Request` header.

### Postel's Law (Principle #9) ✓
No client-side validation. Server validates and returns inline errors. Native HTML attributes only.

### Explicit SQL with Composable Helpers (Principle #10) ✓
`dbrepo.NewSelect()`, `dbrepo.NewWhere()`, `dbrepo.Columns()`, `dbrepo.InsertInto()`, `dbrepo.SetClause()`. No ORM.

### Schema as Code with Traits (Principle #11) ✓
`schema.NewTable("Tasks").Columns(...).WithStatus("draft").WithSortOrder().WithVersion().WithNotes().WithArchive().WithReplacement().WithTimestamps().WithSoftDelete().WithSeedRows(...)`

### Domain Patterns as Primitives (Principle #12) ✓
Standalone functions: `dbrepo.SetCreateTimestamps()`, `dbrepo.InitVersion()`, `dbrepo.SetStatus()`, `dbrepo.IncrementVersion()`.

### Errors Are Hypermedia (Principle #14) ✓
`ErrorContext` with controls, request ID, route. 5 error recovery scenarios demonstrated (transient, validation, conflict, stale data, cascade).

### Structured Observability (Principle #15) ✓
Request IDs via `promolog.CorrelationMiddleware`. Error traces persisted to SQLite. Promote-on-error pattern.

### DaisyUI Semantic Classes (Principle #16) ✓
`btn btn-primary`, `btn-error`, `btn-ghost`, `alert alert-error`, `badge badge-sm badge-outline`, `modal`. No raw Tailwind color classes.

### Native HTML (Principle #17) ✓
Native `<dialog>` with `<form method="dialog">`. Native `<select>`, `<input>` elements.

### Link Relations (Principle #19) ✓
`hypermedia.LinkRelation` system with `rel="up"` chains, `BreadcrumbsFromLinks()`, `?from=` bitmask navigation.

### Go Principles (Principle #22) ✓
Small interfaces (`SessionSettingsStore` has 2 methods, `AppRoutes` has 3). Error wrapping with `fmt.Errorf("...: %w", err)`. Clear, straightforward code.
