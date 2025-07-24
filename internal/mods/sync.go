package mods

import (
	"context"

	"github.com/brittonhayes/armactl/internal/rclone"
)

type sync struct{}

type Syncer interface {
	Sync(ctx context.Context, destination string, source string) error
}

func NewSync() Syncer {
	return &sync{}
}

func (s *sync) Sync(ctx context.Context, destination string, source string) error {
	return rclone.New().Sync(ctx, destination, source)
}
