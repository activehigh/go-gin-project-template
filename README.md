# go-gin-project-template

A production-ready template for building RESTful APIs with [Gin](https://github.com/gin-gonic/gin) in Go. This template provides a robust foundation for scalable, maintainable, and testable web services.

## Features
- Fast, minimal, and idiomatic Go web server using Gin
- Healthcheck endpoints (`/live`, `/ready`)
- Example ping endpoint (`/ping`)
- Structured logging with Uber's zap
- Graceful shutdown
- Modular project structure
- Docker and docker-compose support
- Ready for CI/CD and container orchestration

## Tech Stack
- [Go 1.24+](https://golang.org/)
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [zap Logger](https://github.com/uber-go/zap)
- [Docker](https://www.docker.com/)

## Quick Start

### Prerequisites
- Go 1.24 or newer
- Docker (optional, for containerized runs)

### Running Locally
```sh
git clone https://github.com/activehigh/go-gin-project-template.git
cd go-gin-project-template
go run ./cmd/main.go
```

### Using Docker
```sh
docker-compose up --build
```

## Usage
See [docs/usage.md](docs/usage.md) for API endpoints and examples.

## Development
See [docs/development.md](docs/development.md) for setup, environment, and workflow.

## Architecture
See [docs/architecture.md](docs/architecture.md) for project structure and design.

## Contributing
Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

## License
[MIT](LICENSE)