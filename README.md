# armactl

[![Go Reference](https://pkg.go.dev/badge/github.com/brittonhayes/armactl.svg)](https://pkg.go.dev/github.com/brittonhayes/armactl)
![Latest Release](https://img.shields.io/github/v/release/brittonhayes/armactl?label=latest%20release)
[![Go Report Card](https://goreportcard.com/badge/github.com/brittonhayes/armactl)](https://goreportcard.com/report/github.com/brittonhayes/armactl)

`armactl` is a cross-platform CLI tool for managing ARMA 3 servers. Supports Windows, Linux, and OSX.

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

### Scoop (Windows)

```powershell
scoop bucket add armactl https://github.com/brittonhayes/armactl.git
scoop install armactl
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

### Mod Key Management

To see all available commands, run `armactl keys --help`.

List all keys in a directory or copy them to a directory.

```bash
# List all the *.bikeys in a directory
armactl keys ls -d /path/to/keys

# Copy keys to a directory
armactl keys cp --from '/path/to/keys/' --to '/path/to/destination/'
```

### Start a Server

To see all available commands, run `armactl server --help`.

Start a server with the given configuration.

```bash
# Start a server with two headless clients
armactl server --mods-local="@CBA_A3;@RHSUSAF" --headless-clients=2 --verbose

# Start a server in dry-run mode (no server will be started and it will print the commands that would be run)
armactl server  --verbose --dry-run

# View the server help
armactl server --help

# Usage: armactl server

# Run ARMA3 dedicated server.

# Flags:
#   -h, --help                     Show context-sensitive help.
#   -v, --verbose                  Enable verbose logging.

#       --binary="./arma3server_x64"
#                                  ($ARMA_BINARY)
#       --cdlc=STRING              ($ARMA_CDLC)
#       --config="server.cfg"      ($ARMA_CONFIG)
#       --port=2302                ($ARMA_PORT)
#       --skip-install             ($SKIP_INSTALL)
#       --mods-local=;...          ($ARMA_MODS)
#       --mods-server=;...         ($ARMA_MODS_SERVER)
#       --world="empty"            ($ARMA_WORLD)
#       --limit-fps=100            ($ARMA_LIMITFPS)
#       --headless-clients=0       ($HEADLESS_CLIENTS)
#       --headless-clients-profile="hc"
#                                  ($HEADLESS_CLIENTS_PROFILE)
#       --headless-clients-server="127.0.0.1"
#                                  ($HEADLESS_CLIENTS_SERVER)
#       --params=STRING            ($ARMA_PARAMS)
#       --profile="main"           ($ARMA_PROFILE)
#       --steam-branch="public"    ($STEAM_BRANCH)
#       --steam-branch-password=STRING
#                                  ($STEAM_BRANCH_PASSWORD)
#       --steam-user=STRING        ($STEAM_USER)
#       --steam-password=STRING    ($STEAM_PASSWORD)
#   -d, --dry-run                  ($DRY_RUN)
```

### Steam Server Querying

To see all available commands, run `armactl steam --help`.

Query a server for its current status.

```bash
# Inspect the steam server at this address
armactl steam inspect 127.0.0.1:2303

# 6:37PM INF game="Antistasi Plus - Harvest Blue" maxplayers=20 name="Antistasi" players=4
```

## Acknowledgements

- Mod syncing is performed using the packages from the amazing [rclone](https://github.com/rclone/rclone) tool.
- Steam API querying is built ontop of the [go-a2s](github.com/rumblefrog/go-a2s) library.