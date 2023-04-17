package services

import (
	"strconv"

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
func (service *CartService) AddToCart(userid uint64, reqStruct *types.AddToCartStruct) error {
	err := service.repo.AddToCart(userid, reqStruct)
	if err != nil {
		return err
	}
	return nil
}

// RemoveFromCart implements domain.ICartService
func (service *CartService) RemoveFromCart(cartItemId string, userId uint64) error {
	parsedcartItemId, err := strconv.ParseUint(cartItemId, 10, 64)
	if err != nil {
		return &types.CustomError{
			Message: "Invalid Cart item Id",
		}
	}
	err2 := service.repo.RemoveFromCart(parsedcartItemId, userId)
	if err2 != nil {
		return err2
	}
	return nil
}

// ShowCart implements domain.ICartService
func (service *CartService) ShowCart(userid uint64) ([]types.ShowCartStruct, error) {
	items, err := service.repo.ShowCart(userid)
	if err != nil {
		return nil, err
	}
	return items, nil
}
