package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/brittonhayes/armactl/web/api/types"
	"github.com/brittonhayes/armactl/web/services"
	"github.com/go-chi/chi/v5"
)

// SteamHandler handles Steam server related requests
type SteamHandler struct {
	service services.SteamService
}

// NewSteamHandler creates a new steam handler
func NewSteamHandler(service services.SteamService) *SteamHandler {
	return &SteamHandler{
		service: service,
	}
}

// Routes sets up the steam routes
func (h *SteamHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/query/{address}", h.queryServer)
	r.Get("/players/{address}", h.getPlayers)

	return r
}

// queryServer queries a Steam server for its information
func (h *SteamHandler) queryServer(w http.ResponseWriter, r *http.Request) {
	address := chi.URLParam(r, "address")
	if address == "" {
		h.writeError(w, "address parameter is required", http.StatusBadRequest)
		return
	}

	server, err := h.service.QueryServer(r.Context(), address)
	if err != nil {
		h.writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.writeJSON(w, server)
}

// getPlayers returns a list of players currently on the server
func (h *SteamHandler) getPlayers(w http.ResponseWriter, r *http.Request) {
	address := chi.URLParam(r, "address")
	if address == "" {
		h.writeError(w, "address parameter is required", http.StatusBadRequest)
		return
	}

	players, err := h.service.GetPlayers(r.Context(), address)
	if err != nil {
		h.writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.writeJSON(w, map[string][]string{"players": players})
}

// writeJSON writes a JSON response
func (h *SteamHandler) writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// writeError writes an error response
func (h *SteamHandler) writeError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	errorResp := &types.ErrorResponse{
		Error:   message,
		Code:    int32(code),
		Message: message,
	}

	json.NewEncoder(w).Encode(errorResp)
}
