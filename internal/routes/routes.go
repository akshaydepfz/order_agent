package routes

import (
	"encoding/json"
	"net/http"

	"order_agent/internal/handlers"
	"order_agent/internal/middleware"
	"order_agent/internal/repository"
	"order_agent/internal/services"
	"order_agent/pkg/s3"
)

type Dependencies struct {
	JWTSecret   string
	AuthService *services.AuthService
	ShopRepo    *repository.ShopRepo
	S3Client    *s3.Client
}

func Setup(mux *http.ServeMux, deps *Dependencies) {
	// Root
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Order Bot Running 🚀"})
	})

	// Webhook (WhatsApp) - no auth
	webhookHandler := handlers.NewWebhookHandler()
	mux.HandleFunc("GET /webhook/whatsapp", webhookHandler.Verify)
	mux.HandleFunc("POST /webhook/whatsapp", webhookHandler.WhatsAppWebhook)

	// Auth
	authHandler := handlers.NewAuthHandler(deps.AuthService)
	mux.HandleFunc("POST /auth/login", authHandler.Login)

	// Protected routes
	adminHandler := handlers.NewAdminHandler()
	mux.Handle("GET /api/admin/dashboard", middleware.AuthRequired(deps.JWTSecret, http.HandlerFunc(adminHandler.Dashboard)))

	shopHandler := handlers.NewShopHandler(deps.ShopRepo, deps.S3Client)
	mux.Handle("GET /api/shops", middleware.AuthRequired(deps.JWTSecret, http.HandlerFunc(shopHandler.List)))
	mux.Handle("POST /api/shops", middleware.AuthRequired(deps.JWTSecret, http.HandlerFunc(shopHandler.Create)))
	mux.Handle("GET /api/shops/{id}", middleware.AuthRequired(deps.JWTSecret, http.HandlerFunc(shopHandler.GetByID)))
}
