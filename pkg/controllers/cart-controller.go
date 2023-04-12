package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type CartController struct {
	db *gorm.DB
}

func CartControllerInstance(db *gorm.DB) *CartController {
	return &CartController{
		db: db,
	}
}

func (cartController *CartController) AddToCart(c echo.Context) error {
	panic("unimplemented")
}
func (cartController *CartController) RemoveFromCart(c echo.Context) error {
	panic("unimplemented")
}
func (cartController *CartController) ShowCart(c echo.Context) error {
	panic("unimplemented")
}
