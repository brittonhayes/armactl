package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/brittonhayes/armactl/web/api/types"
	"github.com/brittonhayes/armactl/web/services"
	"github.com/go-chi/chi/v5"
)

// KeysHandler handles key management related requests
type KeysHandler struct {
	service services.KeysService
}

// NewKeysHandler creates a new keys handler
func NewKeysHandler(service services.KeysService) *KeysHandler {
	return &KeysHandler{
		service: service,
	}
}

// Routes sets up the keys routes
func (h *KeysHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.listKeys)
	r.Post("/copy", h.copyKeys)

	return r
}

// listKeys lists all .bikey files in a directory
func (h *KeysHandler) listKeys(w http.ResponseWriter, r *http.Request) {
	directory := r.URL.Query().Get("directory")
	if directory == "" {
		h.writeError(w, "directory parameter is required", http.StatusBadRequest)
		return
	}

	keys, err := h.service.ListKeys(r.Context(), directory)
	if err != nil {
		h.writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := types.KeysListResponse{
		Keys: keys,
	}

	h.writeJSON(w, response)
}

// copyKeys copies keys from a source directory to a destination directory
func (h *KeysHandler) copyKeys(w http.ResponseWriter, r *http.Request) {
	var req struct {
		From string `json:"from"`
		To   string `json:"to"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.From == "" || req.To == "" {
		h.writeError(w, "both from and to directories are required", http.StatusBadRequest)
		return
	}

	err := h.service.CopyKeys(r.Context(), req.From, req.To)
	if err != nil {
		h.writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.writeJSON(w, map[string]string{"status": "success", "message": "keys copied successfully"})
}

// writeJSON writes a JSON response
func (h *KeysHandler) writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// writeError writes an error response
func (h *KeysHandler) writeError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	errorResp := &types.ErrorResponse{
		Error:   message,
		Code:    int32(code),
		Message: message,
	}

	json.NewEncoder(w).Encode(errorResp)
}
