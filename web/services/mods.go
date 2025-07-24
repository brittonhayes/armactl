package services

import (
	"context"

	"github.com/brittonhayes/armactl/web/api/types"
)

// ModsService defines the interface for mod management operations
type ModsService interface {
	// ListFromDirectory lists all mods in a directory
	ListFromDirectory(ctx context.Context, directory string) ([]string, error)
	
	// ListFromPreset lists all mods from a preset HTML file
	ListFromPreset(ctx context.Context, presetPath string) ([]string, error)
	
	// ValidatePreset validates that mods in a directory match a preset
	ValidatePreset(ctx context.Context, directory, presetPath string) (*types.ValidationResult, error)
	
	// SyncMods syncs mods from source directories to destination directories
	SyncMods(ctx context.Context, fromDirs, toDirs []string) error
}