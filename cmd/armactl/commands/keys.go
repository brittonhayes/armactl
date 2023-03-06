package commands

import (
	"fmt"

	"github.com/brittonhayes/armactl/mods"
)

type KeysCmd struct {
	List KeysListCmd `cmd:"" help:"List keys in a directory." aliases:"ls"`
	Copy KeysCopyCmd `cmd:"" help:"Copy keys to a directory." aliases:"cp"`
}

type KeysListCmd struct {
	Directory string `help:"Path to directory to list keys in." type:"existingDir" short:"d" required:"true"`
}

func (c *KeysListCmd) Run(ctx *Context) error {
	k := mods.NewBikey()
	keys, err := k.Find(ctx.Ctx, c.Directory)
	if err != nil {
		return err
	}

	if len(keys) == 0 {
		return ErrNoModsFound
	}

	for _, key := range keys {
		fmt.Fprintf(ctx.Output, "%s\n", key)
	}

	return nil
}

type KeysCopyCmd struct {
	Source      string `help:"Source to copy keys from." type:"existingDir" short:"s" required:"true"`
	Destination string `help:"Destination to copy keys to." type:"existingDir" short:"d" required:"true"`
}

func (c *KeysCopyCmd) Run(ctx *Context) error {
	k := mods.NewBikey()
	keys, err := k.Find(ctx.Ctx, c.Source)
	if err != nil {
		return err
	}

	if len(keys) == 0 {
		return ErrNoKeysFound
	}

	return k.Copy(ctx.Ctx, keys, c.Destination)
}
