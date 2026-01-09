package reforger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

// Server interface defines operations for managing Reforger servers
type Server interface {
	Start(ctx context.Context, w io.Writer) error
	Stop(ctx context.Context) error
	Update(ctx context.Context, w io.Writer) error
}

// ReforgerServer implements the Server interface
type ReforgerServer struct {
	Executable      string
	Config          string
	MaxFPS          int
	Profile         string
	Detach          bool
	DryRun          bool
	SkipValidation  bool
	Validator       ConfigValidator
	InstallDir      string
	SteamUser       string
	SteamPassword   string
	SteamBranch     string
	SteamCmdPath    string
}

// New creates a new Reforger server with functional options
func New(options ...ServerOption) Server {
	s := &ReforgerServer{
		Validator: NewConfigValidator(),
	}
	for _, option := range options {
		option(s)
	}
	return s
}

// buildArgs constructs command-line arguments for ArmaReforgerServer
func (s *ReforgerServer) buildArgs() string {
	args := fmt.Sprintf("-config=%q", s.Config)

	if s.MaxFPS > 0 {
		args = fmt.Sprintf("%s -maxFPS=%d", args, s.MaxFPS)
	}

	if s.Profile != "" {
		args = fmt.Sprintf("%s -profile=%q", args, s.Profile)
	}

	return args
}

// Start starts the Reforger server
func (s *ReforgerServer) Start(ctx context.Context, w io.Writer) error {
	// Validate config file exists
	if _, err := os.Stat(s.Config); os.IsNotExist(err) {
		return fmt.Errorf("config file not found: %s. Use 'armactl reforger init-config' to create one", s.Config)
	}

	// Validate config against schema (unless skipped)
	if !s.SkipValidation {
		if err := s.validateConfig(ctx); err != nil {
			return fmt.Errorf("config validation failed (use --skip-validation to bypass): %w", err)
		}
	}

	// Build command arguments
	args := s.buildArgs()

	// Log/print the command that will be executed
	fmt.Fprintf(w, "%s %s\n", s.Executable, args)

	if s.DryRun {
		return nil
	}

	// Create command
	cmd := exec.CommandContext(ctx, s.Executable, args)

	if s.Detach {
		// Detach mode: Start process and return immediately
		if err := cmd.Start(); err != nil {
			return fmt.Errorf("failed to start server: %w", err)
		}

		fmt.Fprintf(w, "Server started with PID: %d\n", cmd.Process.Pid)
		return nil
	}

	// Foreground mode: Stream output in real-time
	cmd.Stdout = w
	cmd.Stderr = w

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("server exited with error: %w", err)
	}

	return nil
}

// validateConfig validates the server config JSON against the schema
func (s *ReforgerServer) validateConfig(ctx context.Context) error {
	file, err := os.ReadFile(s.Config)
	if err != nil {
		return err
	}

	var config interface{}
	if err := json.Unmarshal(file, &config); err != nil {
		return fmt.Errorf("invalid JSON: %w", err)
	}

	return s.Validator.Validate(ctx, config)
}

// Stop stops a running Reforger server
// Note: This method is not used directly; the StopCmd calls StopProcess() instead
func (s *ReforgerServer) Stop(ctx context.Context) error {
	return fmt.Errorf("stop operation should be called via 'armactl reforger stop --pid <PID>'")
}

// Update downloads or updates the Reforger server using SteamCMD
func (s *ReforgerServer) Update(ctx context.Context, w io.Writer) error {
	// Determine steamcmd executable
	steamcmd := s.SteamCmdPath
	if steamcmd == "" {
		steamcmd = "steamcmd.exe"
	}

	// Determine install directory
	installDir := s.InstallDir
	if installDir == "" {
		installDir = "C:\\ArmaReforgerServer"
	}

	// Build SteamCMD arguments
	// App ID for Arma Reforger Dedicated Server is 1874900
	args := []string{
		"+force_install_dir", installDir,
		"+app_update", "1874900",
	}

	// Add login credentials if provided
	if s.SteamUser != "" {
		args = append([]string{"+login", s.SteamUser, s.SteamPassword}, args...)
	} else {
		args = append([]string{"+login", "anonymous"}, args...)
	}

	// Add branch if specified (e.g., experimental)
	if s.SteamBranch != "" {
		args = append(args, "-beta", s.SteamBranch)
	}

	// Validate and quit
	args = append(args, "validate", "+quit")

	fmt.Fprintf(w, "Updating Arma Reforger server via SteamCMD...\n")
	fmt.Fprintf(w, "%s %v\n", steamcmd, args)

	if s.DryRun {
		return nil
	}

	// Execute SteamCMD
	cmd := exec.CommandContext(ctx, steamcmd, args...)
	cmd.Stdout = w
	cmd.Stderr = w

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("steamcmd failed: %w", err)
	}

	fmt.Fprintf(w, "\nServer installed/updated to: %s\n", installDir)
	return nil
}

// DiscoverExecutable attempts to find the Reforger server executable
// Priority: explicit path -> env var -> common install paths (SteamCMD and default)
func DiscoverExecutable(explicitPath string) (string, error) {
	// 1. Check explicit path
	if explicitPath != "" {
		if _, err := os.Stat(explicitPath); err == nil {
			return filepath.Abs(explicitPath)
		}
		return "", fmt.Errorf("executable not found at: %s", explicitPath)
	}

	// 2. Check environment variable
	if envPath := os.Getenv("ARMA_REFORGER_SERVER_EXE"); envPath != "" {
		if _, err := os.Stat(envPath); err == nil {
			return filepath.Abs(envPath)
		}
		return "", fmt.Errorf("executable not found at ARMA_REFORGER_SERVER_EXE: %s", envPath)
	}

	// 3. Check common paths
	commonPaths := []string{
		"C:\\ArmaReforgerServer\\ArmaReforgerServer.exe",                    // Default standalone
		"C:\\steamcmd\\steamapps\\common\\Arma Reforger Server\\ArmaReforgerServer.exe", // SteamCMD default
		"C:\\Program Files (x86)\\Steam\\steamapps\\common\\Arma Reforger Server\\ArmaReforgerServer.exe", // Steam client
	}

	for _, path := range commonPaths {
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}

	return "", fmt.Errorf("ArmaReforgerServer.exe not found. Use --exe flag, set ARMA_REFORGER_SERVER_EXE, or run 'armactl reforger update' to install via SteamCMD")
}
