# go-htmx-demo

<!--toc:start-->
- [go-htmx-demo](#go-htmx-demo)
  - [Features](#features)
  - [Quick Start](#quick-start)
    - [From Release Binary](#from-release-binary)
    - [From Source](#from-source)
  - [Tech Stack](#tech-stack)
  - [Project Structure](#project-structure)
  - [Template Setup](#template-setup)
    - [Interactive Setup (with gum)](#interactive-setup-with-gum)
    - [Non-interactive Setup](#non-interactive-setup)
  - [Development](#development)
    - [Prerequisites](#prerequisites)
    - [Running the Dev Server](#running-the-dev-server)
    - [HTTPS Development Setup](#https-development-setup)
  - [Testing](#testing)
  - [Mage Targets](#mage-targets)
<!--toc:end-->

A demonstration of building modern, interactive web applications with Go and HTMX. The app runs as a single binary with all assets embedded -- no external dependencies or configuration required.

## Features

- **SSE Real-time Dashboard** -- Live system stats, metrics, service health, and event streams powered by Server-Sent Events with HTMX
- **Hypermedia Controls** -- Buttons, toggles, and interactive UI patterns using HTMX attributes
- **CRUD Operations** -- Create, read, update, and delete with inline editing and form handling
- **Interactive Tables** -- Sorting, filtering, pagination, and bulk operations on a SQLite-backed demo dataset
- **List Manipulation** -- Dynamic list management patterns (add, remove, reorder)
- **State Management** -- Client-side state patterns using HTMX and Hyperscript
- **Out-of-Band Updates** -- HTMX OOB swaps for updating multiple page regions from a single response
- **Polling** -- Polling-based counter as an alternative to SSE

## Quick Start

### From Release Binary

Download the latest release for your platform from the [Releases](../../releases) page and run it:

```bash
# Linux
chmod +x go-htmx-demo-linux-amd64
./go-htmx-demo-linux-amd64

# Windows
go-htmx-demo-windows-amd64.exe
```

The app starts on `http://localhost:8080` by default. Override the port with:

```bash
SERVER_LISTEN_PORT=3000 ./go-htmx-demo-linux-amd64
```

### From Source

```bash
git clone https://github.com/catgoose/go-htmx-demo.git
cd go-htmx-demo
go build -o go-htmx-demo .
./go-htmx-demo
```

## Tech Stack

- [**Go**](https://go.dev/) -- Backend language
- [**Echo**](https://echo.labstack.com/) -- High performance, minimalist Go web framework
- [**HTMX**](https://htmx.org/) -- Frontend interactivity with minimal JavaScript
- [**templ**](https://templ.guide/) -- Type-safe HTML templating for Go
- [**Tailwind CSS**](https://tailwindcss.com/) -- Utility-first CSS framework
- [**DaisyUI**](https://daisyui.com/) -- Tailwind CSS component library
- [**Hyperscript**](https://hyperscript.org/) -- Lightweight scripting for DOM interactions
- [**SQLite**](https://www.sqlite.org/) -- Embedded database for demo data
- [**Air**](https://github.com/air-verse/air) -- Live reloading for Go development
- [**Mage**](https://magefile.org/) -- Make/rake-like build tool for Go

## Project Structure

```
go-htmx-demo/
├── main.go                    # Application entrypoint
├── magefile.go                # Build automation targets
├── internals/
│   ├── config/                # Configuration management
│   ├── logger/                # Structured logging (slog)
│   ├── routes/                # HTTP routes, handlers, middleware
│   │   ├── handler/           # Component rendering utilities
│   │   ├── middleware/        # Request validation, error handling
│   │   ├── response/         # HTMX OOB response builders
│   │   ├── hypermedia/       # Navigation, filters, table state
│   │   └── routes_realtime.go # SSE endpoints and publishers
│   ├── demo/                  # SQLite demo database
│   ├── ssebroker/             # Topic-based pub/sub for SSE
│   └── domain/                # Data models
├── web/
│   ├── views/                 # Page-level templ components
│   ├── components/core/       # Reusable UI components
│   ├── styles/                # Tailwind CSS input
│   └── assets/public/         # Static assets (embedded in binary)
│       ├── js/                # HTMX, Hyperscript, SSE extension
│       └── css/               # Tailwind, DaisyUI
└── tests/                     # Test utilities
```

## Template Setup

This repo doubles as a template for new Go + HTMX projects. Run `mage setup` to customize the module path, ports, and features for your own app.

### Interactive Setup (with gum)

Install [`gum`](https://github.com/charmbracelet/gum) for the interactive wizard:

```bash
go install github.com/charmbracelet/gum@latest
go tool mage setup
```

The wizard walks you through:

1. **Copy to new directory** -- Optionally copy the template to a new location and run `git init`
2. **App name** -- Human-readable name (e.g. "My App"), used to derive the binary name
3. **Module path** -- Go module path (e.g. `github.com/you/my-app`)
4. **Base port** -- 5-digit port number; the app uses `BASE_PORT`, templ proxy uses `BASE_PORT+1`, Caddy uses `BASE_PORT+2`
5. **Feature selection** -- Multi-select which features to include:

| Feature              | Description                                  | Default    |
| -------------------- | -------------------------------------------- | ---------- |
| Auth (Crooner)       | Azure AD authentication via Crooner          | Selected   |
| Graph API            | Microsoft Graph SDK integration              | Selected   |
| Avatar Photos        | User photo sync from Azure (requires Graph)  | Selected   |
| Database (MSSQL)     | SQL Server database with SQLx                | Selected   |
| SSE                  | Server-Sent Events real-time updates (requires Caddy) | Selected |
| Caddy (HTTPS)        | Caddy reverse proxy with TLS termination     | Selected   |
| Demo Content         | SQLite demo tables, hypermedia examples       | Unselected |

Deselected features have their code, routes, imports, and related files stripped from the project. Dependencies are auto-resolved (SSE includes Caddy, Avatar includes Graph).

### Non-interactive Setup

```bash
go tool mage setup -n "My App" -m "github.com/me/my-app" -p 12345
go tool mage setup -n "My App" --features sse,demo
go tool mage setup -n "My App" --features none
go tool mage setup -n "My App" --features all
```

| Flag           | Description                                                        |
| -------------- | ------------------------------------------------------------------ |
| `-n APP_NAME`  | App name (required)                                                |
| `-m MODULE`    | Go module path                                                     |
| `-p PORT`      | 5-digit base port (< 60000)                                       |
| `--features`   | Comma-separated: `auth,graph,avatar,database,sse,caddy,demo`, `all`, or `none` |
| `--force`      | Re-run setup on an already customized project                      |

After setup, review `.env.development` and start the dev server with `go tool mage watch`.

## Development

### Prerequisites

- Go 1.26+ (latest)
- Node.js 20+ (for Tailwind CSS compilation)
- (Optional) [`gum`](https://github.com/charmbracelet/gum) for interactive setup

### Running the Dev Server

```bash
# Install npm dependencies (first time)
npm ci

# Start development with live reload (Tailwind, Templ, Air)
go tool mage watch
```

The dev server starts with TLS on the configured port. Edit `.env.development` to change settings.

### HTTPS Development Setup

The dev server uses TLS with self-signed certificates. Generate them with:

```bash
openssl req -x509 -newkey rsa:2048 -keyout localhost.key -out localhost.crt \
  -days 365 -nodes -subj "/CN=localhost" \
  -addext "subjectAltName=DNS:localhost,IP:127.0.0.1"
```

Trust the certificate on your system:

- **Linux**: `sudo cp localhost.crt /usr/local/share/ca-certificates/ && sudo update-ca-certificates`
- **macOS**: Open Keychain Access, drag cert to System, set Trust to Always Trust
- **Windows**: Right-click cert, Install Certificate, Local Machine, Trusted Root CAs

## Testing

```bash
go tool mage test              # Run all tests
go tool mage testverbose       # Verbose output
go tool mage testcoverage      # Coverage report
go tool mage testrace          # Race detection
go tool mage testwatch         # Watch mode
```

## Mage Targets

All commands run with `go tool mage <target>`.

| Command             | Category     | Description                                                   |
| ------------------- | ------------ | ------------------------------------------------------------- |
| `watch`             | Development  | Start dev mode with live reload (Tailwind, Templ, Air)        |
| `air`               | Development  | Run Air live reload for Go                                    |
| `templ`             | Development  | Run Templ in watch mode                                       |
| `templgenerate`     | Development  | Generate Templ files once                                     |
| `build`             | Build        | Clean, update assets, and build the project                   |
| `compile`           | Build        | Build the Go binary                                           |
| `run`               | Build        | Build and execute                                             |
| `updateassets`      | Assets       | Update all assets (Hyperscript, HTMX, DaisyUI, Tailwind)     |
| `tailwind`          | Assets       | Run Tailwind CSS compilation                                  |
| `tailwindwatch`     | Assets       | Run Tailwind in watch mode                                    |
| `daisyupdate`       | Assets       | Update DaisyUI CSS                                            |
| `htmxupdate`        | Assets       | Update HTMX files                                             |
| `test`              | Testing      | Run all tests                                                 |
| `testverbose`       | Testing      | Tests with verbose output                                     |
| `testcoverage`      | Testing      | Tests with coverage report                                    |
| `testcoveragehtml`  | Testing      | Generate HTML coverage report                                 |
| `testbenchmark`     | Testing      | Run benchmark tests                                           |
| `testrace`          | Testing      | Tests with race detection                                     |
| `testwatch`         | Testing      | Tests in watch mode                                           |
| `lint`              | Code Quality | Run static analysis (golangci-lint, golint, fieldalignment)   |
| `fixfieldalignment` | Code Quality | Auto-fix field alignment                                      |
| `clean`             | Utility      | Remove build and debug files                                  |
| `caddyinstall`      | HTTPS        | Install Caddy for local dev                                   |
| `caddystart`        | HTTPS        | Start Caddy with TLS termination                              |
