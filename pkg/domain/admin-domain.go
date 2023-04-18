package domain

import (
	"github.com/Bappy60/ecommerce_in_echo/pkg/models"
	"github.com/Bappy60/ecommerce_in_echo/pkg/types"
	"github.com/labstack/echo/v4"
)

type IAdminController interface {
	AddProduct(c echo.Context) error
	DeleteProduct(c echo.Context) error
	//TODO: need to implement
	//GetUser(c echo.Context)
	//SetUserRole(c echo.Context)
	//DeleteUser(c echo.Context)
}

type IAdminService interface {
AddProduct(product types.CreateProduct) error
DeleteProduct(id string) error
}

type IAdminRepo interface {
AddProduct(product models.Product) error
DeleteProduct(id uint64) error
}
