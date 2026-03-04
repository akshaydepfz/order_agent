package main

import (
	"database/sql"
	"log"
	"net/http"

	"order_agent/internal/config"
	"order_agent/internal/database"
	"order_agent/internal/repository"
	"order_agent/internal/routes"
	"order_agent/internal/services"
	"order_agent/pkg/s3"
)

func main() {
	cfg := config.Load()

	mux := http.NewServeMux()

	var db *sql.DB
	if cfg.DatabaseURL != "" {
		var err error
		db, err = database.NewPostgres(cfg.DatabaseURL)
		if err != nil {
			log.Printf("Database connection failed: %v (running without DB)", err)
		} else {
			defer db.Close()
		}
	}

	s3Client, err := s3.NewClient(s3.Config{
		Bucket: cfg.S3Bucket,
		Region: cfg.S3Region,
		Key:    cfg.AWSKey,
		Secret: cfg.AWSSecret,
	})
	if err != nil {
		log.Printf("S3 client init failed: %v", err)
	}

	deps := &routes.Dependencies{
		JWTSecret:   cfg.JWTSecret,
		AuthService: services.NewAuthService(cfg.JWTSecret),
		ShopRepo:    repository.NewShopRepo(db),
		S3Client:    s3Client,
	}

	routes.Setup(mux, deps)

	log.Printf("Server running on :%s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, mux))
}
