// setup:feature:demo

package routes

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"catgoose/go-htmx-demo/internals/demo"
	"catgoose/go-htmx-demo/internals/routes/handler"
	"catgoose/go-htmx-demo/internals/ssebroker"
	"catgoose/go-htmx-demo/web/views"

	"github.com/labstack/echo/v4"
)

type feedRoutes struct {
	actLog *demo.ActivityLog
	broker *ssebroker.SSEBroker
}

func (ar *appRoutes) initFeedRoutes(actLog *demo.ActivityLog, broker *ssebroker.SSEBroker) {
	f := &feedRoutes{actLog: actLog, broker: broker}
	ar.e.GET("/tables/feed", f.handleFeedPage)
	ar.e.GET("/tables/feed/more", f.handleFeedMore)
	ar.e.GET("/sse/activity", f.handleActivitySSE)
}

func (f *feedRoutes) handleFeedPage(c echo.Context) error {
	events := f.actLog.Recent(20)
	lastID := 0
	if len(events) > 0 {
		lastID = events[len(events)-1].ID
	}
	return handler.RenderBaseLayout(c, views.FeedPage(events, lastID))
}

func (f *feedRoutes) handleFeedMore(c echo.Context) error {
	beforeID, _ := strconv.Atoi(c.QueryParam("before"))
	events := f.actLog.Recent(50)
	// Filter events with ID < beforeID
	var filtered []demo.ActivityEvent
	for _, e := range events {
		if e.ID < beforeID {
			filtered = append(filtered, e)
		}
	}
	// Take last 20
	if len(filtered) > 20 {
		filtered = filtered[:20]
	}
	lastID := 0
	if len(filtered) > 0 {
		lastID = filtered[len(filtered)-1].ID
	}
	hasMore := len(filtered) == 20
	return handler.RenderComponent(c, views.FeedMoreItems(filtered, lastID, hasMore))
}

func (f *feedRoutes) handleActivitySSE(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "text/event-stream")
	c.Response().Header().Set("Cache-Control", "no-cache")
	c.Response().Header().Set("Connection", "keep-alive")
	c.Response().WriteHeader(http.StatusOK)

	flusher, ok := c.Response().Writer.(http.Flusher)
	if !ok {
		return fmt.Errorf("streaming unsupported")
	}

	ch, unsub := f.broker.Subscribe(ssebroker.TopicActivityFeed)
	defer unsub()

	ctx := c.Request().Context()
	for {
		select {
		case <-ctx.Done():
			return nil
		case msg, ok := <-ch:
			if !ok {
				return nil
			}
			fmt.Fprint(c.Response(), msg)
			flusher.Flush()
		}
	}
}

// BroadcastActivity publishes an activity event to the SSE feed.
func BroadcastActivity(broker *ssebroker.SSEBroker, e demo.ActivityEvent) {
	if !broker.HasSubscribers(ssebroker.TopicActivityFeed) {
		return
	}
	buf := statsBufPool.Get().(*bytes.Buffer)
	buf.Reset()
	if err := views.FeedItemOOB(e).Render(context.Background(), buf); err != nil {
		statsBufPool.Put(buf)
		return
	}
	msg := ssebroker.NewSSEMessage("activity-event", buf.String()).String()
	statsBufPool.Put(buf)
	broker.Publish(ssebroker.TopicActivityFeed, msg)
}
