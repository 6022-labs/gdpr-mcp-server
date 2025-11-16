# AGENTS.md

## Project Overview

- Purpose: GDPR Model Context Protocol (MCP) server and data library to expose structured GDPR content (recitals, chapters, articles, and article paragraphs) for AI coding agents and MCP-aware clients.
- Structure:
  - `data/v1/`: Canonical GDPR data in JSON, organized into `recitals/`, `chapters/`, and `articles/<art-n>/` folders.
  - `src/gdpr_mcp_server`: Go library with core models for GDPR artifacts.
  - `src/gdpr_mcp_server_host`: Intended host binary for MCP server (entrypoint). Currently scaffolding; `main.go` is empty.
  - `src/gdpr_mcp_server_tools`: Reserved for internal tools (e.g., generators). Currently empty.
- Technologies: Go 1.24+, Uber Dig (DI), Uber Zap (logging).
- Intended runtime: An MCP server executable speaking stdio to MCP clients.

## Setup Commands

- Prerequisites:
  - Go 1.24+ installed and on PATH
  - macOS zsh shell (repo author’s default)
- Clone and prepare modules:

```bash
git clone https://github.com/6022-labs/gdpr-mcp-server.git
cd gdpr-mcp-server
go mod download
```

- Optional developer tools:

```bash
# Format and imports
go install golang.org/x/tools/cmd/goimports@latest

# Linters (choose one)
go install honnef.co/go/tools/cmd/staticcheck@latest
# or
brew install golangci-lint # if using Homebrew
```

## Development Workflow

- Build all packages (no binaries produced yet):

```bash
go build ./...
```

- Library development (`src/gdpr_mcp_server`):

  - Add functionality under this module and unit tests alongside using `_test.go` files.

- Host binary (`src/gdpr_mcp_server_host`):
  - Goal: implement an MCP stdio server. The current `main.go` is empty; create an entrypoint before running.
  - Example minimal stub to unblock local runs:

```go
// file: src/gdpr_mcp_server_host/main.go
package main

import "fmt"

func main() { fmt.Println("gdpr-mcp-server host stub") }
```

- After a minimal `main.go` exists, you can run/build:

```bash
# Run from repo root
go run ./src/gdpr_mcp_server_host

# Build a binary
go build -o bin/gdpr-mcp ./src/gdpr_mcp_server_host
```

- Data usage:
  - GDPR JSON data lives under `data/v1`. Prefer relative paths from repo root or inject a data root via an env var (e.g., `GDPR_DATA_ROOT`), then resolve files under `recitals/`, `chapters/`, and `articles/`.

## Testing Instructions

- Run all tests:

```bash
go test ./...
```

- With race detector and verbose output:

```bash
go test -race -v ./...
```

- Coverage report (HTML):

```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

- Test layout:
  - Place `_test.go` files next to the code being tested.
  - Use table-driven tests and keep fixtures small and deterministic.
  - Add/adjust tests with every code change.

## Code Style

- Formatting and imports:

```bash
gofmt -s -w .
# optional, ensures imports are grouped and ordered
goimports -w .
```

- Static analysis:

```bash
go vet ./...
staticcheck ./...        # if installed
# or
golangci-lint run         # if using golangci-lint
```

- Conventions:
  - Package names: short, all-lowercase, no underscores (except existing paths).
  - Filenames: `snake_case.go` as per Go norm; tests end with `_test.go`.
  - Errors: return wrapped errors with context where helpful.
  - Logging: prefer structured logging via Zap when host/server code is implemented.

## Build and Deployment

- Build library packages only:

```bash
go build ./src/gdpr_mcp_server/...
```

- Build host binary (after `main.go` is implemented):

```bash
mkdir -p bin
go build -o bin/gdpr-mcp ./src/gdpr_mcp_server_host
```

- Docker (note: current Dockerfile references a different project and should be updated):
  - The file `src/gdpr_mcp_server_host/Dockerfile` is a placeholder. Update it to target this repo/module. Example multi-stage Dockerfile:

```dockerfile
# file: src/gdpr_mcp_server_host/Dockerfile
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /out/gdpr-mcp ./src/gdpr_mcp_server_host

FROM gcr.io/distroless/base-debian12
COPY --from=builder /out/gdpr-mcp /usr/local/bin/gdpr-mcp
USER nonroot:nonroot
ENTRYPOINT ["/usr/local/bin/gdpr-mcp"]
```

- Build and run container (after Dockerfile is corrected):

```bash
# build from repo root
docker build -f src/gdpr_mcp_server_host/Dockerfile -t 6022labs/gdpr-mcp-server:dev .
docker run --rm -it 6022labs/gdpr-mcp-server:dev
```

- MCP client integration (stdio):
  - Point your MCP client to the built binary (`bin/gdpr-mcp`) and use stdio transport. Pass environment like `GDPR_DATA_ROOT` if introduced.

## Security Considerations

- Secrets: none required; do not commit secrets. Use environment variables if future integrations need tokens.
- Input handling: validate and sanitize any user-provided selectors or search parameters once request handling is added.
- Logging: avoid logging PII; redact or omit request details that may contain personal data.
- Supply chain: pin dependencies in `go.mod` and periodically update.

## Pull Request Guidelines

- Branch naming: `feat/...`, `fix/...`, `chore/...`, `docs/...` (use your team’s convention consistently).
- Title format: `[gdpr-mcp] Brief description`.
- Required local checks before pushing:

```bash
gofmt -s -w . && go vet ./... && go test ./...
```

- Add or update tests for any behavior change. Keep coverage trending upward where practical.
- Keep PRs focused and small; include a brief rationale in the description.

## Debugging and Troubleshooting

- Build succeeds but nothing runs:

  - The host `main.go` is currently empty. Implement a minimal entrypoint as shown above, then `go run ./src/gdpr_mcp_server_host`.

- Docker build issues:

  - Ensure the Dockerfile is updated to this repo’s paths (not the placeholder in the current file).

- Data not found at runtime:
  - Confirm working directory is repo root or that `GDPR_DATA_ROOT` points at `data/v1`.

## Notes for Agents

- Prefer reading and writing within `src/gdpr_mcp_server` for reusable logic. Keep `src/gdpr_mcp_server_host` thin (wiring, config, DI, and process lifecycle).
- When adding MCP methods, structure protocol-facing code in the host, and domain/data code in the library package.
- If adding generators/tools, place them under `src/gdpr_mcp_server_tools` and document usage in this file.
