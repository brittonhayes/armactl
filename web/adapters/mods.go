package adapters

import (
	"context"
	"errors"
	"os"

	"github.com/brittonhayes/armactl/internal/mods"
	"github.com/brittonhayes/armactl/internal/preset"
	"github.com/brittonhayes/armactl/web/api/types"
	"github.com/brittonhayes/armactl/web/services"
	"github.com/google/go-cmp/cmp"
)

// modsAdapter implements the ModsService interface
type modsAdapter struct {
	ctx *WebContext
}

// NewModsAdapter creates a new mods adapter
func NewModsAdapter(ctx *WebContext) services.ModsService {
	return &modsAdapter{
		ctx: ctx,
	}
}

// ListFromDirectory lists all mods in a directory
func (m *modsAdapter) ListFromDirectory(ctx context.Context, directory string) ([]string, error) {
	modsList, err := mods.FromDirectory(directory)
	if err != nil {
		return nil, err
	}

	if len(modsList) == 0 {
		return nil, errors.New("no mods found")
	}

	m.ctx.Log.Debug().Str("directory", directory).Strs("mods", modsList).Send()
	return modsList, nil
}

// ListFromPreset lists all mods from a preset HTML file
func (m *modsAdapter) ListFromPreset(ctx context.Context, presetPath string) ([]string, error) {
	f, err := os.Open(presetPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	p, err := preset.New().Parse(f)
	if err != nil {
		return nil, err
	}

	if len(p.Mods) == 0 {
		return nil, errors.New("no mods found")
	}

	m.ctx.Log.Debug().Str("preset", presetPath).Strs("mods", p.Mods).Send()
	return p.Mods, nil
}

// ValidatePreset validates that mods in a directory match a preset
func (m *modsAdapter) ValidatePreset(ctx context.Context, directory, presetPath string) (*types.ValidationResult, error) {
	if directory == "" || presetPath == "" {
		return nil, errors.New("both directory and preset must be provided")
	}

	installedMods, err := m.ListFromDirectory(ctx, directory)
	if err != nil {
		return &types.ValidationResult{
			Valid:        false,
			ErrorMessage: func() *string { s := err.Error(); return &s }(),
		}, nil
	}

	presetMods, err := m.ListFromPreset(ctx, presetPath)
	if err != nil {
		return &types.ValidationResult{
			Valid:        false,
			ErrorMessage: func() *string { s := err.Error(); return &s }(),
		}, nil
	}

	if diff := cmp.Diff(installedMods, presetMods); diff != "" {
		// Find missing and extra mods
		installedMap := make(map[string]bool)
		presetMap := make(map[string]bool)

		for _, mod := range installedMods {
			installedMap[mod] = true
		}
		for _, mod := range presetMods {
			presetMap[mod] = true
		}

		var missing, extra []string
		for _, mod := range presetMods {
			if !installedMap[mod] {
				missing = append(missing, mod)
			}
		}
		for _, mod := range installedMods {
			if !presetMap[mod] {
				extra = append(extra, mod)
			}
		}

		return &types.ValidationResult{
			Valid:   false,
			Missing: missing,
			Extra:   extra,
		}, nil
	}

	m.ctx.Log.Info().Msg("mods in preset match mod directory")
	return &types.ValidationResult{Valid: true}, nil
}

// SyncMods syncs mods from source directories to destination directories
func (m *modsAdapter) SyncMods(ctx context.Context, fromDirs, toDirs []string) error {
	if len(fromDirs) != len(toDirs) {
		return errors.New("number of source directories must match number of destination directories")
	}

	s := mods.NewSync()
	for i := range fromDirs {
		m.ctx.Log.Debug().Str("to", toDirs[i]).Str("from", fromDirs[i]).Msg("syncing mods...")
		err := s.Sync(m.ctx.Ctx, toDirs[i], fromDirs[i])
		if err != nil {
			return err
		}

		m.ctx.Log.Info().Str("to", toDirs[i]).Str("from", fromDirs[i]).Msg("synced mods")
	}

	return nil
}
