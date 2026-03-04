package handlers

import (
	"encoding/json"
	"net/http"
)

type WebhookHandler struct {
	// Add WhatsApp service, AI service, order service as needed
}

func NewWebhookHandler() *WebhookHandler {
	return &WebhookHandler{}
}

func (h *WebhookHandler) WhatsAppWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid method"})
		return
	}

	var payload map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON"})
		return
	}

	// TODO: Process WhatsApp webhook payload, extract message, call AI service, create order, send reply
	_ = payload

	w.WriteHeader(http.StatusOK)
}

func (h *WebhookHandler) Verify(w http.ResponseWriter, r *http.Request) {
	// WhatsApp webhook verification
	mode := r.URL.Query().Get("hub.mode")
	token := r.URL.Query().Get("hub.verify_token")
	challenge := r.URL.Query().Get("hub.challenge")

	if mode == "subscribe" && token != "" {
		w.Write([]byte(challenge))
		return
	}
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("Forbidden"))
}
