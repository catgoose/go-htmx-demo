// setup:feature:demo

package routes

import (
	"bytes"
	"context"
	"fmt"
	"math"
	"math/rand/v2"
	"net/http"
	"sync"
	"time"

	"catgoose/dothog/internal/routes/handler"
	"catgoose/dothog/web/views"

	"github.com/catgoose/tavern"
	"github.com/labstack/echo/v4"
)

const numericalBase = hypermediaBase + "/numerical"

// ── Simulation state ────────────────────────────────────────────────────────

type numSim struct {
	txnSec    float64
	revenue   float64
	users     float64
	queue     float64
	cacheHit  float64
	errors    float64
	p99       float64
	cpu       float64
	mem       float64
	uptime    time.Time
	deploys   int
	sla       float64
	incidents int

	prevTxn   float64
	prevUsers float64
	prevQueue float64
	prevCache float64
	prevP99   float64
	prevCPU   float64
}

func newNumSim() *numSim {
	return &numSim{
		txnSec:   3200,
		revenue:  48500,
		users:    11200,
		queue:    12,
		cacheHit: 96.5,
		errors:   87,
		p99:      28,
		cpu:      42,
		mem:      9.8,
		uptime:   time.Now().Add(-14*24*time.Hour - 7*time.Hour - 23*time.Minute),
		deploys:  2,
		sla:      99.97,
	}
}

func (s *numSim) tick() {
	s.prevTxn = s.txnSec
	s.prevUsers = s.users
	s.prevQueue = s.queue
	s.prevCache = s.cacheHit
	s.prevP99 = s.p99
	s.prevCPU = s.cpu

	s.txnSec = clampF(s.txnSec+(rand.Float64()-0.48)*300, 800, 9000)
	s.revenue += s.txnSec * 0.035 * (0.8 + rand.Float64()*0.4)
	s.users = clampF(s.users+(rand.Float64()-0.5)*150, 3000, 30000)
	s.queue = clampF(s.queue+(rand.Float64()-0.55)*4, 0, 200)
	s.cacheHit = clampF(s.cacheHit+(rand.Float64()-0.48)*0.8, 82, 99.9)
	s.p99 = clampF(s.p99+(rand.Float64()-0.48)*6, 5, 500)
	s.cpu = clampF(s.cpu+(rand.Float64()-0.5)*8, 5, 98)
	s.mem = clampF(s.mem+(rand.Float64()-0.5)*0.2, 4, 15.5)

	// Correlate p99 with cpu load
	if s.cpu > 80 {
		s.p99 += (s.cpu - 80) * 0.5
	}

	// Error accumulation with occasional spike
	if rand.Float64() < 0.03 {
		s.errors += float64(10 + rand.IntN(30))
		s.incidents++
	} else {
		s.errors += float64(rand.IntN(3))
	}

	// SLA derived from error rate
	errorRate := s.errors / math.Max(s.txnSec*86400, 1) * 100
	s.sla = clampF(100.0-errorRate*5, 95, 100)

	// Occasional deploy
	if rand.Float64() < 0.001 {
		s.deploys++
	}
}

func (s *numSim) buildTiles() []views.NumTile {
	uptimeDur := time.Since(s.uptime)
	days := int(uptimeDur.Hours()) / 24
	hours := int(uptimeDur.Hours()) % 24
	mins := int(uptimeDur.Minutes()) % 60

	return []views.NumTile{
		{
			ID: "num-txn", Title: "Transactions/sec",
			Value:   fmtCommas(int(s.txnSec)),
			Delta:   fmtDelta(s.txnSec, s.prevTxn),
			DeltaUp: s.txnSec >= s.prevTxn,
			Color:   "info",
		},
		{
			ID: "num-revenue", Title: "Revenue Today",
			Value:    fmt.Sprintf("$%s", fmtMoney(s.revenue)),
			Delta:    fmt.Sprintf("$%.0f/min", s.txnSec*0.035*60),
			DeltaUp:  true,
			Subtitle: "accumulating",
			Color:    "success",
		},
		{
			ID: "num-users", Title: "Active Users",
			Value:   fmtCommas(int(s.users)),
			Delta:   fmtDelta(s.users, s.prevUsers),
			DeltaUp: s.users >= s.prevUsers,
			Color:   "info",
		},
		{
			ID: "num-queue", Title: "Queue Depth",
			Value:   fmt.Sprintf("%d", int(s.queue)),
			Delta:   fmtDelta(s.queue, s.prevQueue),
			DeltaUp: s.queue <= s.prevQueue, // lower is better
			Color:   queueColor(s.queue),
		},
		{
			ID: "num-cache", Title: "Cache Hit Rate",
			Value:   fmt.Sprintf("%.1f%%", s.cacheHit),
			Delta:   fmtDeltaPct(s.cacheHit, s.prevCache),
			DeltaUp: s.cacheHit >= s.prevCache,
			Color:   cacheColor(s.cacheHit),
		},
		{
			ID: "num-errors", Title: "Errors (24h)",
			Value:    fmtCommas(int(s.errors)),
			Subtitle: fmt.Sprintf("%d incidents", s.incidents),
			Color:    errorCountColor(s.errors),
		},
		{
			ID: "num-p99", Title: "P99 Latency",
			Value:   fmt.Sprintf("%.0fms", s.p99),
			Delta:   fmtDelta(s.p99, s.prevP99),
			DeltaUp: s.p99 <= s.prevP99, // lower is better
			Color:   latencyColor(s.p99),
		},
		{
			ID: "num-cpu", Title: "CPU Load",
			Value:   fmt.Sprintf("%.0f%%", s.cpu),
			Delta:   fmtDeltaPct(s.cpu, s.prevCPU),
			DeltaUp: s.cpu <= s.prevCPU, // lower is better
			Color:   cpuColor(s.cpu),
		},
		{
			ID: "num-mem", Title: "Memory",
			Value:    fmt.Sprintf("%.1f GB", s.mem),
			Subtitle: fmt.Sprintf("of 16 GB (%.0f%%)", s.mem/16*100),
			Color:    memColor(s.mem),
		},
		{
			ID: "num-uptime", Title: "Uptime",
			Value:   fmt.Sprintf("%dd %dh %dm", days, hours, mins),
			Neutral: true,
			Color:   "success",
		},
		{
			ID: "num-deploys", Title: "Deploys Today",
			Value:   fmt.Sprintf("%d", s.deploys),
			Neutral: true,
			Color:   "info",
		},
		{
			ID: "num-sla", Title: "SLA Compliance",
			Value: fmt.Sprintf("%.2f%%", s.sla),
			Color: slaColor(s.sla),
		},
	}
}

// ── Color thresholds ────────────────────────────────────────────────────────

func queueColor(v float64) string {
	switch {
	case v > 100:
		return "error"
	case v > 50:
		return "warning"
	default:
		return "success"
	}
}

func cacheColor(v float64) string {
	switch {
	case v < 85:
		return "error"
	case v < 95:
		return "warning"
	default:
		return "success"
	}
}

func errorCountColor(v float64) string {
	switch {
	case v > 500:
		return "error"
	case v > 200:
		return "warning"
	default:
		return ""
	}
}

func latencyColor(v float64) string {
	switch {
	case v > 200:
		return "error"
	case v > 100:
		return "warning"
	default:
		return "success"
	}
}

func cpuColor(v float64) string {
	switch {
	case v > 85:
		return "error"
	case v > 70:
		return "warning"
	default:
		return "success"
	}
}

func memColor(v float64) string {
	switch {
	case v > 14:
		return "error"
	case v > 12:
		return "warning"
	default:
		return ""
	}
}

func slaColor(v float64) string {
	switch {
	case v < 99.0:
		return "error"
	case v < 99.9:
		return "warning"
	default:
		return "success"
	}
}

// ── Formatting helpers ──────────────────────────────────────────────────────

func fmtCommas(n int) string {
	s := fmt.Sprintf("%d", n)
	if len(s) <= 3 {
		return s
	}
	var b []byte
	rem := len(s) % 3
	if rem > 0 {
		b = append(b, s[:rem]...)
	}
	for i := rem; i < len(s); i += 3 {
		if len(b) > 0 {
			b = append(b, ',')
		}
		b = append(b, s[i:i+3]...)
	}
	return string(b)
}

func fmtMoney(v float64) string {
	whole := int(v)
	cents := int((v - float64(whole)) * 100)
	return fmt.Sprintf("%s.%02d", fmtCommas(whole), cents)
}

func fmtDelta(cur, prev float64) string {
	d := cur - prev
	if math.Abs(d) < 0.5 {
		return "—"
	}
	return fmt.Sprintf("%.0f", math.Abs(d))
}

func fmtDeltaPct(cur, prev float64) string {
	d := cur - prev
	if math.Abs(d) < 0.01 {
		return "—"
	}
	return fmt.Sprintf("%.1f%%", math.Abs(d))
}

func clampF(v, lo, hi float64) float64 {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

// ── Routes ──────────────────────────────────────────────────────────────────

var numBufPool = sync.Pool{New: func() any { return new(bytes.Buffer) }}

func (ar *appRoutes) initNumericalRoutes(broker *tavern.SSEBroker) {
	ar.e.GET(numericalBase, ar.handleNumericalPage)
	ar.e.GET(numericalBase+"/sse-connect", handleNumericalSSEConnect)
	ar.e.GET("/sse/numerical", handleSSENumerical(broker))

	go ar.publishNumerical(broker)
}

func (ar *appRoutes) handleNumericalPage(c echo.Context) error {
	sim := newNumSim()
	tiles := sim.buildTiles()
	return handler.RenderBaseLayout(c, views.NumericalPage(tiles))
}

func handleNumericalSSEConnect(c echo.Context) error {
	return handler.RenderComponent(c, views.NumericalSSEBlock())
}

func handleSSENumerical(broker *tavern.SSEBroker) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Content-Type", "text/event-stream")
		c.Response().Header().Set("Cache-Control", "no-cache")
		c.Response().Header().Set("Connection", "keep-alive")
		c.Response().WriteHeader(200)
		flusher, ok := c.Response().Writer.(http.Flusher)
		if !ok {
			return fmt.Errorf("streaming not supported")
		}
		flusher.Flush()

		ch, unsub := broker.Subscribe(TopicNumericalDash)
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
}

// ── Publisher ────────────────────────────────────────────────────────────────

func (ar *appRoutes) publishNumerical(broker *tavern.SSEBroker) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	sim := newNumSim()
	ctx := context.Background()

	for {
		select {
		case <-ar.ctx.Done():
			return
		case <-ticker.C:
			if !broker.HasSubscribers(TopicNumericalDash) {
				continue
			}

			sim.tick()
			tiles := sim.buildTiles()

			buf := numBufPool.Get().(*bytes.Buffer)
			buf.Reset()
			if err := views.NumericalOOB(tiles).Render(ctx, buf); err != nil {
				numBufPool.Put(buf)
				continue
			}

			msg := tavern.NewSSEMessage("numerical-dash", buf.String()).String()
			numBufPool.Put(buf)
			broker.Publish(TopicNumericalDash, msg)
		}
	}
}
