package domain

import (
	"github.com/labstack/echo/v4"
)

type ICartController interface {
	AddToCart(c echo.Context) error
	RemoveFromCart(c echo.Context)error
	ShowCart(c echo.Context)error
}

// type ICartService interface {
// 	AddToCart(c echo.Context) error
// }
// type ICartRepo interface {
// 	AddToCart(c echo.Context) error
// }
