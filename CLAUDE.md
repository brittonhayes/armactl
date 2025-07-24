# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

`armactl` is a cross-platform CLI tool for managing ARMA 3 servers, written in Go. It provides commands for mod management, server configuration, Steam server querying, and mod key management. The project also includes support for ARMA Reforger dedicated servers.

- This project is designed for game server administrators running arma servers for small communities

## Architecture

The codebase follows a modular CLI architecture using the Kong library for command parsing:

- **Main Entry Point**: `cmd/armactl/main.go` - Sets up CLI with Kong parser and zerolog
- **Commands**: `cmd/armactl/commands/` - Individual command implementations (mods, keys, server, steam, reforger, version)
- **Core Modules**:
  - `mods/` - Mod listing, syncing, and preset management
  - `server/` - ARMA 3 dedicated server configuration and startup
  - `steam/` - Steam server querying functionality
  - `reforger/` - ARMA Reforger config validation with JSON schema
  - `preset/` - HTML preset file parsing
- **Utilities**:
  - `internal/fs/` - Filesystem utilities
  - `internal/rclone/` - File syncing using rclone
  - `templates/` - Configuration file templates

The CLI uses a command context pattern (`commands.Context`) to pass shared state (logger, version, output) between commands.

## Development Commands

The project uses [Task](https://taskfile.dev/) for build automation. Key commands:

```bash
# Build the application
task build

# Run tests
task test

# Run tests with verbose output
task test:v

# Run tests with coverage
task test:cov

# Lint code
task lint

# Clean build artifacts
task clean

# Download and tidy dependencies
task mod

# Development setup
task dev-setup

# Generate documentation
task docs:gen

# Serve local documentation
task docs:serve

# Initialize React UI (first time setup)
task ui:init

# Start React dev server (for development)
task ui:dev

# Build React UI for production
task ui:build
```

## Web Interface

The project includes a modern React-based web interface that provides a GUI for all CLI functionality:

- **Start the web interface**: `armactl ui --port 3000`
- **Development mode**: `armactl ui --dev --port 3000` (enables CORS for hot reload)
- **Architecture**: RESTful API with React frontend, embedded as static assets in the binary

### Web Architecture

- **Backend**: Go REST API using chi router with hexagonal architecture
- **Frontend**: React + TypeScript + Vite + shadcn/ui + Tailwind CSS  
- **API Endpoints**: 
  - `/api/health` - Health check
  - `/api/server/*` - Server management
  - `/api/mods/*` - Mod management
  - `/api/steam/*` - Steam server querying
  - `/api/keys/*` - Key management

## Key Features

- **Mod Management**: List, sync, and validate mods against HTML presets
- **Server Management**: Start ARMA 3 dedicated servers with custom configurations
- **Key Management**: Handle `.bikey` files for mod verification
- **Steam Integration**: Query Steam servers for status information
- **Reforger Support**: Validate ARMA Reforger server configurations against JSON schema
- **Web Interface**: Modern React-based web UI for managing servers via `armactl ui`
- **Cross-platform**: Supports Windows, Linux, and macOS

## Testing

Tests are located alongside source files (e.g., `preset_test.go`, `schema_test.go`). The project uses standard Go testing with testify for assertions.

## Configuration

The tool supports both command-line flags and environment variables for configuration. Server templates are stored in `templates/` directory for generating configuration files.

## Project Standards

- This project uses idiomatic golang project structure and code standards

## Project Requirements

- This project should always be portable and downloadable as a single binary