package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/brittonhayes/armactl/reforger"
)

type ReforgerCmd struct {
	Validate   ValidateCmd   `cmd:"" help:"Validate a reforger server config."`
	Start      StartCmd      `cmd:"" help:"Start Arma Reforger dedicated server."`
	Stop       StopCmd       `cmd:"" help:"Stop a running Arma Reforger server."`
	Update     UpdateCmd     `cmd:"" help:"Download or update Arma Reforger server via SteamCMD."`
	InitConfig InitConfigCmd `cmd:"" help:"Generate a Reforger server config template." aliases:"init"`
}

type ValidateCmd struct {
	Config string `arg:"" help:"Path to reforger config file." type:"existingFile"`
}

func (c *ValidateCmd) Run(ctx *Context) error {
	file, err := os.ReadFile(c.Config)
	if err != nil {
		return err
	}

	var config interface{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return err
	}

	validator := reforger.NewConfigValidator()
	err = validator.Validate(context.Background(), config)
	if err != nil {
		return err
	}

	fmt.Fprintln(ctx.Output, "Config is valid")
	return nil
}

// StartCmd starts the Reforger dedicated server
type StartCmd struct {
	Exe            string `help:"Path to ArmaReforgerServer.exe" env:"ARMA_REFORGER_SERVER_EXE" type:"path"`
	Config         string `help:"Path to server config JSON file." required:"true" type:"existingFile"`
	MaxFPS         int    `help:"Maximum FPS limit for server." default:"60"`
	Profile        string `help:"Path to profile directory." type:"path"`
	Detach         bool   `help:"Start server in background (detached mode)." default:"false" short:"d"`
	SkipValidation bool   `help:"Skip config schema validation (useful for custom configs)." default:"false"`
	DryRun         bool   `help:"Print command without executing." default:"false"`
}

func (c *StartCmd) Run(ctx *Context) error {
	// Discover executable path
	exe, err := reforger.DiscoverExecutable(c.Exe)
	if err != nil {
		return err
	}

	ctx.Log.Info().
		Str("executable", exe).
		Str("config", c.Config).
		Bool("detach", c.Detach).
		Msg("Starting Reforger server")

	// Convert config to absolute path
	absConfig, err := filepath.Abs(c.Config)
	if err != nil {
		return fmt.Errorf("failed to resolve config path: %w", err)
	}

	// Create server instance
	s := reforger.New(
		reforger.WithExecutable(exe),
		reforger.WithConfig(absConfig),
		reforger.WithMaxFPS(c.MaxFPS),
		reforger.WithProfile(c.Profile),
		reforger.WithDetach(c.Detach),
		reforger.WithSkipValidation(c.SkipValidation),
		reforger.WithDryRun(c.DryRun),
	)

	if c.DryRun {
		ctx.Log.Info().Msg("Dry-run enabled, server will not be started")
	}

	// Start server
	if err := s.Start(ctx.Ctx, ctx.Output); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	if !c.Detach && !c.DryRun {
		ctx.Log.Info().Msg("Server stopped")
	}

	return nil
}

// StopCmd stops a running Reforger server
type StopCmd struct {
	PID int `help:"Process ID of the server to stop." required:"true"`
}

func (c *StopCmd) Run(ctx *Context) error {
	ctx.Log.Info().Int("pid", c.PID).Msg("Stopping Reforger server")

	if err := reforger.StopProcess(c.PID); err != nil {
		return err
	}

	fmt.Fprintf(ctx.Output, "Server with PID %d stopped successfully\n", c.PID)
	return nil
}

// UpdateCmd downloads or updates the Reforger server via SteamCMD
type UpdateCmd struct {
	InstallDir    string `help:"Server installation directory." default:"C:\\ArmaReforgerServer" env:"ARMA_REFORGER_INSTALL_DIR"`
	SteamUser     string `help:"Steam username (use 'anonymous' for anonymous login)." env:"STEAM_USER"`
	SteamPassword string `help:"Steam password." env:"STEAM_PASSWORD"`
	SteamBranch   string `help:"Steam branch to install (e.g., 'experimental')." env:"STEAM_BRANCH"`
	SteamCmdPath  string `help:"Path to steamcmd executable." default:"steamcmd.exe" env:"STEAMCMD_PATH"`
	DryRun        bool   `help:"Print command without executing." default:"false"`
}

func (c *UpdateCmd) Run(ctx *Context) error {
	ctx.Log.Info().
		Str("installDir", c.InstallDir).
		Str("steamUser", c.SteamUser).
		Msg("Updating Reforger server")

	// Create server instance with Steam options
	s := reforger.New(
		reforger.WithInstallDir(c.InstallDir),
		reforger.WithSteamUser(c.SteamUser),
		reforger.WithSteamPassword(c.SteamPassword),
		reforger.WithSteamBranch(c.SteamBranch),
		reforger.WithSteamCmdPath(c.SteamCmdPath),
		reforger.WithDryRun(c.DryRun),
	)

	if c.DryRun {
		ctx.Log.Info().Msg("Dry-run enabled, SteamCMD will not be executed")
	}

	// Update server
	if err := s.Update(ctx.Ctx, ctx.Output); err != nil {
		return fmt.Errorf("failed to update server: %w", err)
	}

	return nil
}

// InitConfigCmd generates a config template
type InitConfigCmd struct {
	Output    string `help:"Output path for config file." default:"reforger-server.json" short:"o"`
	Commented bool   `help:"Include comments in the template." default:"true" short:"c"`
}

func (c *InitConfigCmd) Run(ctx *Context) error {
	ctx.Log.Info().Str("output", c.Output).Msg("Generating config template")

	// Check if file already exists
	if _, err := os.Stat(c.Output); err == nil {
		return fmt.Errorf("file already exists: %s (remove it first)", c.Output)
	}

	// Generate template
	if err := reforger.GenerateConfigTemplate(c.Output, c.Commented); err != nil {
		return fmt.Errorf("failed to generate config: %w", err)
	}

	fmt.Fprintf(ctx.Output, "Config template created: %s\n", c.Output)

	if c.Commented {
		fmt.Fprintln(ctx.Output, "Note: Remove comments before using with the server (JSON doesn't support comments)")
	}

	return nil
}
