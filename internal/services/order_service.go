package services

import (
	"context"

	"github.com/google/uuid"
	"order_agent/internal/models"
	"order_agent/internal/repository"
)

type OrderService struct {
	orderRepo *repository.OrderRepo
}

func NewOrderService(orderRepo *repository.OrderRepo) *OrderService {
	return &OrderService{orderRepo: orderRepo}
}

func (s *OrderService) CreateOrder(ctx context.Context, shopID uuid.UUID, customer, phone, items string, total float64) (*models.Order, error) {
	order := &models.Order{
		ID:       uuid.New(),
		ShopID:   shopID,
		Customer: customer,
		Phone:    phone,
		Items:    items,
		Status:   "pending",
		Total:    total,
	}
	return order, s.orderRepo.Create(ctx, order)
}

func (s *OrderService) GetOrdersByShop(ctx context.Context, shopID uuid.UUID) ([]*models.Order, error) {
	return s.orderRepo.GetByShopID(ctx, shopID)
}
