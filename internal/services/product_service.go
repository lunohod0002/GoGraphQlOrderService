package services

import (
	"OzonOrderService/graph/model"
	"OzonOrderService/internal/repositories"
	"context"
)

type ProductService struct {
	productRepository repositories.ProductRepository
}

func NewProductService(repository *repositories.ProductRepository) *ProductService {
	return &ProductService{productRepository: *repository}
}
func (r *ProductService) Create(ctx context.Context, input model.ProductCreateInput) *model.Product {
	productInput := &model.Product{Name: input.Name, Price: input.Price}
	product, err := r.productRepository.Create(productInput)
	if err != nil {
		//
	}
	return product
}
