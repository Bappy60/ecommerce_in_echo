package controllers

import (
	"net/http"
	"strconv"

	"github.com/Bappy60/ecommerce_in_echo/pkg/consts"
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
	reqStruct := types.AddToCart{}
	if err := c.Bind(&reqStruct); err != nil {
		return err
	}
	//TODO: validate the productId and cartId

	userid := c.Get("userId").(uint64)
	err := cartController.service.AddToCart(userid, &reqStruct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, consts.AddedToCart)
}

func (cartController *CartController) ShowCart(c echo.Context) error {
	userid := c.Get("userId").(uint64)
	items, err := cartController.service.ShowCart(userid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if len(items) == 0 {
		return c.JSON(http.StatusOK, consts.EmptyCart)
	}
	return c.JSON(http.StatusOK, items)
}

func (cartController *CartController) RemoveFromCart(c echo.Context) error {
	cartItemId := c.Param("cartItemId")
	userid := c.Get("userId").(uint64)
	parsedcartItemId, err := strconv.ParseUint(cartItemId, 10, 64)
	if err != nil {
		return &types.CustomError{
			Message: consts.ParseErr,
		}
	}
	if err := cartController.service.RemoveFromCart(parsedcartItemId, userid); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, consts.CartItemRemoved)
}
