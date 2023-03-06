package rclone

import (
	"context"
	"io"
	"os"
	"path/filepath"

	_ "github.com/rclone/rclone/backend/local"
	"github.com/rclone/rclone/fs"
	"github.com/rclone/rclone/fs/sync"
	"github.com/rclone/rclone/fs/walk"
)

type rclone struct{}

type FS interface {
	Find(ctx context.Context, path string, ext string) ([]string, error)
	Copy(ctx context.Context, destination string, source string) error
	Sync(ctx context.Context, destination string, source string) error
}

func New() FS {
	return &rclone{}
}

// Find retrieves all files in a directory with a given extension
func (r *rclone) Find(ctx context.Context, path string, ext string) ([]string, error) {
	var matches []string

	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	dir, err := fs.NewFs(context.Background(), path)
	if err != nil {
		return nil, err
	}

	err = walk.Walk(context.Background(), dir, "", true, 0, func(p string, d fs.DirEntries, err error) error {
		if err != nil {
			return err
		}

		for _, entry := range d {
			if filepath.Ext(entry.String()) == ext {
				matches = append(matches, filepath.Join(abs, entry.String()))
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return matches, nil
}

// Copy copies the file to the destination
func (r *rclone) Copy(ctx context.Context, destination string, file string) error {
	info, err := os.Stat(destination)
	if err != nil {
		return err
	}
	if info.IsDir() {
		destination = filepath.Join(destination, filepath.Base(file))
	}

	src, err := os.Open(file)
	if err != nil {
		return err
	}
	defer src.Close()

	dest, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer dest.Close()

	_, err = io.Copy(src, dest)
	if err != nil {
		return err
	}

	return nil
}

// Sync synchronizes all files from the source to the destination
func (r *rclone) Sync(ctx context.Context, destination string, source string) error {
	dest, err := fs.NewFs(ctx, destination)
	if err != nil {
		return err
	}

	src, err := fs.NewFs(ctx, source)
	if err != nil {
		return err
	}

	return sync.Sync(ctx, dest, src, false)
}
