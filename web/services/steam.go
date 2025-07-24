package services

import (
	"context"

	"github.com/brittonhayes/armactl/web/api/types"
)

// SteamService defines the interface for Steam server operations
type SteamService interface {
	// QueryServer queries a Steam server for its information
	QueryServer(ctx context.Context, address string) (*types.ServerInfo, error)
	
	// GetPlayers returns a list of players currently on the server
	GetPlayers(ctx context.Context, address string) ([]string, error)
}