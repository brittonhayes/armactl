package mods

import (
	"context"

	_ "github.com/rclone/rclone/backend/local"

	"github.com/rclone/rclone/fs"
	rsync "github.com/rclone/rclone/fs/sync"
)

type sync struct{}

type Sync interface {
	Sync(ctx context.Context, destination string, source string) error
}

func NewSync() Sync {
	return &sync{}
}

func (s *sync) Sync(ctx context.Context, destination string, source string) error {
	dest, err := fs.NewFs(ctx, destination)
	if err != nil {
		return err
	}

	src, err := fs.NewFs(ctx, source)
	if err != nil {
		return err
	}

	return rsync.Sync(ctx, dest, src, false)
}
