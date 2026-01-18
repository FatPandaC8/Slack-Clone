package http

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	auth *AuthService
}

func NewHandler(auth *AuthService) *Handler {
	return &Handler{
		auth: auth,
	}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name 		string `json:"name"`
		Email 		string `json:"email"`
		Password 	string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	userID, access, refresh, err := h.auth.Register(req.Name, req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"user_id":       userID,
		"access_token":  access,
		"refresh_token": refresh,
	})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	id, access, refresh, err := h.auth.Login(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"user_id":       id,
		"access_token":  access,
		"refresh_token": refresh,
	})
}