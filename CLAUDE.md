# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

`armactl` is a cross-platform CLI tool for managing ARMA 3 and ARMA Reforger dedicated servers. It supports Windows, Linux, and macOS.

**Requirements:** Go 1.18 or higher

**Key capabilities:**
- Start ARMA 3 dedicated servers with headless client support
- Manage mods (list, sync, compare against presets)
- Parse and compare ARMA 3 HTML preset files
- Query Steam servers via A2S protocol
- Validate ARMA Reforger server configurations against JSON schema
- Manage bikey files for mod signatures

## Build and Development Commands

This project uses [Task](https://taskfile.dev) for build automation. View all available commands with `task -l`.

### Common commands:
```bash
# Development setup
task dev-setup              # Install dependencies and build

# Building
task build                  # Build binary to bin/armactl
task install               # Install via go install

# Testing
task test                  # Run all tests
task test:v                # Run tests verbose
task test:cov              # Run tests with coverage report

# Code quality
task lint                  # Run golangci-lint

# Dependencies
task mod                   # Download and tidy go modules

# Cleanup
task clean                 # Remove binaries and caches

# Documentation
task docs:gen              # Generate docs via gomarkdown
task docs:serve            # Serve godoc at localhost:6060
```

### Testing individual packages:
```bash
go test ./preset -v              # Test preset HTML parsing
go test ./reforger -v            # Test Reforger config validation

# Run tests with race detection (as done in CI)
go test -race ./...

# View coverage details
go tool cover -html=coverage.out
```

### Release process:
This project uses GoReleaser for releases. Releases are automated via GitHub Actions and publish to:
- GitHub Releases (binaries)
- Homebrew tap ([brittonhayes/armactl](https://github.com/brittonhayes/armactl))
- Scoop bucket ([brittonhayes/armactl](https://github.com/brittonhayes/armactl))
- Docker image (ghcr.io/brittonhayes/armactl)

## Architecture Overview

### Command Structure
The CLI uses [Kong](https://github.com/alecthomas/kong) for command parsing with a clean separation between command definitions and domain logic:

1. **Entry point**: [cmd/armactl/main.go](cmd/armactl/main.go) - Initializes Kong parser, sets up zerolog, and dispatches to commands
2. **Command layer**: [cmd/armactl/commands/](cmd/armactl/commands/) - Kong command structs that handle CLI flags and delegate to domain packages
3. **Domain layer**: Top-level packages ([mods/](mods/), [server/](server/), [preset/](preset/), [steam/](steam/), [reforger/](reforger/)) - Business logic
4. **Internal utilities**: [internal/](internal/) - Shared utilities (rclone wrapper for file operations)

### Context Pattern
All commands receive a `commands.Context` struct containing:
- `context.Context` - For cancellation and timeouts
- `zerolog.Logger` - Structured logging (debug level when `--verbose`)
- `io.Writer` - Output destination (typically `os.Stdout`)
- `Version` string

### Domain Package Responsibilities

**[server/](server/)** - ARMA 3 server management
- Constructs command-line arguments for `arma3server_x64` binary
- Spawns dedicated server process
- Manages headless client processes (spawned concurrently with sync.WaitGroup)
- Server interface defines Start/Stop/StartHeadless/StopHeadless operations

**[mods/](mods/)** - Mod management
- `FromDirectory()` - Scans directories for mods by reading `meta.cpp` files
- `Meta()` - Parses mod names from `meta.cpp` using regex
- `NewSync()` - Wraps rclone for syncing mod directories
- `keys.go` - Bikey file management

**[preset/](preset/)** - HTML preset parsing
- Uses [goquery](https://github.com/PuerkitoBio/goquery) to parse ARMA 3 Launcher HTML exports
- Extracts mod list from `[data-type='ModContainer']` elements
- Returns sorted mod names using natural string sorting

**[steam/](steam/)** - Steam server querying
- Uses [go-a2s](https://github.com/rumblefrog/go-a2s) for A2S protocol
- `Inspect()` - Query server info (name, game, player count)
- `Players()` - List connected players

**[reforger/](reforger/)** - ARMA Reforger support
- Validates server config JSON files against embedded JSON schema
- Schema embedded via `//go:embed config.schema.json`
- Uses [jsonschema](https://github.com/santhosh-tekuri/jsonschema) for validation

**[internal/rclone/](internal/rclone/)** - File operations wrapper
- Wraps [rclone library](https://github.com/rclone/rclone) for file operations
- `Sync()` - Synchronize directories (used for mod syncing)
- `Copy()` - Copy individual files
- `Find()` - Recursively find files by extension (used for bikey discovery)

## Key Dependencies

- **Kong** - CLI framework with struct tags for flag definitions
- **zerolog** - Structured logging with console writer
- **rclone** - File syncing operations (imported as library, not CLI)
- **goquery** - HTML parsing for preset files
- **go-a2s** - Steam A2S protocol client
- **jsonschema** - JSON schema validation
- **natural** - Natural string sorting for mod lists

## Adding New Commands

1. Create command struct in [cmd/armactl/commands/](cmd/armactl/commands/) with Kong tags
2. Implement `Run(ctx *Context) error` method
3. Add command field to `cli` struct in [main.go](cmd/armactl/main.go)
4. Create domain package if needed for business logic
5. Add task to [Taskfile.yml](Taskfile.yml) if command needs build/test tasks

Example command structure:
```go
type NewCmd struct {
    Flag string `help:"Description" default:"value" short:"f"`
}

func (c *NewCmd) Run(ctx *Context) error {
    ctx.Log.Info().Msg("Starting command")
    // Delegate to domain package
    return domainpkg.DoWork(ctx.Ctx, c.Flag)
}
```

## Configuration

Most server options can be set via environment variables (see `env:` tags in [commands/server.go](cmd/armactl/commands/server.go)):
- `ARMA_BINARY` - Path to server binary
- `ARMA_MODS` - Semicolon-separated mod list
- `ARMA_CONFIG` - Server config file path
- `ARMA_PORT` - Server port
- `STEAM_USER`, `STEAM_PASSWORD` - SteamCMD credentials
- `HEADLESS_CLIENTS` - Number of headless clients to spawn

## Testing

Test coverage is currently minimal:
- [preset/preset_test.go](preset/preset_test.go) - Tests HTML preset parsing
- [reforger/schema_test.go](reforger/schema_test.go) - Tests config validation

When adding features, consider adding tests for:
- Domain logic (parsing, validation, transformations)
- Error handling edge cases
- Cross-platform path handling

## Code Style

- Use functional options pattern for server configuration (see [server/options.go](server/options.go))
- Define interfaces for testability (Parser, Syncer, Client, etc.)
- Use natural sorting for user-facing lists (mod names, etc.)
- Log at Info level for user actions, Debug level for implementation details
- Return errors, don't panic (except in initialization like steam.New)

## Important Implementation Notes

### Server Command Dry-Run Mode
The server command supports `--dry-run` mode which prints the commands that would be executed without actually starting the server. This is useful for debugging server configuration issues. The server builds command arguments using the functional options pattern in [server/options.go](server/options.go).

### Mod Discovery
Mods are discovered by scanning directories for `meta.cpp` files. The parser uses regex to extract mod names from the C++ style configuration format. If mods aren't being discovered, check that they contain a valid `meta.cpp` file with a `name` field.

### Headless Client Management
Headless clients are spawned concurrently using `sync.WaitGroup` when the server starts. Each headless client connects back to the server address specified by `--headless-clients-server` (defaults to `127.0.0.1`). The number of headless clients is controlled by `--headless-clients` flag.

### Error Handling Patterns
Domain packages define custom error types (e.g., `ErrNoKeysFound`, `ErrNoModsFound`, `ErrModMismatch`) for common failure scenarios. These can be checked with `errors.Is()` for specific error handling in commands.
