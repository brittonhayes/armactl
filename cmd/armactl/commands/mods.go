package commands

import (
	"errors"
	"fmt"
	"os"

	"github.com/brittonhayes/armactl/mods"
	"github.com/brittonhayes/armactl/preset"
	"github.com/google/go-cmp/cmp"
)

var (
	ErrNoKeysFound        = errors.New("no *.bikeys found in directory")
	ErrNoModsFound        = errors.New("no mods found in preset")
	ErrModMismatch        = errors.New("mods in preset do not match mods in directory")
	ErrMissingCompareArgs = errors.New("directory and preset must be provided to compare")
)

type ModsCmd struct {
	Keys KeysCmd     `cmd:"" help:"List keys in a directory." aliases:"k"`
	List ListModsCmd `cmd:"" help:"List mods in directory or HTML mod preset." aliases:"ls"`
	Sync SyncModsCmd `cmd:"" help:"Sync mods from a source directory to a destination." aliases:"s"`
}

type ListModsCmd struct {
	Check     bool   `help:"Check if mods from preset exist in mods directory." default:"false" short:"c"`
	Preset    string `help:"Path to preset file." type:"existingFile" short:"p"`
	Directory string `help:"Path to directory to check mods exist in." type:"existingDir" short:"d"`
}

func (c *ListModsCmd) Run(ctx *Context) error {
	var (
		presetMods    []string
		installedMods []string
	)

	if c.Directory == "" && c.Preset == "" {
		return errors.New("either directory or preset must be provided to list mods")
	}

	if c.Directory != "" {
		m, err := c.listFromDirectory(ctx)
		if err != nil {
			return err
		}

		installedMods = m
		fmt.Fprintln(ctx.Output, "\nInstalled Mods:")
		for i, mod := range m {
			fmt.Fprintf(ctx.Output, "%d. @%s\n", i+1, mod)
		}
	}

	if c.Preset != "" {
		p, err := c.listFromPreset(ctx)
		if err != nil {
			return err
		}
		presetMods = p
		fmt.Fprintln(ctx.Output, "\nPreset Mods:")
		for i, mod := range p {
			fmt.Fprintf(ctx.Output, "%d. @%s\n", i+1, mod)
		}
	}

	if c.Check {
		ctx.Log.Debug().Msg("checking if mods exist in directory...")
		err := c.compare(installedMods, presetMods)
		if err != nil {
			ctx.Log.Err(ErrModMismatch).Send()
			return err
		}

		ctx.Log.Info().Msg("mods in preset match mod directory")
	}

	return nil
}

func (c *ListModsCmd) listFromDirectory(ctx *Context) ([]string, error) {
	m, err := mods.FromDirectory(c.Directory)
	if err != nil {
		return nil, err
	}

	if len(m) == 0 {
		return nil, ErrNoModsFound
	}

	ctx.Log.Debug().Str("directory", c.Directory).Strs("mods", m).Send()
	return m, nil
}

func (c *ListModsCmd) listFromPreset(ctx *Context) ([]string, error) {
	f, err := os.Open(c.Preset)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	p, err := preset.New().Parse(f)
	if err != nil {
		return nil, err
	}

	if len(p.Mods) == 0 {
		return nil, ErrNoModsFound
	}

	ctx.Log.Debug().Str("preset", c.Preset).Strs("mods", p.Mods).Send()

	return p.Mods, nil
}

func (c *ListModsCmd) compare(installedMods, presetMods []string) error {
	if c.Directory == "" || c.Preset == "" {
		return ErrMissingCompareArgs
	}

	if diff := cmp.Diff(installedMods, presetMods); diff != "" {
		fmt.Println(diff)
		return ErrModMismatch
	}

	return nil
}

type SyncModsCmd struct {
	From []string `help:"Path(s) to directory to sync mods from." type:"existingDir" short:"f" required:"true"`
	To   []string `help:"Path(s) to directory to sync mods to." type:"existingDir" short:"t" required:"true"`
}

func (c *SyncModsCmd) Run(ctx *Context) error {
	if len(c.From) != len(c.To) {
		return fmt.Errorf("number of 'from' directories (%d) does not match number of 'to' directories (%d)", len(c.From), len(c.To))
	}

	s := mods.NewSync()
	for i := range c.From {
		ctx.Log.Debug().Str("to", c.To[i]).Str("from", c.From[i]).Msg("syncing mods...")
		err := s.Sync(ctx.Ctx, c.To[i], c.From[i])
		if err != nil {
			return err
		}

		ctx.Log.Info().Str("to", c.To[i]).Str("from", c.From[i]).Msg("synced mods")
	}
	return nil
}
