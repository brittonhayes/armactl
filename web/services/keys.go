package services

import (
	"context"
)

// KeysService defines the interface for key management operations
type KeysService interface {
	// ListKeys lists all .bikey files in a directory
	ListKeys(ctx context.Context, directory string) ([]string, error)
	
	// CopyKeys copies keys from a source directory to a destination directory
	CopyKeys(ctx context.Context, fromDir, toDir string) error
}