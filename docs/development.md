# Development Guide

## Prerequisites
- [Go 1.24+](https://golang.org/)
- [Docker](https://www.docker.com/) (optional)

## Setup
1. Clone the repository:
   ```sh
   git clone https://github.com/activehigh/go-gin-project-template.git
   cd go-gin-project-template
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Copy and edit environment variables as needed:
   ```sh
   cp private.env .env
   # Edit .env as needed
   ```

## Running the Server
- Locally: `go run ./cmd/main.go`
- With Docker: `docker-compose up --build`

## Development Workflow
- Follow Go best practices for code structure and formatting.
- Use feature branches for new features or bug fixes.
- Write tests for new code.
- Run tests with:
  ```sh
  go test ./...
  ```
- Use `justfile` for common tasks (if [just](https://github.com/casey/just) is installed):
  ```sh
  just
  ```

## Linting & Formatting
- Use `gofmt` and `golangci-lint` for code quality.

## Testing
- Unit and integration tests are located in the `internal/` and `tests/` directories.
- Run all tests:
  ```sh
  go test ./...
  ```

## Contributing
- Fork the repo and create your branch from `main`.
- Ensure your code passes all tests and lints.
- Open a pull request with a clear description of your changes.
