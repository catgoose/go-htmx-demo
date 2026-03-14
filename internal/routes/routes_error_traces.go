package routes

import (
	"net/url"
	"strconv"

	"catgoose/dothog/internal/requestlog"
	"catgoose/dothog/internal/routes/handler"
	"catgoose/dothog/internal/routes/hypermedia"
	"catgoose/dothog/web/views"

	hx "catgoose/dothog/internal/routes/htmx"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

const errorTracesBase = "/admin/error-traces"

func (ar *appRoutes) initErrorTracesRoutes() {
	if ar.reqLogStore == nil {
		return
	}
	ar.e.GET(errorTracesBase, ar.handleErrorTracesPage)
	ar.e.GET(errorTracesBase+"/list", ar.handleErrorTracesList)
	ar.e.GET(errorTracesBase+"/:requestID", ar.handleErrorTraceDetail)
	ar.e.DELETE(errorTracesBase+"/:requestID", ar.handleErrorTraceDelete)
}

func (ar *appRoutes) handleErrorTracesPage(c echo.Context) error {
	bar, container, err := ar.buildErrorTracesContent(c)
	if err != nil {
		return handler.HandleHypermediaError(c, 500, "Failed to load error traces", err)
	}
	return handler.RenderBaseLayout(c, views.ErrorTracesPage(bar, container))
}

func (ar *appRoutes) handleErrorTracesList(c echo.Context) error {
	_, container, err := ar.buildErrorTracesContent(c)
	if err != nil {
		return handler.HandleHypermediaError(c, 500, "Failed to load error traces", err)
	}
	if hx.IsHTMX(c) {
		pushURL := errorTracesBase
		if q := c.Request().URL.RawQuery; q != "" {
			pushURL += "?" + q
		}
		hx.ReplaceURL(c, pushURL)
	}
	return handler.RenderComponent(c, container)
}

func (ar *appRoutes) handleErrorTraceDetail(c echo.Context) error {
	requestID := c.Param("requestID")
	trace := ar.reqLogStore.Get(requestID)
	if trace == nil {
		return handler.HandleHypermediaError(c, 404, "Error trace not found", nil)
	}
	return handler.RenderComponent(c, views.ErrorTraceDetailContent(trace))
}

func (ar *appRoutes) handleErrorTraceDelete(c echo.Context) error {
	requestID := c.Param("requestID")
	if err := ar.reqLogStore.DeleteTrace(requestID); err != nil {
		return handler.HandleHypermediaError(c, 500, "Failed to delete trace", err)
	}
	// Re-apply current filters from HX-Current-URL
	if raw := c.Request().Header.Get("HX-Current-URL"); raw != "" {
		if u, err := url.Parse(raw); err == nil && u.RawQuery != "" {
			c.Request().URL.RawQuery = u.RawQuery
		}
	}
	_, container, err := ar.buildErrorTracesContent(c)
	if err != nil {
		return handler.HandleHypermediaError(c, 500, "Failed to reload traces", err)
	}
	return handler.RenderComponent(c, container)
}

func (ar *appRoutes) buildErrorTracesContent(c echo.Context) (hypermedia.FilterBar, templ.Component, error) {
	const perPage = 20
	q := c.QueryParam("q")
	status := c.QueryParam("status")
	sort := c.QueryParam("sort")
	dir := c.QueryParam("dir")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}

	traces, total, err := ar.reqLogStore.ListTraces(q, status, sort, dir, page, perPage)
	if err != nil {
		return hypermedia.FilterBar{}, nil, err
	}

	target := "#error-traces-table-container"
	listURL := errorTracesBase + "/list"

	bar := hypermedia.NewFilterBar(listURL, target,
		hypermedia.SearchField("q", "Search routes, errors, request IDs\u2026", q),
		hypermedia.SelectField("status", "Status", status,
			hypermedia.SelectOptions(status,
				"", "All",
				"4xx", "4xx Client",
				"5xx", "5xx Server",
			)),
	)

	sortBase := traceStripParams(c.Request().URL, "sort", "dir")
	cols := []hypermedia.TableCol{
		hypermedia.SortableCol("CreatedAt", "Time", sort, dir, sortBase, target, "#filter-form"),
		hypermedia.SortableCol("StatusCode", "Status", sort, dir, sortBase, target, "#filter-form"),
		hypermedia.SortableCol("Method", "Method", sort, dir, sortBase, target, "#filter-form"),
		hypermedia.SortableCol("Route", "Route", sort, dir, sortBase, target, "#filter-form"),
		{Label: "Error"},
		{Label: "IP"},
		{Label: ""},
	}

	pageBase := traceStripParams(c.Request().URL, "page")
	info := hypermedia.PageInfo{
		Page:       page,
		PerPage:    perPage,
		TotalItems: total,
		TotalPages: hypermedia.ComputeTotalPages(total, perPage),
		BaseURL:    pageBase,
		Target:     target,
		Include:    "#filter-form",
	}

	body := views.ErrorTracesBody(traces)
	container := views.ErrorTracesTableContainer(cols, body, info)
	return bar, container, nil
}

// traceStripParams returns a copy of u with the named query params removed.
func traceStripParams(u *url.URL, params ...string) string {
	cp := *u
	q := cp.Query()
	for _, p := range params {
		q.Del(p)
	}
	cp.RawQuery = q.Encode()
	return cp.String()
}

// SeedErrorTraces inserts demo error trace data into the store.
func SeedErrorTraces(store *requestlog.Store) {
	traces := []requestlog.ErrorTrace{
		{
			RequestID: "demo-a1b2c3d4e5f6", ErrorChain: "get user 42: sql: no rows in result set",
			StatusCode: 404, Route: "/api/users/42", Method: "GET",
			UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
			RemoteIP: "192.168.1.100", UserID: "alice@contoso.com",
			Entries: []requestlog.Entry{
				{Level: "INFO", Message: "Request started", Attrs: "method=GET path=/api/users/42"},
				{Level: "INFO", Message: "Querying user by ID", Attrs: "user_id=42"},
				{Level: "ERROR", Message: "User not found", Attrs: "user_id=42 error=sql: no rows in result set"},
			},
		},
		{
			RequestID: "demo-f6e5d4c3b2a1", ErrorChain: "process order: validate inventory: insufficient stock for SKU-1234",
			StatusCode: 422, Route: "/api/orders", Method: "POST",
			UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7)",
			RemoteIP: "10.0.0.50", UserID: "bob@contoso.com",
			Entries: []requestlog.Entry{
				{Level: "INFO", Message: "Request started", Attrs: "method=POST path=/api/orders"},
				{Level: "INFO", Message: "Parsing order payload", Attrs: "items=3"},
				{Level: "INFO", Message: "Validating inventory", Attrs: "sku=SKU-1234 requested=10"},
				{Level: "WARN", Message: "Low stock detected", Attrs: "sku=SKU-1234 available=2 requested=10"},
				{Level: "ERROR", Message: "Insufficient stock", Attrs: "sku=SKU-1234"},
			},
		},
		{
			RequestID: "demo-1122334455aa", ErrorChain: "render dashboard: query metrics: context deadline exceeded",
			StatusCode: 504, Route: "/dashboard", Method: "GET",
			UserAgent: "Mozilla/5.0 (X11; Linux x86_64; rv:109.0)",
			RemoteIP: "172.16.0.25", UserID: "charlie@contoso.com",
			Entries: []requestlog.Entry{
				{Level: "INFO", Message: "Request started", Attrs: "method=GET path=/dashboard"},
				{Level: "INFO", Message: "Fetching metrics", Attrs: "range=7d"},
				{Level: "WARN", Message: "Query slow", Attrs: "elapsed_ms=4500"},
				{Level: "ERROR", Message: "Context deadline exceeded", Attrs: "timeout=5s"},
			},
		},
		{
			RequestID: "demo-aabb11223344", ErrorChain: "upload file: multipart: NextPart: unexpected EOF",
			StatusCode: 400, Route: "/api/files/upload", Method: "POST",
			UserAgent: "curl/8.1.2",
			RemoteIP: "192.168.1.55", UserID: "",
			Entries: []requestlog.Entry{
				{Level: "INFO", Message: "Request started", Attrs: "method=POST path=/api/files/upload"},
				{Level: "INFO", Message: "Parsing multipart form", Attrs: "content_length=1048576"},
				{Level: "ERROR", Message: "Multipart parse failed", Attrs: "error=unexpected EOF"},
			},
		},
		{
			RequestID: "demo-55667788ccdd", ErrorChain: "save settings: database is locked",
			StatusCode: 500, Route: "/settings/theme", Method: "POST",
			UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64)",
			RemoteIP: "10.0.0.12", UserID: "dana@contoso.com",
			Entries: []requestlog.Entry{
				{Level: "INFO", Message: "Request started", Attrs: "method=POST path=/settings/theme"},
				{Level: "INFO", Message: "Updating theme", Attrs: "theme=dark session=abc-123"},
				{Level: "ERROR", Message: "Database write failed", Attrs: "error=database is locked table=SessionSettings"},
			},
		},
		{
			RequestID: "demo-99aabbccddee", ErrorChain: "fetch report: connect: connection refused",
			StatusCode: 502, Route: "/api/reports/monthly", Method: "GET",
			UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7)",
			RemoteIP: "172.16.0.30", UserID: "eve@contoso.com",
			Entries: []requestlog.Entry{
				{Level: "INFO", Message: "Request started", Attrs: "method=GET path=/api/reports/monthly"},
				{Level: "INFO", Message: "Calling reporting service", Attrs: "url=http://reports-svc:8080/monthly"},
				{Level: "ERROR", Message: "Upstream connection refused", Attrs: "host=reports-svc:8080 error=connection refused"},
			},
		},
		{
			RequestID: "demo-eeff00112233", ErrorChain: "authenticate: token expired at 2026-03-13T23:59:59Z",
			StatusCode: 401, Route: "/api/protected/data", Method: "GET",
			UserAgent: "PostmanRuntime/7.32.3",
			RemoteIP: "192.168.1.200", UserID: "",
			Entries: []requestlog.Entry{
				{Level: "INFO", Message: "Request started", Attrs: "method=GET path=/api/protected/data"},
				{Level: "WARN", Message: "Token validation failed", Attrs: "reason=expired exp=2026-03-13T23:59:59Z"},
				{Level: "ERROR", Message: "Authentication failed", Attrs: "error=token expired"},
			},
		},
		{
			RequestID: "demo-44556677aabb", ErrorChain: "list items: UNIQUE constraint failed: items.name",
			StatusCode: 409, Route: "/demo/inventory/items", Method: "POST",
			UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64)",
			RemoteIP: "10.0.0.5", UserID: "frank@contoso.com",
			Entries: []requestlog.Entry{
				{Level: "INFO", Message: "Request started", Attrs: "method=POST path=/demo/inventory/items"},
				{Level: "INFO", Message: "Creating item", Attrs: "name=Widget category=Electronics"},
				{Level: "ERROR", Message: "Insert failed", Attrs: "error=UNIQUE constraint failed: items.name name=Widget"},
			},
		},
	}

	for i := range traces {
		traces[i].Entries[0].Time = traces[i].Entries[0].Time // time.Time zero is fine for demo
		store.Promote(traces[i])
	}
}
