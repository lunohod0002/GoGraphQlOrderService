package services

import (
	"OzonOrderService/graph/model"
	"OzonOrderService/internal/repositories"
	"context"
)

type CartService struct {
	cartRepository repositories.CartRepository
}

func NewCartService(repository *repositories.CartRepository) *CartService {
	return &CartService{cartRepository: *repository}
}
func (r *CartService) AddToCart(ctx context.Context, input *model.ItemAddInput) *model.Item {
	cart := r.cartRepository.Get(input.CartID)
	if cart == nil {
		//
	}
	item, err := r.cartRepository.AddItem(input)
	if err != nil {
		//
	}
	return item
}
func (r *CartService) AddCart(ctx context.Context, input *model.CartCreateInput) *model.Cart {
	cart, err := r.cartRepository.Create(input.UserID)

	if err != nil {
		//
	}
	return cart
}
