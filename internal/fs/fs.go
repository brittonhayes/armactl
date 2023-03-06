package fs

import (
	"context"
	"errors"
)

type fs struct{}

type FS interface {
	Find(ctx context.Context, path string, ext string) ([]string, error)
	Copy(ctx context.Context, source string, destination string) error
	Sync(ctx context.Context, source string, destination string) error
}

func NewFs() FS {
	return &fs{}
}

// Find retrieves all files in a directory with a given extension
func (r *fs) Find(ctx context.Context, path string, ext string) ([]string, error) {
	return nil, errors.New("not implemented")
}

// Copy copies a file from one location to another
func (r *fs) Copy(ctx context.Context, source string, destination string) error {
	return errors.New("not implemented")
}

// Sync copies the directory from the source to the destination
func (r *fs) Sync(ctx context.Context, source string, destination string) error {
	return errors.New("not implemented")
}
