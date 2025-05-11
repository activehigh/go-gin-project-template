# Architecture Overview

## Project Structure

```
├── cmd/                # Application entry point (main.go)
├── configs/            # Configuration files and logic
├── internal/           # Application core logic
│   ├── handlers/       # HTTP handlers (business logic)
│   │   ├── healthcheck/  # Healthcheck endpoint
│   │   └── ping/         # Ping endpoint
│   ├── logger/         # Logging setup (zap)
│   └── server/         # Server setup, routing, connection management
├── tests/              # Test files
├── docs/               # Documentation
├── Dockerfile          # Docker build file
├── docker-compose.yaml # Docker Compose setup
├── go.mod, go.sum      # Go modules
└── README.md           # Project overview
```

## Main Components

- **cmd/main.go**: Application entry point. Sets up logging, loads config, initializes the server, and handles graceful shutdown.
- **internal/server/router.go**: Sets up Gin routes and binds handlers.
- **internal/handlers/**: Contains business logic for each endpoint (e.g., healthcheck, ping).
- **internal/logger/**: Initializes and configures zap logger.
- **configs/**: Handles configuration loading and CLI arguments.

## Routing

- `/live` and `/ready`: Healthcheck endpoints for liveness and readiness probes.
- `/ping`: Example endpoint for connectivity checks (GET/POST).

## Logging

- Uses Uber's zap for structured, high-performance logging.

## Graceful Shutdown

- Listens for SIGINT/SIGTERM and gracefully shuts down the server, allowing in-flight requests to complete.

## Extending the Template

- Add new endpoints by creating new handler packages in `internal/handlers/` and registering them in `internal/server/router.go`.
- Add configuration options in `configs/` as needed.
