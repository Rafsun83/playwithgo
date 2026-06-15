## Project Summary

`playwithgo` is a Go REST API boilerplate built with the [go-blueprint](https://github.com/Melkeydev/go-blueprint) template. It wires together a Gin HTTP server, a MySQL database, and a clean layered structure ready for real application logic.

---

## Architecture

### Directory Structure

```
playwithgo/
├── cmd/
│   └── server/
│       └── main.go           # Entry point: Server struct, NewServer(), graceful shutdown
├── internal/
│   ├── handlers/             # HTTP layer: route registration + handler implementations
│   │   ├── health.go         # Handler struct, HelloWorldHandler, HealthHandler
│   │   ├── routes.go         # RegisterRoutes() — Gin router + CORS middleware
│   │   └── routes_test.go    # Unit tests for handlers
│   ├── services/             # Business logic layer (to be populated)
│   ├── repository/           # Data access layer
│   │   ├── database.go       # Service interface, MySQL connection pool, Health()
│   │   └── database_test.go  # Integration tests via Testcontainers
│   └── models/               # Shared data structures (to be populated)
├── pkg/                      # Public reusable packages (optional)
├── config/                   # Configuration files
├── migrations/               # SQL migration files
├── .env                      # Local environment variables (not committed)
├── docker-compose.yml        # MySQL container for local development
├── Makefile                  # Dev workflow shortcuts
├── go.mod
└── go.sum
```

### Layer Diagram

```
  HTTP Request
       │
       ▼
┌─────────────────────────────────────┐
│         cmd/server/main.go          │
│  NewServer()                        │
│  ├── Reads PORT from env            │
│  ├── Calls repository.New()         │  ← DB connection pool (singleton)
│  └── Calls handlers.RegisterRoutes()│  ← Wires up Gin router
│  gracefulShutdown()                 │  ← SIGINT/SIGTERM → 5s drain
└─────────────────┬───────────────────┘
                  │
                  ▼
┌─────────────────────────────────────┐
│        internal/handlers/           │
│  RegisterRoutes(db Service)         │
│  ├── Gin router + CORS middleware   │
│  ├── GET /       → HelloWorldHandler│
│  └── GET /health → HealthHandler    │
└─────────────────┬───────────────────┘
                  │
                  ▼
┌─────────────────────────────────────┐
│  internal/services/  (future)       │
│  Business logic between handlers    │
│  and repository                     │
└─────────────────┬───────────────────┘
                  │
                  ▼
┌─────────────────────────────────────┐
│       internal/repository/          │
│  Service interface                  │
│  ├── New()      → singleton *sql.DB │
│  ├── Health()   → connection stats  │
│  └── Close()    → cleanup on exit   │
└─────────────────┬───────────────────┘
                  │
                  ▼
             MySQL (via
          docker-compose)
```

### Request Flow

```
Client
  │  GET /health
  ▼
Gin router (RegisterRoutes)
  │  CORS middleware applied
  ▼
HealthHandler (handlers/health.go)
  │  calls h.DB.Health()
  ▼
repository.Service.Health() (repository/database.go)
  │  PingContext (1s timeout) + sql.DBStats
  ▼
JSON response  {"status":"up", "open_connections":"3", ...}
```

### Key Components

| Component | Location | Responsibility |
|---|---|---|
| `Server` struct | `cmd/server/main.go` | Holds `port` and `db`; wires everything together in `NewServer()` |
| `gracefulShutdown()` | `cmd/server/main.go` | Listens for `SIGINT`/`SIGTERM`, drains in-flight requests within 5 seconds |
| `RegisterRoutes()` | `internal/handlers/routes.go` | Creates the Gin engine, applies CORS, injects `Handler`, registers all routes |
| `Handler` struct | `internal/handlers/health.go` | Holds `DB repository.Service`; implements all HTTP handler methods |
| `repository.Service` | `internal/repository/database.go` | Interface with `Health()` and `Close()`; concrete impl uses a singleton `*sql.DB` with max 50 open/idle connections |

### Environment Variables

| Variable | Used by | Description |
|---|---|---|
| `PORT` | `cmd/server/main.go` | HTTP listen port (e.g. `8080`) |
| `BLUEPRINT_DB_HOST` | `internal/repository/database.go` | MySQL host |
| `BLUEPRINT_DB_PORT` | `internal/repository/database.go` | MySQL port |
| `BLUEPRINT_DB_DATABASE` | `internal/repository/database.go` | Database name |
| `BLUEPRINT_DB_USERNAME` | `internal/repository/database.go` | MySQL user |
| `BLUEPRINT_DB_PASSWORD` | `internal/repository/database.go` | MySQL password |

---

### File Overview

| File | Purpose |
|---|---|
| `cmd/server/main.go` | Entry point. Bootstraps the HTTP server and handles graceful shutdown on `SIGINT`/`SIGTERM` (5-second drain window). |
| `internal/handlers/routes.go` | `RegisterRoutes(db)` creates the Gin router, applies CORS middleware for `localhost:5173`, and wires up all routes. |
| `internal/handlers/health.go` | `Handler` struct holds the DB reference; implements `HelloWorldHandler` and `HealthHandler`. |
| `internal/handlers/routes_test.go` | Unit test. Verifies `HelloWorldHandler` returns `200 OK` with `{"message":"Hello World"}`. |
| `internal/repository/database.go` | Defines the `Service` interface, implements a singleton MySQL connection pool (max 50 conns), and exposes detailed `Health()` stats. |
| `internal/repository/database_test.go` | Integration test. Spins up a real MySQL 8 container via Testcontainers to test `New()`, `Health()`, and `Close()`. |
| `internal/services/` | Business logic layer (to be populated). |
| `internal/models/` | Shared data structures (to be populated). |
| `go.mod` / `go.sum` | Module definition. Key deps: `gin`, `gin-contrib/cors`, `go-sql-driver/mysql`, `godotenv`, `testcontainers-go`. |
| `.env` | Local config. Server port and MySQL credentials, loaded automatically by `godotenv`. |
| `docker-compose.yml` | Runs a MySQL container using `.env` credentials, persisting data in a named volume. |
| `Makefile` | Dev workflow. Shortcuts for build, run, test, integration tests, live reload, and Docker management. |

---

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

Run build make command with tests

```bash
make all
```

Build the application

```bash
make build
```

Run the application

```bash
make run
```

Create DB container

```bash
make docker-run
```

Shutdown DB Container

```bash
make docker-down
```

DB Integrations Test:

```bash
make itest
```

Live reload the application:

```bash
make watch
```

Run the test suite:

```bash
make test
```

Clean up binary from the last build:

```bash
make clean
```
