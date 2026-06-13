## Project Summary

`playwithgo` is a Go REST API boilerplate built with the [go-blueprint](https://github.com/Melkeydev/go-blueprint) template. It wires together a Gin HTTP server, a MySQL database, and a clean layered structure ready for real application logic.

### Architecture

```
main.go
  └── server.NewServer()          ← reads PORT env, creates DB connection
        ├── database.New()         ← singleton MySQL pool via godotenv credentials
        └── s.RegisterRoutes()     ← Gin router with CORS + 2 routes
              ├── GET /            → HelloWorldHandler
              └── GET /health      → db.Health() stats
```

### File Overview

| File | Purpose |
|---|---|
| `cmd/api/main.go` | Entry point. Starts the HTTP server and handles graceful shutdown on `SIGINT`/`SIGTERM` (5-second drain window). |
| `internal/server/server.go` | Server factory. Reads `PORT` from env, initializes the DB, and returns a configured `*http.Server` with sane timeouts. |
| `internal/server/routes.go` | HTTP routes. Registers `GET /` (hello world) and `GET /health` (DB health check) via Gin, plus CORS middleware for `localhost:5173`. |
| `internal/server/routes_test.go` | Route unit test. Tests `HelloWorldHandler` returns `200 OK` with `{"message":"Hello World"}`. |
| `internal/database/database.go` | Database layer. Defines the `Service` interface, implements it with a singleton MySQL connection pool (max 50 conns), and exposes detailed `Health()` stats. |
| `internal/database/database_test.go` | DB integration test. Spins up a real MySQL 8 container via Testcontainers to test `New()`, `Health()`, and `Close()`. |
| `go.mod` / `go.sum` | Module definition. Key deps: `gin`, `gin-contrib/cors`, `go-sql-driver/mysql`, `godotenv`, `testcontainers-go`. |
| `.env` | Local config. Server port (`8080`) and MySQL connection details loaded automatically by `godotenv`. |
| `docker-compose.yml` | Local DB. Runs a MySQL container using `.env` credentials, persisting data in a named volume. |
| `.air.toml` | Live reload config. Configures [Air](https://github.com/air-verse/air) to watch Go/template files and rebuild via `make build` on changes. |
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
