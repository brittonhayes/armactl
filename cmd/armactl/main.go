package main

import (
	"context"
	"os"

	"github.com/alecthomas/kong"
	"github.com/brittonhayes/armactl/cmd/armactl/commands"
	"github.com/rs/zerolog"
)

var (
	version = "vX.X.X"
)

var cli struct {
	Verbose bool                `help:"Enable verbose logging." default:"false" short:"v"`
	Server  commands.ServerCmd  `cmd:"" help:"Query a server for information."`
	Mods    commands.ModsCmd    `cmd:"" help:"Parse an ARMA mods preset."`
	Version commands.VersionCmd `cmd:"" help:"Print the version of armactl." aliases:"v"`
}

func main() {
	ctx := kong.Parse(&cli,
		kong.Name("armactl"),
		kong.Description("CLI for managing ARMA3 dedicated servers."),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
	)

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger().Level(zerolog.InfoLevel).Output(zerolog.ConsoleWriter{Out: os.Stderr})
	if cli.Verbose {
		logger = logger.Level(zerolog.DebugLevel)
	}

	err := ctx.Run(&commands.Context{
		Ctx:     context.Background(),
		Log:     logger,
		Version: version,
		Output:  os.Stdout,
	})

	ctx.FatalIfErrorf(err)
}
