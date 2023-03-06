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
	From string `help:"Directory to recursively find and copy *.bikey files." type:"existingDir" short:"s" required:"true"`
	To   string `help:"Directory to copy keys to." type:"existingDir" short:"d" required:"true"`
}

func (c *KeysCopyCmd) Run(ctx *Context) error {
	bikey := mods.NewBikey()
	keys, err := bikey.Find(ctx.Ctx, c.From)
	if err != nil {
		return err
	}

	if len(keys) == 0 {
		return ErrNoKeysFound
	}

	return bikey.Copy(ctx.Ctx, keys, c.To)
}
