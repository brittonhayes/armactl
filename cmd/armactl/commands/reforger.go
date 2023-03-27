package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/brittonhayes/armactl/reforger"
)

type ReforgerCmd struct {
	Validate ValidateCmd `cmd:"" help:"Validate a reforger server config."`
}

type ValidateCmd struct {
	Config string `arg:"" help:"Path to reforger config file." type:"existingFile"`
}

func (c *ValidateCmd) Run(ctx *Context) error {
	file, err := os.ReadFile(c.Config)
	if err != nil {
		return err
	}

	var config interface{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return err
	}

	validator := reforger.NewConfigValidator()
	err = validator.Validate(context.Background(), config)
	if err != nil {
		return err
	}

	fmt.Fprintln(ctx.Output, "Config is valid")
	return nil
}
