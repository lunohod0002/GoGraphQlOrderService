package graph

import (
	"OzonOrderService/internal/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ProductService *services.ProductService
	UserService    *services.UserService
	CartService    *services.CartService
}

func NewResolver(ProductService *services.ProductService, CartService *services.CartService, UserService *services.UserService) *Resolver {
	return &Resolver{
		ProductService: ProductService,
		UserService:    UserService,
		CartService:    CartService}
}
