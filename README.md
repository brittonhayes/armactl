# armactl

[![Go Reference](https://pkg.go.dev/badge/github.com/brittonhayes/armactl.svg)](https://pkg.go.dev/github.com/brittonhayes/armactl)
![Latest Release](https://img.shields.io/github/v/release/brittonhayes/armactl?label=latest%20release)
[![Go Report Card](https://goreportcard.com/badge/github.com/brittonhayes/armactl)](https://goreportcard.com/report/github.com/brittonhayes/armactl)

`armactl` is a cross-platform CLI tool for managing ARMA 3 servers. Supports Windows, Linux, and MacOS.

## Installation

### Go Install

```bash
go install github.com/brittonhayes/armactl/cmd/armactl@latest
```

### Homebrew (Linux)

```bash
brew tap brittonhayes/armactl
brew install armactl
```

### Docker Image

```
docker run --rm -it ghcr.io/brittonhayes/armactl:latest --help
```

### Manual

Download the latest release from the [releases page](https://github.com/brittonhayes/armactl/releases).

## Usage

To see all available commands, run `armactl --help`.

### Mod Management

To see all available commands, run `armactl mod --help`.

List all mods in a directory, a preset HTML file, or check that all mods in a directory match a preset HTML file.


```bash
# List all mods in a directory
armactl mods ls -d /path/to/mods

# List all mods in preset HTML file
armactl mods ls -p /path/to/preset.html

# Check all mods in directory match preset HTML file
armactl mods ls -p /path/to/preset.html -d /path/to/mods --check
```

Sync your mods between directories

```bash
# Sync mods from one directory to another
armactl mods sync --from '/path/to/source/@CBA_A3' --to '/path/to/destination/@CBA_A3'

# Sync multiple comma-separated mods from one directory to another
armactl mods sync --from '/path/to/@CBA_A3,/path/to/@ace' --to '/path/to/destination/@CBA_A3,/path/to/destination/@ace'
```