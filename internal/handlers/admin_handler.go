package handlers

import (
	"net/http"
	"order_agent/pkg/utils"
)

type AdminHandler struct{}

func NewAdminHandler() *AdminHandler {
	return &AdminHandler{}
}

func (h *AdminHandler) Dashboard(w http.ResponseWriter, r *http.Request) {
	utils.Success(w, map[string]string{
		"message": "Admin dashboard",
	})
}
