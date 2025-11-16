# GDPR MCP Server

[![Unit Tests](https://github.com/6022-labs/gdpr-mcp-server/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/6022-labs/gdpr-mcp-server/actions/workflows/tests.yml)

GDPR Model Context Protocol (MCP) server exposing structured GDPR content (recitals, chapters, articles, and article paragraphs) to MCP-aware clients over HTTP.

### Available MCP Tools

- `GetArticleById(article_id)`
- `GetChapterById(chapter_id)`
- `GetRecitalById(recital_id)`
- `GetArticleParagraphsByArticleId(article_id, index)`

## Technology Stack

- Go 1.24.3
- Dependency Injection: `go.uber.org/dig`
- Logging: `go.uber.org/zap` (with optional `slog-zap` bridge)
- MCP SDK: `github.com/modelcontextprotocol/go-sdk`
- Testing: `testing`, `github.com/stretchr/testify`

## Quick Start

```zsh
git clone https://github.com/6022-labs/gdpr-mcp-server.git
cd gdpr-mcp-server
go mod download
```

Minimal env (from `.env.example`):

```zsh
export APP_NAME="gdpr-mcp-server"
export DAL_ARTICLES_DATA_FILE_PATH="$(pwd)/data/v1/articles"
export DAL_CHAPTERS_DATA_FILE_PATH="$(pwd)/data/v1/chapters"
export DAL_RECITALS_DATA_FILE_PATH="$(pwd)/data/v1/recitals"
# Optional: API_PORT=3000 LOG_LEVEL=info STATELESS=false JSON_RESPONSE=false
```

You can place these in a `.env` file to load it at startup.

### Run the server:

```zsh
go run ./src/gdpr_mcp_server_host
# listens on :${API_PORT:-3000}
```

### Build:

```zsh
mkdir -p bin && go build -o bin/gdpr-mcp ./src/gdpr_mcp_server_host
```

### Docker

Build the Docker image:

```zsh
docker build \
    -t gdpr-mcp-server:dev \
    -f src/gdpr_mcp_server_host/Dockerfile \
    -e APP_NAME="gdpr-mcp-server" \
    -e API_PORT=3000 \
    .
```

### Project Structure

- `src/gdpr_mcp_server_host`: composition, DI, logging, settings, MCP HTTP server
- `src/gdpr_mcp_server`: domain models + repository interfaces
- `src/gdpr_mcp_server_dal`: JSON-backed repositories
- `src/gdpr_mcp_server_tools`: MCP tool controllers
- `data/v1`: canonical GDPR JSON

## Testing

Run tests:

```zsh
go test ./...
```

## Docs

- Architecture and workflows: AGENTS.md
- Coding standards & generator guidance: .github/copilot/copilot-instructions.md

## Contributing

- Follow the architecture and patterns used in:
  - `src/gdpr_mcp_server_host/configurations/*` and `settings/*`
  - `src/gdpr_mcp_server/models/*` and `repositories/*`
  - `src/gdpr_mcp_server_tools/*` for MCP tool registration
- Add tests alongside your changes.
- Keep dependencies minimal and already declared in `go.mod` unless discussed.
- Keep your PRs focused and concise.
- Before submitting a PR:

```zsh
gofmt -s -w . && go vet ./... && go test ./...
```

## License

MIT - See [LICENSE](LICENSE) for details.
