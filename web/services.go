package web

import (
	"context"
	"io"

	"github.com/brittonhayes/armactl/web/adapters"
	"github.com/brittonhayes/armactl/web/api/handlers"
	"github.com/rs/zerolog"
)

// Context represents the context for web services (similar to commands.Context)
type Context struct {
	Ctx     context.Context
	Log     zerolog.Logger
	Output  io.Writer
	Version string
}

// Services holds all the service dependencies
type Services struct {
	Steam  *handlers.SteamHandler
	Mods   *handlers.ModsHandler
	Server *handlers.ServerHandler
	Keys   *handlers.KeysHandler
}

// NewServices creates a new services instance with all adapters
func NewServices(ctx *Context) *Services {
	// Convert web.Context to adapters.WebContext
	webCtx := &adapters.WebContext{
		Ctx:     ctx.Ctx,
		Log:     ctx.Log,
		Output:  ctx.Output,
		Version: ctx.Version,
	}

	return &Services{
		Steam:  handlers.NewSteamHandler(adapters.NewSteamAdapter()),
		Mods:   handlers.NewModsHandler(adapters.NewModsAdapter(webCtx)),
		Server: handlers.NewServerHandler(adapters.NewServerAdapter(webCtx)),
		Keys:   handlers.NewKeysHandler(adapters.NewKeysAdapter(webCtx)),
	}
}