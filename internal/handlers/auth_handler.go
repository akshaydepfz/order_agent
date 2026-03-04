package handlers

import (
	"encoding/json"
	"net/http"
	"order_agent/internal/services"
	"order_agent/pkg/utils"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.Error(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// TODO: Validate credentials against user repo
	token, err := h.authService.GenerateToken(input.Email)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}
	utils.Success(w, map[string]string{"token": token})
}
