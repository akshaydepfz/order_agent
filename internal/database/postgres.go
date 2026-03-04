package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func NewPostgres(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	if err := MigrateShops(db); err != nil {
		log.Printf("Migration warning: %v", err)
	}

	return db, nil
}

func MigrateShops(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS shops (
			id VARCHAR(36) PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			description TEXT DEFAULT '',
			logo_url TEXT DEFAULT '',
			owner_name VARCHAR(255) DEFAULT '',
			phone VARCHAR(50) NOT NULL,
			email VARCHAR(255) DEFAULT '',
			whatsapp_number VARCHAR(50) DEFAULT '',
			phone_number_id VARCHAR(100) DEFAULT '',
			access_token TEXT DEFAULT '',
			verify_token VARCHAR(255) DEFAULT '',
			address VARCHAR(500) DEFAULT '',
			city VARCHAR(100) DEFAULT '',
			state VARCHAR(100) DEFAULT '',
			pincode VARCHAR(20) DEFAULT '',
			latitude DOUBLE PRECISION DEFAULT 0,
			longitude DOUBLE PRECISION DEFAULT 0,
			opening_time VARCHAR(20) DEFAULT '',
			closing_time VARCHAR(20) DEFAULT '',
			delivery_radius_km INTEGER DEFAULT 0,
			delivery_fee DOUBLE PRECISION DEFAULT 0,
			minimum_order DOUBLE PRECISION DEFAULT 0,
			delivery_time_min INTEGER DEFAULT 0,
			cod_enabled BOOLEAN DEFAULT false,
			upi_enabled BOOLEAN DEFAULT false,
			ai_context TEXT DEFAULT '',
			selected_plan VARCHAR(50) DEFAULT '',
			plan_start_date TIMESTAMP DEFAULT NULL,
			plan_expire_date TIMESTAMP DEFAULT NULL,
			status VARCHAR(50) DEFAULT 'active',
			created_at TIMESTAMP DEFAULT NOW()
		)
	`)
	if err != nil {
		return fmt.Errorf("create shops table: %w", err)
	}

	// Add missing columns if table existed with old schema
	alters := []string{
		`ALTER TABLE shops ADD COLUMN IF NOT EXISTS description TEXT DEFAULT ''`,
		`ALTER TABLE shops ADD COLUMN IF NOT EXISTS logo_url TEXT DEFAULT ''`,
		`ALTER TABLE shops ADD COLUMN IF NOT EXISTS owner_name VARCHAR(255) DEFAULT ''`,
		`ALTER TABLE shops ADD COLUMN IF NOT EXISTS email VARCHAR(255) DEFAULT ''`,
		`ALTER TABLE shops ADD COLUMN IF NOT EXISTS whatsapp_number VARCHAR(50) DEFAULT ''`,
		`ALTER TABLE shops ADD COLUMN IF NOT EXISTS phone_number_id VARCHAR(100) DEFAULT ''`,
		`ALTER TABLE shops ADD COLUMN IF NOT EXISTS access_token TEXT DEFAULT ''`,
		`ALTER TABLE shops ADD COLUMN IF NOT EXISTS verify_token VARCHAR(255) DEFAULT ''`,
		`ALTER TABLE shops ADD COLUMN IF NOT EXISTS city VARCHAR(100) DEFAULT ''`,
		`ALTER TABLE shops ADD COLUMN IF NOT EXISTS state VARCHAR(100) DEFAULT ''`,
		`ALTER TABLE shops ADD COLUMN IF NOT EXISTS pincode VARCHAR(20) DEFAULT ''`,
		`ALTER TABLE shops ADD COLUMN IF NOT EXISTS latitude DOUBLE PRECISION DEFAULT 0`,
		`ALTER TABLE shops ADD COLUMN IF NOT EXISTS longitude DOUBLE PRECISION DEFAULT 0`,
		`ALTER TABLE shops ADD COLUMN IF NOT EXISTS opening_time VARCHAR(20) DEFAULT ''`,
		`ALTER TABLE shops ADD COLUMN IF NOT EXISTS closing_time VARCHAR(20) DEFAULT ''`,
		`ALTER TABLE shops ADD COLUMN IF NOT EXISTS delivery_radius_km INTEGER DEFAULT 0`,
		`ALTER TABLE shops ADD COLUMN IF NOT EXISTS delivery_fee DOUBLE PRECISION DEFAULT 0`,
		`ALTER TABLE shops ADD COLUMN IF NOT EXISTS minimum_order DOUBLE PRECISION DEFAULT 0`,
		`ALTER TABLE shops ADD COLUMN IF NOT EXISTS delivery_time_min INTEGER DEFAULT 0`,
		`ALTER TABLE shops ADD COLUMN IF NOT EXISTS cod_enabled BOOLEAN DEFAULT false`,
		`ALTER TABLE shops ADD COLUMN IF NOT EXISTS upi_enabled BOOLEAN DEFAULT false`,
		`ALTER TABLE shops ADD COLUMN IF NOT EXISTS ai_context TEXT DEFAULT ''`,
		`ALTER TABLE shops ADD COLUMN IF NOT EXISTS selected_plan VARCHAR(50) DEFAULT ''`,
		`ALTER TABLE shops ADD COLUMN IF NOT EXISTS plan_start_date TIMESTAMP DEFAULT NULL`,
		`ALTER TABLE shops ADD COLUMN IF NOT EXISTS plan_expire_date TIMESTAMP DEFAULT NULL`,
		`ALTER TABLE shops ADD COLUMN IF NOT EXISTS status VARCHAR(50) DEFAULT 'active'`,
	}
	for _, q := range alters {
		if _, err := db.Exec(q); err != nil {
			log.Printf("Migration alter warning: %v", err)
		}
	}

	return nil
}
