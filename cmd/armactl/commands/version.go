package commands

import "fmt"

type VersionCmd struct{}

func (cmd *VersionCmd) Run(ctx *Context) error {
	fmt.Printf("armactl %s", ctx.Version)
	return nil
}
