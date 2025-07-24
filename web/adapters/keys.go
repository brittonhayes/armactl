package adapters

import (
	"context"
	"errors"

	"github.com/brittonhayes/armactl/internal/mods"
	"github.com/brittonhayes/armactl/web/services"
)

// keysAdapter implements the KeysService interface
type keysAdapter struct {
	ctx *WebContext
}

// NewKeysAdapter creates a new keys adapter
func NewKeysAdapter(ctx *WebContext) services.KeysService {
	return &keysAdapter{
		ctx: ctx,
	}
}

// ListKeys lists all .bikey files in a directory
func (k *keysAdapter) ListKeys(ctx context.Context, directory string) ([]string, error) {
	bikey := mods.NewBikey()
	keys, err := bikey.Find(ctx, directory)
	if err != nil {
		return nil, err
	}

	if len(keys) == 0 {
		return nil, errors.New("no keys found")
	}

	k.ctx.Log.Debug().Str("directory", directory).Strs("keys", keys).Send()
	return keys, nil
}

// CopyKeys copies keys from a source directory to a destination directory
func (k *keysAdapter) CopyKeys(ctx context.Context, fromDir, toDir string) error {
	bikey := mods.NewBikey()
	keys, err := bikey.Find(ctx, fromDir)
	if err != nil {
		return err
	}

	if len(keys) == 0 {
		return errors.New("no keys found")
	}

	k.ctx.Log.Debug().Str("from", fromDir).Str("to", toDir).Strs("keys", keys).Msg("copying keys...")
	err = bikey.Copy(ctx, keys, toDir)
	if err != nil {
		return err
	}

	k.ctx.Log.Info().Str("from", fromDir).Str("to", toDir).Msg("copied keys successfully")
	return nil
}
