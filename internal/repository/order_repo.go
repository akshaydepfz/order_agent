package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"order_agent/internal/models"
)

type OrderRepo struct {
	db *sql.DB
}

func NewOrderRepo(db *sql.DB) *OrderRepo {
	return &OrderRepo{db: db}
}

func (r *OrderRepo) Create(ctx context.Context, order *models.Order) error {
	query := `INSERT INTO orders (id, shop_id, customer, phone, items, status, total, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())`
	_, err := r.db.ExecContext(ctx, query,
		order.ID, order.ShopID, order.Customer, order.Phone, order.Items, order.Status, order.Total,
	)
	return err
}

func (r *OrderRepo) GetByShopID(ctx context.Context, shopID uuid.UUID) ([]*models.Order, error) {
	query := `SELECT id, shop_id, customer, phone, items, status, total, created_at, updated_at
		FROM orders WHERE shop_id = $1 ORDER BY created_at DESC`
	rows, err := r.db.QueryContext(ctx, query, shopID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*models.Order
	for rows.Next() {
		var o models.Order
		if err := rows.Scan(&o.ID, &o.ShopID, &o.Customer, &o.Phone, &o.Items, &o.Status, &o.Total, &o.CreatedAt, &o.UpdatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, &o)
	}
	return orders, rows.Err()
}
