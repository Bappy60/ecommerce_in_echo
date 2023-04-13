package controllers

import (
	"net/http"

	"github.com/Bappy60/ecommerce_in_echo/pkg/domain"
	"github.com/Bappy60/ecommerce_in_echo/pkg/types"
	"github.com/labstack/echo/v4"
)

type CartController struct {
	service domain.ICartService
}

func CartControllerInstance(CartService domain.ICartService) domain.ICartController {
	return &CartController{
		service: CartService,
	}
}

func (cartController *CartController) AddToCart(c echo.Context) error {
	reqStruct := types.AddToCartStruct{}
	if err := c.Bind(&reqStruct); err != nil {
		return err
	}
	//TODO: validate the productId and cartId

	userid := c.Get("userId").(uint64)
	err := cartController.service.AddToCart(userid, &reqStruct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Product added to cart successfully"})
}

func (cartController *CartController) ShowCart(c echo.Context) error {
	userid := c.Get("userId").(uint64)
	items, err := cartController.service.ShowCart(userid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, items)
}

func (cartController *CartController) RemoveFromCart(c echo.Context) error {
	cartItemId := c.Param("cartItemId")
	userid := c.Get("userId").(uint64)
	err := cartController.service.RemoveFromCart(cartItemId, userid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Cart item removed successfully"})
}
