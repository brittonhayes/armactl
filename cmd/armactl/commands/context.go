package commands

import (
	"context"

	"github.com/rs/zerolog"
)

type Context struct {
	Ctx     context.Context
	Log     zerolog.Logger
	Version string
}
