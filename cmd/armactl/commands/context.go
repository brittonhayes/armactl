package commands

import (
	"context"
	"io"

	"github.com/rs/zerolog"
)

type Context struct {
	Ctx     context.Context
	Log     zerolog.Logger
	Output  io.Writer
	Version string
}
