package services

import (
	"OzonOrderService/graph/model"
	"OzonOrderService/internal/repositories"
	"context"
)

type UserService struct {
	userRepository repositories.UserRepository
	cartRepository repositories.CartRepository
}

func NewUserService(cartRepo *repositories.CartRepository, repository *repositories.UserRepository) *UserService {
	return &UserService{userRepository: *repository}
}
func (r *UserService) Create(ctx context.Context, input model.UserCreateInput) *model.User {
	userInput := &model.User{Fio: input.Fio, Balance: input.Balance}
	user, err := r.userRepository.Create(userInput)
	if err != nil {
		//
	}
	return user
}
