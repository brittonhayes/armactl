package commands

import (
	"fmt"
	"net/http"
	"time"

	"github.com/brittonhayes/armactl/web"
	"github.com/brittonhayes/armactl/web/api"
)

// UICmd represents the UI command
type UICmd struct {
	Port int    `help:"Port to serve the web interface on" default:"3000" env:"ARMACTL_UI_PORT"`
	Host string `help:"Host to bind the web server to" default:"localhost" env:"ARMACTL_UI_HOST"`
	Dev  bool   `help:"Enable development mode (CORS, etc.)" default:"false" env:"ARMACTL_UI_DEV"`
}

// Run executes the UI command
func (c *UICmd) Run(ctx *Context) error {
	// Convert commands.Context to web.Context
	webCtx := &web.Context{
		Ctx:     ctx.Ctx,
		Log:     ctx.Log,
		Output:  ctx.Output,
		Version: ctx.Version,
	}

	// Create the router with middleware
	router := api.NewRouter(webCtx, c.Dev)

	// Create HTTP server
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", c.Host, c.Port),
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	ctx.Log.Info().
		Str("host", c.Host).
		Int("port", c.Port).
		Bool("dev", c.Dev).
		Msg("Starting armactl web interface")

	if c.Dev {
		ctx.Log.Info().Msg("Development mode enabled - CORS is active")
		ctx.Log.Info().Msg("React dev server should be running on port 5173")
	}

	ctx.Log.Info().
		Str("url", fmt.Sprintf("http://%s:%d", c.Host, c.Port)).
		Msg("Web interface is ready")

	// Start the server
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		ctx.Log.Error().Err(err).Msg("Failed to start web server")
		return err
	}

	return nil
}
