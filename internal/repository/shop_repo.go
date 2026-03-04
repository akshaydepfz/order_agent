package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"order_agent/internal/models"
)

var ErrDBNotConfigured = errors.New("database not configured")

type ShopRepo struct {
	db *sql.DB
}

func NewShopRepo(db *sql.DB) *ShopRepo {
	return &ShopRepo{db: db}
}

func (r *ShopRepo) Create(ctx context.Context, shop *models.Shop) error {
	if r.db == nil {
		return ErrDBNotConfigured
	}
	query := `INSERT INTO shops (
		id, name, description, logo_url, owner_name, phone, email,
		whatsapp_number, phone_number_id, access_token, verify_token,
		address, city, state, pincode, latitude, longitude,
		opening_time, closing_time, delivery_radius_km, delivery_fee, minimum_order, delivery_time_min,
		cod_enabled, upi_enabled, ai_context, selected_plan, plan_start_date, plan_expire_date,
		status, created_at
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17,
		$18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, NOW()
	)`
	_, err := r.db.ExecContext(ctx, query,
		shop.ID, shop.Name, shop.Description, shop.LogoURL, shop.OwnerName, shop.Phone, shop.Email,
		shop.WhatsAppNumber, shop.PhoneNumberID, shop.AccessToken, shop.VerifyToken,
		shop.Address, shop.City, shop.State, shop.Pincode, shop.Latitude, shop.Longitude,
		shop.OpeningTime, shop.ClosingTime, shop.DeliveryRadiusKM, shop.DeliveryFee, shop.MinimumOrder, shop.DeliveryTimeMin,
		shop.CODEnabled, shop.UPIEnabled, shop.AIContext, shop.SelectedPlan,
		nullTime(shop.PlanStartDate), nullTime(shop.PlanExpireDate),
		shop.Status,
	)
	return err
}

func nullTime(t time.Time) interface{} {
	if t.IsZero() {
		return nil
	}
	return t
}

func (r *ShopRepo) List(ctx context.Context) ([]*models.Shop, error) {
	if r.db == nil {
		return nil, nil
	}
	query := `SELECT id, name, description, logo_url, owner_name, phone, email,
		whatsapp_number, phone_number_id, access_token, verify_token,
		address, city, state, pincode, latitude, longitude,
		opening_time, closing_time, delivery_radius_km, delivery_fee, minimum_order, delivery_time_min,
		cod_enabled, upi_enabled, ai_context, selected_plan, plan_start_date, plan_expire_date,
		status, created_at FROM shops ORDER BY created_at DESC`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shops []*models.Shop
	for rows.Next() {
		shop, err := scanShop(rows)
		if err != nil {
			return nil, err
		}
		shops = append(shops, shop)
	}
	return shops, rows.Err()
}

func (r *ShopRepo) GetByID(ctx context.Context, id string) (*models.Shop, error) {
	if r.db == nil {
		return nil, ErrDBNotConfigured
	}
	query := `SELECT id, name, description, logo_url, owner_name, phone, email,
		whatsapp_number, phone_number_id, access_token, verify_token,
		address, city, state, pincode, latitude, longitude,
		opening_time, closing_time, delivery_radius_km, delivery_fee, minimum_order, delivery_time_min,
		cod_enabled, upi_enabled, ai_context, selected_plan, plan_start_date, plan_expire_date,
		status, created_at FROM shops WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)
	return scanShopRow(row)
}

func scanShop(rows *sql.Rows) (*models.Shop, error) {
	var shop models.Shop
	var planStart, planExpire sql.NullTime
	err := rows.Scan(
		&shop.ID, &shop.Name, &shop.Description, &shop.LogoURL, &shop.OwnerName, &shop.Phone, &shop.Email,
		&shop.WhatsAppNumber, &shop.PhoneNumberID, &shop.AccessToken, &shop.VerifyToken,
		&shop.Address, &shop.City, &shop.State, &shop.Pincode, &shop.Latitude, &shop.Longitude,
		&shop.OpeningTime, &shop.ClosingTime, &shop.DeliveryRadiusKM, &shop.DeliveryFee, &shop.MinimumOrder, &shop.DeliveryTimeMin,
		&shop.CODEnabled, &shop.UPIEnabled, &shop.AIContext, &shop.SelectedPlan,
		&planStart, &planExpire,
		&shop.Status, &shop.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	if planStart.Valid {
		shop.PlanStartDate = planStart.Time
	}
	if planExpire.Valid {
		shop.PlanExpireDate = planExpire.Time
	}
	return &shop, nil
}

func scanShopRow(row *sql.Row) (*models.Shop, error) {
	var shop models.Shop
	var planStart, planExpire sql.NullTime
	err := row.Scan(
		&shop.ID, &shop.Name, &shop.Description, &shop.LogoURL, &shop.OwnerName, &shop.Phone, &shop.Email,
		&shop.WhatsAppNumber, &shop.PhoneNumberID, &shop.AccessToken, &shop.VerifyToken,
		&shop.Address, &shop.City, &shop.State, &shop.Pincode, &shop.Latitude, &shop.Longitude,
		&shop.OpeningTime, &shop.ClosingTime, &shop.DeliveryRadiusKM, &shop.DeliveryFee, &shop.MinimumOrder, &shop.DeliveryTimeMin,
		&shop.CODEnabled, &shop.UPIEnabled, &shop.AIContext, &shop.SelectedPlan,
		&planStart, &planExpire,
		&shop.Status, &shop.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if planStart.Valid {
		shop.PlanStartDate = planStart.Time
	}
	if planExpire.Valid {
		shop.PlanExpireDate = planExpire.Time
	}
	return &shop, nil
}
