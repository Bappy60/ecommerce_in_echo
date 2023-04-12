package domain

import "github.com/labstack/echo/v4"

type IAdminController interface {
	AddProduct(c echo.Context) error
	DeleteProduct(c echo.Context) error
}

// type IAdminService interface {
//AddProduct(c echo.Context) error
// DeleteProduct(c echo.Context) error
// }

// type IAdminRepo interface {
//AddProduct(c echo.Context) error
// DeleteProduct(c echo.Context) error
// }
