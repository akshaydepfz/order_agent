package services

import (
	"order_agent/pkg/utils"
)

type AuthService struct {
	jwtSecret string
}

func NewAuthService(jwtSecret string) *AuthService {
	return &AuthService{jwtSecret: jwtSecret}
}

func (s *AuthService) GenerateToken(userID string) (string, error) {
	return utils.GenerateJWT(userID, s.jwtSecret)
}

func (s *AuthService) ValidateToken(token string) (string, error) {
	return utils.ValidateJWT(token, s.jwtSecret)
}
