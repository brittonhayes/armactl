package adapters

import (
	"bytes"
	"context"
	"errors"
	"os/exec"
	"strconv"
	"strings"

	"github.com/brittonhayes/armactl/internal/server"
	"github.com/brittonhayes/armactl/web/api/types"
	"github.com/brittonhayes/armactl/web/services"
)

// serverAdapter implements the ServerService interface
type serverAdapter struct {
	ctx *WebContext
}

// NewServerAdapter creates a new server adapter
func NewServerAdapter(ctx *WebContext) services.ServerService {
	return &serverAdapter{
		ctx: ctx,
	}
}

// GetStatus returns the current status of the ARMA server
func (s *serverAdapter) GetStatus(ctx context.Context) (*types.ServerStatus, error) {
	// Check if arma3server_x64 process is running
	cmd := exec.CommandContext(ctx, "pgrep", "-f", "arma3server")
	output, err := cmd.Output()

	if err != nil {
		// Process not found
		return &types.ServerStatus{
			Running: false,
		}, nil
	}

	// Parse PID if process is running
	pidStr := strings.TrimSpace(string(output))
	if pidStr != "" {
		pid, err := strconv.Atoi(strings.Split(pidStr, "\n")[0])
		if err != nil {
			return &types.ServerStatus{
				Running: false,
				Error:   func() *string { s := "failed to parse process ID"; return &s }(),
			}, nil
		}

		return &types.ServerStatus{
			Running: true,
			Pid:     func() *int32 { p := int32(pid); return &p }(),
		}, nil
	}

	return &types.ServerStatus{
		Running: false,
	}, nil
}

// Start starts the ARMA server with the given options
func (s *serverAdapter) Start(ctx context.Context, opts *types.ServerOptions) error {
	// Convert API options to server options
	serverOpts := []server.ServerOption{
		server.WithBinary(opts.Binary),
		server.WithConfig(opts.Config),
		server.WithPort(uint16(opts.Port)),
		server.WithModsLocal(opts.ModsLocal),
		server.WithModsServer(opts.ModsServer),
		server.WithWorld(opts.World),
		server.WithLimitFPS(uint8(opts.LimitFps)),
		server.WithHeadlessClients(uint8(opts.HeadlessClients)),
		server.WithHeadlessClientsProfile(opts.HeadlessClientsProfile),
		server.WithHeadlessClientsServer(opts.HeadlessClientsServer),
		server.WithProfile(opts.Profile),
	}

	// Create server instance
	srv := server.New(serverOpts...)

	// Use a buffer to capture output
	var buf bytes.Buffer

	s.ctx.Log.Info().Msg("Starting ARMA server...")
	err := srv.Start(ctx, &buf)
	if err != nil {
		s.ctx.Log.Error().Err(err).Msg("Failed to start server")
		return err
	}

	// Start headless clients if requested
	if opts.HeadlessClients > 0 {
		s.ctx.Log.Info().Uint8("clients", uint8(opts.HeadlessClients)).Msg("Starting headless clients")
		err = srv.StartHeadless(ctx, &buf)
		if err != nil {
			s.ctx.Log.Error().Err(err).Msg("Failed to start headless clients")
			return err
		}
	}

	s.ctx.Log.Info().Msg("Server started successfully")
	return nil
}

// Stop stops the currently running ARMA server
func (s *serverAdapter) Stop(ctx context.Context) error {
	// Find and kill arma3server processes
	cmd := exec.CommandContext(ctx, "pkill", "-f", "arma3server")
	err := cmd.Run()
	if err != nil {
		// If pkill fails, the process might not be running
		s.ctx.Log.Debug().Err(err).Msg("No server process found to stop")
		return errors.New("no server process found or failed to stop server")
	}

	s.ctx.Log.Info().Msg("Server stopped successfully")
	return nil
}
