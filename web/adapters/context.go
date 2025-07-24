package adapters

import (
	"context"
	"io"

	"github.com/rs/zerolog"
)

// WebContext represents the context for web adapters
type WebContext struct {
	Ctx     context.Context
	Log     zerolog.Logger
	Output  io.Writer
	Version string
}