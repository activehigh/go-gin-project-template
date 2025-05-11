# Usage Guide

## Running the Project

### Locally
```sh
go run ./cmd/main.go
```

### With Docker
```sh
docker-compose up --build
```

The server will start on `http://localhost:8080` by default.

## API Endpoints

### Healthchecks
- **GET /live** — Liveness probe
- **GET /ready** — Readiness probe

#### Example Response
```json
{
  "message": "I am alive!"
}
```

### Ping
- **GET /ping** — Returns a pong message
- **POST /ping** — Returns a pong message

#### Example Response
```json
{
  "message": "pong"
}
```

## Environment Variables
- `TERMINATION_GRACE_PERIOD_IN_SECONDS` — Graceful shutdown period (default: 5)

## Logs
Structured logs are output using zap. See the console for logs when running locally or in Docker.
