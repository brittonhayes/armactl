package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/brittonhayes/armactl/web/api/types"
	"github.com/brittonhayes/armactl/web/services"
	"github.com/go-chi/chi/v5"
)

// ServerHandler handles server management related requests
type ServerHandler struct {
	service services.ServerService
}

// NewServerHandler creates a new server handler
func NewServerHandler(service services.ServerService) *ServerHandler {
	return &ServerHandler{
		service: service,
	}
}

// Routes sets up the server routes
func (h *ServerHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/status", h.getStatus)
	r.Post("/start", h.startServer)
	r.Post("/stop", h.stopServer)

	return r
}

// getStatus returns the current status of the ARMA server
func (h *ServerHandler) getStatus(w http.ResponseWriter, r *http.Request) {
	status, err := h.service.GetStatus(r.Context())
	if err != nil {
		h.writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.writeJSON(w, status)
}

// startServer starts the ARMA server with the given options
func (h *ServerHandler) startServer(w http.ResponseWriter, r *http.Request) {
	var opts types.ServerOptions

	if err := json.NewDecoder(r.Body).Decode(&opts); err != nil {
		h.writeError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// Set defaults if not provided
	if opts.Binary == "" {
		opts.Binary = "./arma3server_x64"
	}
	if opts.Config == "" {
		opts.Config = "server.cfg"
	}
	if opts.Port == 0 {
		opts.Port = 2302
	}
	if opts.World == "" {
		opts.World = "empty"
	}
	if opts.LimitFps == 0 {
		opts.LimitFps = 100
	}
	if opts.Profile == "" {
		opts.Profile = "main"
	}
	if opts.HeadlessClientsProfile == "" {
		opts.HeadlessClientsProfile = "hc"
	}
	if opts.HeadlessClientsServer == "" {
		opts.HeadlessClientsServer = "127.0.0.1"
	}

	err := h.service.Start(r.Context(), &opts)
	if err != nil {
		h.writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.writeJSON(w, map[string]string{"status": "success", "message": "server started successfully"})
}

// stopServer stops the currently running ARMA server
func (h *ServerHandler) stopServer(w http.ResponseWriter, r *http.Request) {
	err := h.service.Stop(r.Context())
	if err != nil {
		h.writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.writeJSON(w, map[string]string{"status": "success", "message": "server stopped successfully"})
}

// writeJSON writes a JSON response
func (h *ServerHandler) writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// writeError writes an error response
func (h *ServerHandler) writeError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	errorResp := &types.ErrorResponse{
		Error:   message,
		Code:    int32(code),
		Message: message,
	}

	json.NewEncoder(w).Encode(errorResp)
}
