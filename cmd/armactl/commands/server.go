package commands

import (
	"github.com/brittonhayes/armactl/server"
)

type ServerOptions struct {
	Binary                 string   `env:"ARMA_BINARY" default:"./arma3server_x64"`
	CDLC                   string   `env:"ARMA_CDLC"`
	Config                 string   `env:"ARMA_CONFIG" default:"server.cfg"`
	Port                   uint16   `env:"ARMA_PORT" default:"2302"`
	SkipInstall            bool     `env:"SKIP_INSTALL" default:"false"`
	ModsLocal              []string `env:"ARMA_MODS" sep:";" default:""`
	ModsServer             []string `env:"ARMA_MODS_SERVER" sep:";"  default:""`
	World                  string   `env:"ARMA_WORLD" default:"empty"`
	LimitFPS               uint8    `env:"ARMA_LIMITFPS" default:"100"`
	HeadlessClients        uint8    `env:"HEADLESS_CLIENTS" aliases:"hc" default:"0"`
	HeadlessClientsProfile string   `env:"HEADLESS_CLIENTS_PROFILE" aliases:"hcp" default:"hc"`
	HeadlessClientsServer  string   `env:"HEADLESS_CLIENTS_SERVER" default:"127.0.0.1"`
	Params                 string   `env:"ARMA_PARAMS"`
	Profile                string   `env:"ARMA_PROFILE" default:"main"`
	SteamBranch            string   `env:"STEAM_BRANCH" default:"public"`
	SteamBranchPassword    string   `env:"STEAM_BRANCH_PASSWORD"`
	SteamUser              string   `env:"STEAM_USER"`
	SteamPassword          string   `env:"STEAM_PASSWORD"`
	DryRun                 bool     `env:"DRY_RUN" short:"d" default:"false"`
}

type ServerCmd struct {
	*ServerOptions `embed:""`
}

func (c *ServerCmd) Run(ctx *Context) error {

	s := server.New(
		server.WithBinary(c.Binary),
		server.WithCDLC(c.CDLC),
		server.WithConfig(c.Config),
		server.WithPort(c.Port),
		server.WithModsLocal(c.ModsLocal),
		server.WithModsServer(c.ModsServer),
		server.WithWorld(c.World),
		server.WithLimitFPS(c.LimitFPS),
		server.WithHeadlessClients(c.HeadlessClients),
		server.WithHeadlessClientsProfile(c.HeadlessClientsProfile),
		server.WithHeadlessClientsServer(c.HeadlessClientsServer),
		server.WithParams(c.Params),
		server.WithProfile(c.Profile),
		server.WithSteamBranch(c.SteamBranch),
		server.WithSteamBranchPassword(c.SteamBranchPassword),
		server.WithSteamUser(c.SteamUser),
		server.WithSteamPassword(c.SteamPassword),
		server.WithDryRun(c.DryRun),
	)

	if c.DryRun {
		ctx.Log.Info().Msg("Dry-run enabled the server will not be started")
	}

	ctx.Log.Info().Msg("Starting dedicated server")
	err := s.Start(ctx.Ctx, ctx.Output)
	if err != nil {
		return err
	}

	if c.HeadlessClients > 0 {
		ctx.Log.Info().Uint8("clients", c.HeadlessClients).Msg("Starting headless clients")
		err = s.StartHeadless(ctx.Ctx, ctx.Output)
		if err != nil {
			return err
		}
	}

	return nil
}
