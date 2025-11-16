# GitHub Copilot Instructions

## Priority Guidelines

When generating code for this repository:

1. Version Compatibility: Detect and respect the exact versions used here
2. Context Files: Prioritize this file (`.github/copilot/copilot-instructions.md`)
3. Codebase Patterns: Mirror patterns found in existing source files
4. Architectural Consistency: Maintain our Hexagonal architecture boundaries
5. Code Quality: Prioritize maintainability, performance, security, and testability

## Technology Version Detection

From the codebase (do not assume beyond these):

- Language: Go 1.24.3 (from `go.mod` and Dockerfile base)
- Dependencies (from `go.mod`):
  - `go.uber.org/dig`
  - `go.uber.org/zap`
- HTTP framework: `net/http`.

Rules:

- Use only Go features compatible with `go 1.24.3`.
- Use only APIs available in the listed dependency versions.
- Do not reference unlisted libraries (e.g., Fiber) unless they are first added to `go.mod`.

## Context Files

This repository currently provides this instructions file as the source of truth.

## Codebase Scanning Instructions

When generating or modifying code:

1. Identify similar files and mirror their patterns.
2. Focus areas to analyze:
   - Naming conventions in `src/gdpr_mcp_server/models/*`
   - DI patterns in `src/gdpr_mcp_server_host/configurations/di_configuration.go`
   - Runtime settings and env handling in `src/gdpr_mcp_server_host/settings/host_settings.go`
3. Follow the most consistent and recent patterns; avoid introducing new patterns.
4. Do not add frameworks or libraries that are not already present in `go.mod`.

## Architecture

Adopt and preserve Hexagonal architecture boundaries evident in the structure:

- Host/Composition (`src/gdpr_mcp_server_host`)
  - Entry point and runtime composition (DI, logging, settings)
  - Files: `configurations/*`, `settings/*`, `main.go`
- Domain (`src/gdpr_mcp_server`)
  - Domain models and application/core logic
  - Files: `models/*`, `services/*`, `use_cases/*`
- Data
  - Static GDPR data in `data/v1/*` for articles, chapters, recitals
- Primary Adapters
  - Include input adapters (e.g., HTTP handlers) here when added
- Secondary Adapters
  - Include output adapters (e.g., database repositories, access to file in the systems, access to external services) here when added

Guidelines:

- Keep dependency flow :
  - From host → domain; domain should not import host.
  - From primary adapters → domain; primary adapters can use domain interfaces/implementations.
  - From domain → secondary adapters; domain defines interfaces, secondary adapters implement them.
- Inject dependencies via DI (dig) from the host layer.

## Code Patterns

### Naming and Organization

- Packages: lower_snake for directory names; concise package names (`models`, `settings`, `configurations`).
- Types: `PascalCase` structs with JSON tags as `camelCase`.
- Files group by responsibility (e.g., `logging_configuration.go`, `di_configuration.go`).

### Dependency Injection (dig)

- Container creation and provisioning follows:

```go
// src/gdpr_mcp_server_host/configurations/di_configuration.go
func ConfigureDI() *dig.Container {
    container := dig.New()
    container.Provide(settings.NewHostSettings)
    return container
}
```

- Dependencies are injected via constructor params. Example:

```go
// src/gdpr_mcp_server_host/settings/host_settings.go
func NewHostSettings(logger *zap.Logger) *HostSettings { /* ... */ }
```

- When adding new services/components:
  - Define constructors that accept required dependencies as parameters.
  - Register constructors with `container.Provide`.
  - Use `container.Invoke` to compose and run at the host layer (match existing style when added).

### Logging (zap)

- Logging is configured centrally using a production config and env-driven level:

```go
// src/gdpr_mcp_server_host/configurations/logging_configuration.go
zapConfig := zap.NewProductionConfig()
// LOG_LEVEL supports: debug|info|warn|error|fatal|panic; defaults to info
logger, err := zapConfig.Build()
container.Provide(func() *zap.Logger { return logger })
```

- Error policy observed:

  - Configuration build failure: `panic(err)`
  - DI provision failure: `logger.Fatal("Failed to provide logger", zap.Error(err))`
  - Settings validation:
    - Missing `API_PORT`: warn and default to 3000
    - Invalid `API_PORT`: warn and default to 3000
    - Missing `APP_NAME`: `logger.Fatal(...)`

- When adding code that logs:
  - Accept `*zap.Logger` via DI and use structured fields (`zap.Error(err)`).
  - Do not use global loggers or other logging libraries.

### Configuration and Environment

- Environment variables used:

  - `LOG_LEVEL`: controls zap level (`debug|info|warn|error|fatal|panic`)
  - `API_PORT`: defaults to `3000` with warnings on missing/invalid
  - `APP_NAME`: required; missing value is fatal

- Pattern: read envs in constructors, validate, and log warnings/fatal consistently.

### HTTP and Transport

- `net/http` official package as transport.

## Code Quality Standards

### Maintainability

- Write self-documenting code with clear naming.
- Follow the existing file naming and package organization patterns.
- Keep constructors focused and validate inputs early with logging.

### Performance

- Prefer injection of shared dependencies (e.g., `*zap.Logger`) rather than creating new instances.
- Avoid unnecessary allocations in hot paths; mirror simplicity of existing models.

### Security

- Treat environment-derived configuration as untrusted: validate and log issues as shown in `NewHostSettings`.
- Avoid logging sensitive values; log only metadata and errors.

### Testability

- Constructors should accept interfaces and dependencies via params to enable substitution in tests.
- Do not introduce external testing libraries unless added to `go.mod`.

## Documentation Requirements

- Match current minimal style: brief, focused code with limited comments.
- Add comments only where behavior is non-obvious (e.g., env parsing rules, defaults, or failure policies).

## Go-Specific Guidelines

- Modules and versions:
  - Respect `module github.com/6022-labs/gdpr-mcp-server` and `go 1.24.3`.
  - Keep imports internal to this module unless explicitly required.
- Error handling:
  - Follow observed patterns: `panic` for irrecoverable configuration build errors; `logger.Fatal` for critical DI issues; warnings for recoverable env issues.
- JSON models:
  - Keep field names `PascalCase` with explicit JSON tags as `camelCase`.
  - Avoid embedding behavior in models; keep them as simple data structures.

## Project-Specific Guidance

- Scan the codebase before generating code; mirror patterns exactly.
- Respect and preserve boundaries between host composition and domain code.
- Use `dig` for DI wiring and `zap` for logging; do not introduce alternatives.
- Prefer small, focused files under appropriate folders (`configurations`, `settings`, `models`).

## Examples from the Codebase

- DI provisioning and settings constructor:

```go
// Provide settings in the container
container.Provide(settings.NewHostSettings)

// Settings constructor with logger injection and env handling
func NewHostSettings(logger *zap.Logger) *HostSettings {
    apiPortStr := os.Getenv("API_PORT")
    if len(strings.TrimSpace(apiPortStr)) == 0 { logger.Warn("environment variable API_PORT is not set, defaulting to 3000"); apiPortStr = "3000" }
    apiPort, err := strconv.Atoi(apiPortStr)
    if err != nil { logger.Warn("invalid API_PORT environment variable value, defaulting to 3000"); apiPort = 3000 }
    appName := os.Getenv("APP_NAME")
    if len(strings.TrimSpace(appName)) == 0 { logger.Fatal("please set your APP_NAME value in your environment") }
    return &HostSettings{ ApiPort: apiPort, AppName: appName }
}
```

- Logging configuration with level mapping:

```go
zapConfig := zap.NewProductionConfig()
switch os.Getenv("LOG_LEVEL") {
case "debug": zapConfig.Level.SetLevel(zap.DebugLevel)
case "info": zapConfig.Level.SetLevel(zap.InfoLevel)
case "warn": zapConfig.Level.SetLevel(zap.WarnLevel)
case "error": zapConfig.Level.SetLevel(zap.ErrorLevel)
case "fatal": zapConfig.Level.SetLevel(zap.FatalLevel)
case "panic": zapConfig.Level.SetLevel(zap.PanicLevel)
default: zapConfig.Level.SetLevel(zap.InfoLevel)
}
logger, err := zapConfig.Build()
if err != nil { panic(err) }
_ = container.Provide(func() *zap.Logger { return logger })
```

Follow these examples precisely when adding similar capabilities.
