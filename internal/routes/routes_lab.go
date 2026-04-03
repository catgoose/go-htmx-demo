// setup:feature:demo

package routes

import (
	"context"
	"fmt"
	"math"
	"math/rand/v2"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"catgoose/dothog/internal/routes/handler"
	"catgoose/dothog/web/views"

	"github.com/catgoose/tavern"
	"github.com/labstack/echo/v4"
)

const labBase = hypermediaBase + "/lab"

// ── Firework constants ──────────────────────────────────────────────────────

const (
	fwWidth      = 400
	fwHeight     = 300
	fwGravity    = 0.035
	fwDefaultFPS = 30
	fwDefaultInt = 5 // intensity 1-10
)

var fwPalettes = [][3]uint8{
	{255, 220, 80},  // gold
	{255, 80, 80},   // red
	{80, 140, 255},  // blue
	{80, 255, 130},  // green
	{230, 80, 255},  // purple
	{255, 255, 230}, // white
	{80, 255, 255},  // cyan
	{255, 150, 60},  // orange
	{255, 110, 200}, // pink
}

// ── Particle types ──────────────────────────────────────────────────────────

type fwParticle struct {
	x, y, vx, vy float64
	life, size    float64
	r, g, b       uint8
	launching     bool
}

type fwSim struct {
	particles []fwParticle
	stars     string
	frame     int
	spawnRate int // frames between launches
	peak      int
}

func newFWSim(intensity int) *fwSim {
	// Generate static stars
	var stars strings.Builder
	for range 60 {
		x := rand.Float64() * fwWidth
		y := rand.Float64() * fwHeight * 0.7
		r := 0.3 + rand.Float64()*0.8
		o := 0.2 + rand.Float64()*0.5
		fmt.Fprintf(&stars, `<circle cx="%.1f" cy="%.1f" r="%.1f" fill="#fff" opacity="%.2f"/>`, x, y, r, o)
	}
	return &fwSim{
		stars:     stars.String(),
		spawnRate: intensityToRate(intensity),
	}
}

func intensityToRate(intensity int) int {
	// intensity 1 = every 50 frames, 10 = every 5 frames
	rate := 55 - intensity*5
	if rate < 5 {
		rate = 5
	}
	return rate
}

func (sim *fwSim) tick() {
	sim.frame++

	// Spawn new firework
	if sim.frame%sim.spawnRate == 0 {
		sim.spawn()
	}

	// Update particles
	alive := sim.particles[:0]
	var explodes []fwParticle
	for i := range sim.particles {
		p := &sim.particles[i]
		p.x += p.vx
		p.y += p.vy
		p.vy += fwGravity

		if p.launching {
			p.life -= 0.025
			if p.vy >= -1.5 || p.life <= 0 {
				explodes = append(explodes, *p)
				continue
			}
		} else {
			p.life -= 0.014
			p.size *= 0.997
		}

		if p.life > 0 && p.y < fwHeight+20 {
			alive = append(alive, *p)
		}
	}
	sim.particles = alive

	// Process explosions
	for _, p := range explodes {
		sim.explode(p)
	}

	if len(sim.particles) > sim.peak {
		sim.peak = len(sim.particles)
	}
}

func (sim *fwSim) spawn() {
	x := fwWidth * (0.15 + rand.Float64()*0.7)
	vy := -(6.0 + rand.Float64()*3.5)
	vx := (rand.Float64() - 0.5) * 1.5
	c := fwPalettes[rand.IntN(len(fwPalettes))]
	sim.particles = append(sim.particles, fwParticle{
		x: x, y: fwHeight,
		vx: vx, vy: vy,
		r: c[0], g: c[1], b: c[2],
		life: 1.0, size: 2.5,
		launching: true,
	})
}

func (sim *fwSim) explode(p fwParticle) {
	count := 60 + rand.IntN(80)
	willow := rand.Float64() < 0.3 // 30% chance of willow type

	for range count {
		angle := rand.Float64() * 2 * math.Pi
		speed := 1.0 + rand.Float64()*3.5
		if willow {
			speed *= 0.6
		}

		// Slight color variation
		r := clampByte(int(p.r) + rand.IntN(50) - 25)
		g := clampByte(int(p.g) + rand.IntN(50) - 25)
		b := clampByte(int(p.b) + rand.IntN(50) - 25)

		life := 0.7 + rand.Float64()*0.3
		if willow {
			life = 0.9 + rand.Float64()*0.1
		}

		sim.particles = append(sim.particles, fwParticle{
			x: p.x, y: p.y,
			vx: math.Cos(angle)*speed + p.vx*0.3,
			vy: math.Sin(angle)*speed + p.vy*0.3,
			r: r, g: g, b: b,
			life: life,
			size: 1.0 + rand.Float64()*1.5,
		})
	}

	// Add sparkle particles
	sparkles := 10 + rand.IntN(20)
	for range sparkles {
		angle := rand.Float64() * 2 * math.Pi
		speed := 0.5 + rand.Float64()*2
		sim.particles = append(sim.particles, fwParticle{
			x: p.x, y: p.y,
			vx: math.Cos(angle) * speed,
			vy: math.Sin(angle) * speed,
			r: 255, g: 255, b: 240,
			life: 0.4 + rand.Float64()*0.3,
			size: 0.5 + rand.Float64()*0.5,
		})
	}
}

func (sim *fwSim) render() string {
	var buf strings.Builder
	buf.Grow(len(sim.particles)*90 + len(sim.stars) + 300)

	fmt.Fprintf(&buf, `<svg viewBox="0 0 %d %d" xmlns="http://www.w3.org/2000/svg" style="display:block;width:100%%">`, fwWidth, fwHeight)

	// Night sky gradient
	buf.WriteString(`<defs><linearGradient id="sky" x1="0" y1="0" x2="0" y2="1">`)
	buf.WriteString(`<stop offset="0" stop-color="#0a0a2e"/>`)
	buf.WriteString(`<stop offset="1" stop-color="#1a1a3e"/>`)
	buf.WriteString(`</linearGradient></defs>`)
	fmt.Fprintf(&buf, `<rect width="%d" height="%d" fill="url(#sky)"/>`, fwWidth, fwHeight)

	buf.WriteString(sim.stars)

	// Render particles (glow + core for each)
	for i := range sim.particles {
		p := &sim.particles[i]
		o := p.life
		if o > 1 {
			o = 1
		}
		// Glow
		fmt.Fprintf(&buf, `<circle cx="%.1f" cy="%.1f" r="%.1f" fill="rgb(%d,%d,%d)" opacity="%.2f"/>`,
			p.x, p.y, p.size*2.5, p.r, p.g, p.b, o*0.25)
		// Core
		fmt.Fprintf(&buf, `<circle cx="%.1f" cy="%.1f" r="%.1f" fill="rgb(%d,%d,%d)" opacity="%.2f"/>`,
			p.x, p.y, p.size*p.life, p.r, p.g, p.b, o)
	}

	buf.WriteString("</svg>")
	return buf.String()
}

func clampByte(v int) uint8 {
	if v < 0 {
		return 0
	}
	if v > 255 {
		return 255
	}
	return uint8(v)
}

// ── State ───────────────────────────────────────────────────────────────────

var labState struct {
	cancel    context.CancelFunc
	fps       int
	intensity int
	mu        sync.Mutex
}

// ── Routes ──────────────────────────────────────────────────────────────────

func (ar *appRoutes) initLabRoutes(broker *tavern.SSEBroker) {
	labState.fps = fwDefaultFPS
	labState.intensity = fwDefaultInt
	ar.e.GET(labBase, handleLabPage(ar.ctx, broker))
	ar.e.POST(labBase+"/settings", handleLabSettings(ar.ctx, broker))
	ar.e.POST(labBase+"/reset", handleLabReset(ar.ctx, broker))
	ar.e.GET("/sse/lab", handleSSELab(broker))
}

func handleLabPage(appCtx context.Context, broker *tavern.SSEBroker) echo.HandlerFunc {
	return func(c echo.Context) error {
		sky := renderEmptySky()

		labState.mu.Lock()
		if labState.cancel != nil {
			labState.cancel()
		}
		ctx, cancel := context.WithCancel(appCtx)
		labState.cancel = cancel
		labState.mu.Unlock()

		go publishFireworks(ctx, broker)

		return handler.RenderBaseLayout(c, views.LabPage(sky))
	}
}

func handleLabSettings(appCtx context.Context, broker *tavern.SSEBroker) echo.HandlerFunc {
	return func(c echo.Context) error {
		fps, _ := strconv.Atoi(c.FormValue("fps"))
		if fps < 10 {
			fps = 10
		} else if fps > 60 {
			fps = 60
		}
		intensity, _ := strconv.Atoi(c.FormValue("intensity"))
		if intensity < 1 {
			intensity = 1
		} else if intensity > 10 {
			intensity = 10
		}

		labState.mu.Lock()
		labState.fps = fps
		labState.intensity = intensity
		if labState.cancel != nil {
			labState.cancel()
		}
		ctx, cancel := context.WithCancel(appCtx)
		labState.cancel = cancel
		labState.mu.Unlock()

		go publishFireworks(ctx, broker)

		return handler.RenderComponent(c, views.LabSettingsApplied())
	}
}

func handleLabReset(appCtx context.Context, broker *tavern.SSEBroker) echo.HandlerFunc {
	return func(c echo.Context) error {
		labState.mu.Lock()
		labState.fps = fwDefaultFPS
		labState.intensity = fwDefaultInt
		if labState.cancel != nil {
			labState.cancel()
		}
		ctx, cancel := context.WithCancel(appCtx)
		labState.cancel = cancel
		labState.mu.Unlock()

		go publishFireworks(ctx, broker)

		return handler.RenderComponent(c, views.LabResetResponse(fwDefaultFPS, fwDefaultInt))
	}
}

func handleSSELab(broker *tavern.SSEBroker) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Content-Type", "text/event-stream")
		c.Response().Header().Set("Cache-Control", "no-cache")
		c.Response().Header().Set("Connection", "keep-alive")
		c.Response().WriteHeader(http.StatusOK)
		flusher, ok := c.Response().Writer.(http.Flusher)
		if !ok {
			return fmt.Errorf("streaming not supported")
		}
		flusher.Flush()

		ch, unsub := broker.Subscribe(TopicLab)
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

func publishFireworks(ctx context.Context, broker *tavern.SSEBroker) {
	for !broker.HasSubscribers(TopicLab) {
		select {
		case <-ctx.Done():
			return
		case <-time.After(50 * time.Millisecond):
		}
	}

	labState.mu.Lock()
	intensity := labState.intensity
	labState.mu.Unlock()

	sim := newFWSim(intensity)

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		sim.tick()
		frame := sim.render()

		var buf strings.Builder
		buf.WriteString(`<div id="lab-canvas" hx-swap-oob="innerHTML">`)
		buf.WriteString(frame)
		buf.WriteString(`</div>`)
		fmt.Fprintf(&buf, `<div id="lab-status" hx-swap-oob="innerHTML"><span class="font-mono">%s</span> particles · peak <span class="font-mono">%s</span></div>`,
			formatCommas(len(sim.particles)), formatCommas(sim.peak))

		msg := tavern.NewSSEMessage("lab-stream", buf.String()).String()
		broker.Publish(TopicLab, msg)

		labState.mu.Lock()
		fps := labState.fps
		newIntensity := labState.intensity
		labState.mu.Unlock()

		if newIntensity != intensity {
			intensity = newIntensity
			sim.spawnRate = intensityToRate(intensity)
		}

		time.Sleep(time.Second / time.Duration(fps))
	}
}

// ── Helpers ─────────────────────────────────────────────────────────────────

func renderEmptySky() string {
	var buf strings.Builder
	fmt.Fprintf(&buf, `<svg viewBox="0 0 %d %d" xmlns="http://www.w3.org/2000/svg" style="display:block;width:100%%">`, fwWidth, fwHeight)
	buf.WriteString(`<defs><linearGradient id="sky" x1="0" y1="0" x2="0" y2="1">`)
	buf.WriteString(`<stop offset="0" stop-color="#0a0a2e"/>`)
	buf.WriteString(`<stop offset="1" stop-color="#1a1a3e"/>`)
	buf.WriteString(`</linearGradient></defs>`)
	fmt.Fprintf(&buf, `<rect width="%d" height="%d" fill="url(#sky)"/>`, fwWidth, fwHeight)
	buf.WriteString("</svg>")
	return buf.String()
}

func formatCommas(n int) string {
	s := fmt.Sprintf("%d", n)
	if len(s) <= 3 {
		return s
	}
	var b strings.Builder
	remainder := len(s) % 3
	if remainder > 0 {
		b.WriteString(s[:remainder])
	}
	for i := remainder; i < len(s); i += 3 {
		if b.Len() > 0 {
			b.WriteByte(',')
		}
		b.WriteString(s[i : i+3])
	}
	return b.String()
}
