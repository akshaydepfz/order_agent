package models

import (
	"time"
)

type Shop struct {
	ID string `json:"id" db:"id"`

	// Basic Info
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	LogoURL     string `json:"logo_url" db:"logo_url"`

	// Owner Info
	OwnerName string `json:"owner_name" db:"owner_name"`
	Phone     string `json:"phone" db:"phone"`
	Email     string `json:"email" db:"email"`

	// WhatsApp Integration
	WhatsAppNumber string `json:"whatsapp_number" db:"whatsapp_number"`
	PhoneNumberID  string `json:"phone_number_id" db:"phone_number_id"`
	AccessToken    string `json:"access_token" db:"access_token"`
	VerifyToken    string `json:"verify_token" db:"verify_token"`

	// Location
	Address   string  `json:"address" db:"address"`
	City      string  `json:"city" db:"city"`
	State     string  `json:"state" db:"state"`
	Pincode   string  `json:"pincode" db:"pincode"`
	Latitude  float64 `json:"latitude" db:"latitude"`
	Longitude float64 `json:"longitude" db:"longitude"`

	// Business Hours
	OpeningTime string `json:"opening_time" db:"opening_time"`
	ClosingTime string `json:"closing_time" db:"closing_time"`

	// Delivery Settings
	DeliveryRadiusKM int     `json:"delivery_radius_km" db:"delivery_radius_km"`
	DeliveryFee      float64 `json:"delivery_fee" db:"delivery_fee"`
	MinimumOrder     float64 `json:"minimum_order" db:"minimum_order"`
	DeliveryTimeMin  int     `json:"delivery_time_min" db:"delivery_time_min"`

	// Payment Options
	CODEnabled bool `json:"cod_enabled" db:"cod_enabled"`
	UPIEnabled bool `json:"upi_enabled" db:"upi_enabled"`

	// AI Context
	AIContext string `json:"ai_context" db:"ai_context"`

	// Subscription
	SelectedPlan   string    `json:"selected_plan" db:"selected_plan"`
	PlanStartDate  time.Time `json:"plan_start_date" db:"plan_start_date"`
	PlanExpireDate time.Time `json:"plan_expire_date" db:"plan_expire_date"`

	// Status
	Status string `json:"status" db:"status"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
