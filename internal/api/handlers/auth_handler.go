package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ruanv123/acme-hotel-api/internal/service"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

type registrationRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registrationResponse struct {
	User struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	}
	Error string `json:"error,omitempty"`
}

type emailRegistrationRequest struct {
	Email string `json:"email"`
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type authResponse struct {
	Token string `json:"token,omitempty"`
	Admin bool   `json:"admin"`
	Error string `json:"error,omitempty"`
}

type validateResponse struct {
	Validate string `json:"validate"`
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req registrationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.authService.Register(r.Context(), req.Email, req.Password, req.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := registrationResponse{}
	resp.User.ID = user.ID.String()
	resp.User.Email = user.Email

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
