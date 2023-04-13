package domain

import "github.com/labstack/echo/v4"

type IAdminController interface {
	AddProduct(c echo.Context) error
	DeleteProduct(c echo.Context) error
	//TODO: need to implement
	//GetUser(c echo.Context)
	//SetUserRole(c echo.Context)
	//DeleteUser(c echo.Context)
}

// type IAdminService interface {
//AddProduct(c echo.Context) error
// DeleteProduct(c echo.Context) error
// }

// type IAdminRepo interface {
//AddProduct(c echo.Context) error
// DeleteProduct(c echo.Context) error
// }
