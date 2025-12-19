package services

import (
	"OzonOrderService/graph/model"
	"OzonOrderService/internal/repositories"
	"context"
	"fmt"
)

type CartService struct {
	cartRepository repositories.CartRepository
	itemRepository repositories.ItemRepository
}

func NewCartService(
	cartRepository *repositories.CartRepository,
	itemRepository *repositories.ItemRepository) *CartService {
	return &CartService{cartRepository: *cartRepository,
		itemRepository: *itemRepository}
}
func (r *CartService) AddToCart(ctx context.Context, input *model.ItemUpdateInput) (*model.Item, error) {
	cart := r.cartRepository.Get(input.CartID)
	if cart == nil {
		return nil, fmt.Errorf("Не удалось выполнить запрос: Корзины не существует")

	}
	item, err := r.itemRepository.GetItem(input.CartID, input.ProductID)

	if err != nil {
		//
	}
	if item == nil {
		newItem, _ := r.itemRepository.AddItem(input)
		return newItem, nil

	} else {
		r.itemRepository.UpdateItem(item.ID, input.ProductID)
	}
	return item, nil
}
func (r *CartService) RemoveFromCart(ctx context.Context, input *model.ItemUpdateInput) (*model.Item, error) {
	cart := r.cartRepository.Get(input.CartID)
	if cart == nil {
		return nil, fmt.Errorf("Не удалось выполнить запрос: Корзины не существует")

	}
	item, err := r.itemRepository.GetItem(input.CartID, input.ProductID)

	if err != nil {
		return nil, fmt.Errorf("Не удалось выполнить запрос")

	}
	if item == nil {
		return nil, fmt.Errorf("Нельзя удалить товар, котого нет в корзине")
	} else if item.Quantity < input.Quantity {
		return nil, fmt.Errorf("Нельзя удалить больше товара, чем есть в корзине")
	} else if item.Quantity == input.Quantity {
		r.itemRepository.DeleteItem(item.ID)
		return nil, nil
	}
	r.itemRepository.UpdateItem(item.ID, int(item.Quantity-input.Quantity))

	return item, nil
}
func (r *CartService) AddCart(ctx context.Context, input *model.CartCreateInput) *model.Cart {
	cart, err := r.cartRepository.Create(input.UserID)

	if err != nil {
		//
	}
	return cart
}
