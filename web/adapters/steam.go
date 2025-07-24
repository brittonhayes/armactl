package adapters

import (
	"context"

	"github.com/brittonhayes/armactl/internal/steam"
	"github.com/brittonhayes/armactl/web/api/types"
	"github.com/brittonhayes/armactl/web/services"
)

// steamAdapter implements the SteamService interface
type steamAdapter struct{}

// NewSteamAdapter creates a new steam adapter
func NewSteamAdapter() services.SteamService {
	return &steamAdapter{}
}

// QueryServer queries a Steam server for its information
func (s *steamAdapter) QueryServer(ctx context.Context, address string) (*types.ServerInfo, error) {
	client := steam.New(address)
	defer client.Close()

	server, err := client.Inspect()
	if err != nil {
		return nil, err
	}

	return &types.ServerInfo{
		Name:       server.Name,
		Game:       server.Game,
		Players:    uint32(server.PlayerCount),
		MaxPlayers: uint32(server.MaxPlayerCount),
	}, nil
}

// GetPlayers returns a list of players currently on the server
func (s *steamAdapter) GetPlayers(ctx context.Context, address string) ([]string, error) {
	client := steam.New(address)
	defer client.Close()

	return client.Players()
}
