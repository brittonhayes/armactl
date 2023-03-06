package main

import (
	"context"
	"os"

	"github.com/alecthomas/kong"
	"github.com/brittonhayes/armactl/cmd/armactl/commands"
	"github.com/rs/zerolog"
)

var (
	Version = "vX.X.X"
)

var cli struct {
	Verbose bool `help:"Enable verbose logging." default:"false" short:"v"`

	Version commands.VersionCmd `cmd:"" help:"Print version of armactl." aliases:"v"`
	Steam   commands.SteamCmd   `cmd:"" help:"Query ARMA Steam server for information."`
	Keys    commands.KeysCmd    `cmd:"" help:"Manage bikeys for ARMA modifications."`
	Mods    commands.ModsCmd    `cmd:"" help:"Manage ARMA modifications and presets."`
	Server  commands.ServerCmd  `cmd:"" help:"Run ARMA3 dedicated server."`
}

func main() {
	ctx := kong.Parse(&cli,
		kong.Name("armactl"),
		kong.Description("ARMACTL is a command line tool to simplify the management of ARMA 3 dedicated servers."),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			NoExpandSubcommands: true,
		}),
	)

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger().Level(zerolog.InfoLevel).Output(zerolog.ConsoleWriter{Out: os.Stderr})
	if cli.Verbose {
		logger = logger.Level(zerolog.DebugLevel)
	}

	err := ctx.Run(&commands.Context{
		Ctx:     context.Background(),
		Log:     logger,
		Version: Version,
		Output:  os.Stdout,
	})

	ctx.FatalIfErrorf(err)
}
