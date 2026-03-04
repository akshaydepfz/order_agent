package config

import (
	"fmt"
	"net/url"
	"os"
)

type Config struct {
	Port          string
	DatabaseURL   string
	JWTSecret     string
	WhatsAppToken string

	// AWS S3
	S3Bucket  string
	S3Region  string
	AWSKey    string
	AWSSecret string
}

func Load() *Config {
	cfg := &Config{
		Port:          getEnv("PORT", "8080"),
		DatabaseURL:   getEnv("DATABASE_URL", ""),
		JWTSecret:     getEnv("JWT_SECRET", "your-256-bit-secret"),
		WhatsAppToken: getEnv("WHATSAPP_TOKEN", ""),

		S3Bucket:  getEnv("S3_BUCKET", "oryoo-bucket"),
		S3Region:  getEnv("S3_REGION", "ap-southeast-2"),
		AWSKey:    getEnv("AWS_ACCESS_KEY_ID", ""),
		AWSSecret: getEnv("AWS_SECRET_ACCESS_KEY", ""),
	}

	// Build DATABASE_URL from individual credentials if not set
	if cfg.DatabaseURL == "" {
		host := getEnv("DATABASE_HOST", "")
		user := getEnv("DATABASE_USER", "")
		pass := getEnv("DATABASE_PASSWORD", "")
		name := getEnv("DATABASE_NAME", "")
		port := getEnv("DATABASE_PORT", "5432")
		if host != "" && user != "" && name != "" {
			cfg.DatabaseURL = buildPostgresURL(host, port, user, pass, name)
		}
	}

	return cfg
}

func buildPostgresURL(host, port, user, password, dbname string) string {
	password = url.QueryEscape(password)
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=require", user, password, host, port, dbname)
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
