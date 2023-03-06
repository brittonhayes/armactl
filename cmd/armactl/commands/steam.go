package commands

import (
	"github.com/brittonhayes/armactl/steam"
)

type SteamCmd struct {
	Inspect SteamInspectCmd `cmd:"" help:"Inspect an arma steam server."`
}

type SteamInspectCmd struct {
	Address string `arg:"true" help:"Host of the server to query. (default localhost:2303)" default:"localhost:2303"`
}

func (c *SteamInspectCmd) Run(ctx *Context) error {
	s := steam.New(c.Address)
	defer s.Close()

	ctx.Log.Debug().Str("address", c.Address).Msg("querying server...")
	i, err := s.Inspect()
	if err != nil {
		return err
	}

	ctx.Log.Info().
		Str("name", i.Name).
		Str("game", i.Game).
		Uint8("players", i.PlayerCount).
		Uint8("maxplayers", i.MaxPlayerCount).
		Send()

	return nil
}
