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
func (r *CartService) AddCart(ctx context.Context, input *model.ItemAddInput) *model.Item {
	item, err := r.cartRepository.AddItem(input)
	if err != nil {
		//
	}
	return item
}
