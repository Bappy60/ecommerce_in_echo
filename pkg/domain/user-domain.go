package domain

import (
	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	ViewProduct(c echo.Context) error
	SearchProduct(c echo.Context) error
}

// type IUserService interface {
// 	SignUp(user *types.RequestStruct)
// 	Login()
// 	ViewProduct()
// 	SearchProduct()
// }
// type IUserRepo interface {
// 	SignUp()
// 	Login()
// 	ViewProduct()
// 	SearchProduct()
// }
