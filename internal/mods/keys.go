package mods

import (
	"context"
	"errors"

	"github.com/brittonhayes/armactl/internal/rclone"
)

var (
	ErrDestinationNotDirectory = errors.New("destination is not a directory")
)

type bikey struct{}

type Bikey interface {
	Find(ctx context.Context, path string) ([]string, error)
	Copy(ctx context.Context, keys []string, destination string) error
}

func NewBikey() Bikey {
	return &bikey{}
}

// Find finds all bikeys in the given path
func (b *bikey) Find(ctx context.Context, path string) ([]string, error) {
	return rclone.New().Find(ctx, path, ".bikey")
}

func (b *bikey) Copy(ctx context.Context, keys []string, destination string) error {
	r := rclone.New()
	for _, key := range keys {
		if err := r.Copy(ctx, destination, key); err != nil {
			return err
		}
	}

	return nil
}
