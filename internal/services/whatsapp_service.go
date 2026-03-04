package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type WhatsAppService struct {
	token string
}

func NewWhatsAppService(token string) *WhatsAppService {
	return &WhatsAppService{token: token}
}

func (s *WhatsAppService) SendMessage(to, message string) error {
	url := fmt.Sprintf("https://graph.facebook.com/v18.0/messages")
	payload := map[string]interface{}{
		"messaging_product": "whatsapp",
		"to":                to,
		"type":              "text",
		"text":              map[string]string{"body": message},
	}
	body, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("whatsapp API error: %d", resp.StatusCode)
	}
	return nil
}
