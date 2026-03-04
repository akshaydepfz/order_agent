package config

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"
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

func Load() (*Config, error) {
	// Local defaults (no env)
	port := "8080"
	dbHost := "localhost"
	dbUser := "postgres"
	dbPass := "postgres"
	dbName := "order_agent"
	dbPort := "5432"
	databaseURL := buildPostgresURL(dbHost, dbPort, dbUser, dbPass, dbName)

	jwtSecret := strings.TrimSpace(os.Getenv("JWT_SECRET"))
	whatsAppToken := strings.TrimSpace(os.Getenv("WHATSAPP_TOKEN"))
	s3Bucket := strings.TrimSpace(os.Getenv("S3_BUCKET"))
	s3Region := strings.TrimSpace(os.Getenv("S3_REGION"))
	awsKey := strings.TrimSpace(os.Getenv("AWS_ACCESS_KEY_ID"))
	awsSecret := strings.TrimSpace(os.Getenv("AWS_SECRET_ACCESS_KEY"))

	if jwtSecret == "" || whatsAppToken == "" || s3Bucket == "" || s3Region == "" || awsKey == "" || awsSecret == "" {
		return nil, errors.New("credentials missing: ensure Koyeb secrets JWT_SECRET, WHATSAPP_TOKEN, S3_BUCKET, S3_REGION, AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY are set")
	}

	cfg := &Config{
		Port:          port,
		DatabaseURL:   databaseURL,
		JWTSecret:     jwtSecret,
		WhatsAppToken: whatsAppToken,
		S3Bucket:      s3Bucket,
		S3Region:      s3Region,
		AWSKey:        awsKey,
		AWSSecret:     awsSecret,
	}

	return cfg, nil
}

func buildPostgresURL(host, port, user, password, dbname string) string {
	password = url.QueryEscape(password)
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
}
