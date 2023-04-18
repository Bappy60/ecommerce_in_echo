package services

import (
	"github.com/Bappy60/ecommerce_in_echo/pkg/domain"
	"github.com/Bappy60/ecommerce_in_echo/pkg/types"
)

type CartService struct {
	repo domain.ICartRepo
}

func CartServiceInstance(cartRepo domain.ICartRepo) domain.ICartService {
	return &CartService{
		repo: cartRepo,
	}
}

// AddToCart implements domain.ICartService
func (service *CartService) AddToCart(userid uint64, reqStruct *types.AddToCart) error {
	err := service.repo.AddToCart(userid, reqStruct)
	if err != nil {
		return err
	}
	return nil
}

// RemoveFromCart implements domain.ICartService
func (service *CartService) RemoveFromCart(cartItemId uint64, userId uint64) error {
	err := service.repo.RemoveFromCart(cartItemId, userId)
	if err != nil {
		return err
	}
	return nil
}

// ShowCart implements domain.ICartService
func (service *CartService) ShowCart(userid uint64) ([]types.ShowCart, error) {
	items, err := service.repo.ShowCart(userid)
	if err != nil {
		return nil, err
	}
	return items, nil
}
