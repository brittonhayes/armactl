package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/brittonhayes/armactl/web/api/types"
	"github.com/brittonhayes/armactl/web/services"
	"github.com/go-chi/chi/v5"
)

// ModsHandler handles mod management related requests
type ModsHandler struct {
	service services.ModsService
}

// NewModsHandler creates a new mods handler
func NewModsHandler(service services.ModsService) *ModsHandler {
	return &ModsHandler{
		service: service,
	}
}

// Routes sets up the mods routes
func (h *ModsHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.listMods)
	r.Post("/validate", h.validatePreset)
	r.Post("/sync", h.syncMods)

	return r
}

// listMods lists mods from directory or preset
func (h *ModsHandler) listMods(w http.ResponseWriter, r *http.Request) {
	directory := r.URL.Query().Get("directory")
	preset := r.URL.Query().Get("preset")

	if directory == "" && preset == "" {
		h.writeError(w, "either directory or preset parameter is required", http.StatusBadRequest)
		return
	}

	response := &types.ModsListResponse{}

	if directory != "" {
		mods, err := h.service.ListFromDirectory(r.Context(), directory)
		if err != nil {
			h.writeError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.InstalledMods = mods
	}

	if preset != "" {
		mods, err := h.service.ListFromPreset(r.Context(), preset)
		if err != nil {
			h.writeError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.PresetMods = mods
	}

	h.writeJSON(w, response)
}

// validatePreset validates that mods in a directory match a preset
func (h *ModsHandler) validatePreset(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Directory string `json:"directory"`
		Preset    string `json:"preset"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.Directory == "" || req.Preset == "" {
		h.writeError(w, "both directory and preset are required", http.StatusBadRequest)
		return
	}

	result, err := h.service.ValidatePreset(r.Context(), req.Directory, req.Preset)
	if err != nil {
		h.writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.writeJSON(w, result)
}

// syncMods syncs mods from source directories to destination directories
func (h *ModsHandler) syncMods(w http.ResponseWriter, r *http.Request) {
	var req types.ModsSyncRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if len(req.From) == 0 || len(req.To) == 0 {
		h.writeError(w, "both from and to directories are required", http.StatusBadRequest)
		return
	}

	err := h.service.SyncMods(r.Context(), req.From, req.To)
	if err != nil {
		h.writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.writeJSON(w, map[string]string{"status": "success", "message": "mods synced successfully"})
}

// writeJSON writes a JSON response
func (h *ModsHandler) writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// writeError writes an error response
func (h *ModsHandler) writeError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	errorResp := &types.ErrorResponse{
		Error:   message,
		Code:    int32(code),
		Message: message,
	}

	json.NewEncoder(w).Encode(errorResp)
}
