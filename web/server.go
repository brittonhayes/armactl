package web

import (
	"embed"
	"io/fs"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

//go:embed ui/dist
var webAssets embed.FS

// NewFileServer creates a file server for embedded web assets
func NewFileServer() http.Handler {
	// Get the embedded filesystem starting from ui/dist
	uiAssets, err := fs.Sub(webAssets, "ui/dist")
	if err != nil {
		// In development or if assets don't exist, return a simple handler
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`
<!DOCTYPE html>
<html>
<head>
    <title>armactl</title>
    <style>
        body { font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif; margin: 2rem; }
        .container { max-width: 600px; margin: 0 auto; text-align: center; }
        .code { background: #f5f5f5; padding: 0.25rem 0.5rem; border-radius: 3px; font-family: monospace; }
    </style>
</head>
<body>
    <div class="container">
        <h1>armactl Web Interface</h1>
        <p>The React UI has not been built yet.</p>
        <p>To build the UI, run:</p>
        <p><code class="code">task build:ui</code></p>
        <p>API endpoints are available at <code class="code">/api/*</code></p>
        <ul style="text-align: left; max-width: 400px; margin: 2rem auto;">
            <li><code class="code">GET /api/health</code> - Health check</li>
            <li><code class="code">GET /api/server/status</code> - Server status</li>
            <li><code class="code">GET /api/steam/query/{address}</code> - Query Steam server</li>
            <li><code class="code">GET /api/mods?directory=/path</code> - List mods</li>
            <li><code class="code">GET /api/keys?directory=/path</code> - List keys</li>
        </ul>
    </div>
</body>
</html>`))
		})
	}

	return http.FileServer(http.FS(uiAssets))
}

// AddStaticRoutes adds routes for serving static web assets
func AddStaticRoutes(r chi.Router) {
	fileServer := NewFileServer()
	
	// Serve static assets
	r.Handle("/static/*", fileServer)
	r.Handle("/assets/*", fileServer)
	
	// Handle SPA routing - serve index.html for all non-API routes
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		// Don't serve SPA for API routes
		if strings.HasPrefix(r.URL.Path, "/api/") {
			http.NotFound(w, r)
			return
		}
		
		// Serve index.html for all other routes (SPA routing)
		fileServer.ServeHTTP(w, r)
	})
}