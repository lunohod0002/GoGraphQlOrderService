package services

import (
	"OzonOrderService/graph/model"
	"OzonOrderService/internal/repositories"
	"context"
)

type OrderService struct {
	orderRepository repositories.OrderRepository
	cartRepository  repositories.CartRepository
}

func NewOrderService(orderRepository *repositories.OrderRepository, cartRepository *repositories.CartRepository) *OrderService {
	return &OrderService{orderRepository: *orderRepository,
		cartRepository: *cartRepository}
}
func (r *OrderService) CreateOrder(ctx context.Context, input *model.OrderCreateInput) *model.Order {
	cart := r.cartRepository.Get(input.CartID)
	order, err := r.orderRepository.Create(cart, input.Name)
	if err != nil {
		//
	}
	return order
}
func (r *OrderService) GetAll(ctx context.Context, user_id int) []*model.Order {
	orders, err := r.orderRepository.GetAll(user_id)
	if err != nil {
		//
	}
	return orders
}
