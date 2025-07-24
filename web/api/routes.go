package api

import (
	"net/http"

	"github.com/brittonhayes/armactl/web"
	"github.com/brittonhayes/armactl/web/api/middleware"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

// NewRouter creates a new chi router with all routes and middleware
func NewRouter(ctx *web.Context, devMode bool) chi.Router {
	services := web.NewServices(ctx)
	
	r := chi.NewRouter()

	// Basic middleware
	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.Recoverer)
	r.Use(middleware.Logger(ctx.Log))

	// CORS middleware for development
	if devMode {
		r.Use(middleware.CORS())
	}

	// API routes
	r.Route("/api", func(r chi.Router) {
		r.Mount("/steam", services.Steam.Routes())
		r.Mount("/mods", services.Mods.Routes())
		r.Mount("/server", services.Server.Routes())
		r.Mount("/keys", services.Keys.Routes())
		
		// Health check endpoint
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"status":"ok","service":"armactl-api"}`))
		})
	})

	// Serve static assets in production mode
	if !devMode {
		web.AddStaticRoutes(r)
	}

	return r
}