package repository

import (
	"context"
	"database/sql"

	"order_agent/internal/models"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (id, shop_id, email, phone, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())`
	_, err := r.db.ExecContext(ctx, query, user.ID, user.ShopID, user.Email, user.Phone, user.Role)
	return err
}

func (r *UserRepo) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT id, shop_id, email, phone, role, created_at, updated_at FROM users WHERE email = $1`
	var user models.User
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID, &user.ShopID, &user.Email, &user.Phone, &user.Role, &user.CreatedAt, &user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}
