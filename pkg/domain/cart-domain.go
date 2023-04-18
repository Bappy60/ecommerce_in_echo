package domain

import (
	"github.com/Bappy60/ecommerce_in_echo/pkg/types"
	"github.com/labstack/echo/v4"
)

type ICartController interface {
	AddToCart(c echo.Context) error
	RemoveFromCart(c echo.Context)error
	ShowCart(c echo.Context)error
	
}

type ICartService interface {
	AddToCart(userid uint64, reqStruct *types.AddToCart) error
	RemoveFromCart(cartItemId string,userId uint64)error
	ShowCart(userid uint64) ([]types.ShowCart,error)
}
type ICartRepo interface {
	AddToCart(userid uint64, reqStruct *types.AddToCart) error
	RemoveFromCart(cartItemId uint64, userId uint64)error
	ShowCart(userid uint64)([]types.ShowCart,error)
}
