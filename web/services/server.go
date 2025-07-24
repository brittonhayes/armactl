package services

import (
	"context"

	"github.com/brittonhayes/armactl/web/api/types"
)

// ServerService defines the interface for ARMA server operations
type ServerService interface {
	// GetStatus returns the current status of the ARMA server
	GetStatus(ctx context.Context) (*types.ServerStatus, error)
	
	// Start starts the ARMA server with the given options
	Start(ctx context.Context, opts *types.ServerOptions) error
	
	// Stop stops the currently running ARMA server
	Stop(ctx context.Context) error
}