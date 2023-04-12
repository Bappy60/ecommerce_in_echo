package controllers

import (
	"net/http"

	"github.com/Bappy60/ecommerce_in_echo/pkg/models"
	"github.com/Bappy60/ecommerce_in_echo/pkg/types"
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
	req := types.AddToCartStruct{}
	if err := c.Bind(&req); err != nil {
		return err
	}
	//TODO: validate the productId and cartId

	userid := c.Get("userId").(uint64)
	cart, err := cartController.getCartIDByUserID(userid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	cartItem := models.CartItem{}
	if err := cartController.db.Where("cart_id = ? AND product_id = ?", cart.ID, req.ProductID).First(&cartItem).Error; err == nil {
		cartItem.Quantity += uint64(req.Quantity)
		cartItem.TotalItemPrice += float64(cartItem.Product.Price) * float64(cartItem.Quantity)
		if err := cartController.db.Save(&cartItem).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, "could not update cart item")
		}
	} else {
		product := models.Product{}
		if err := cartController.db.Where("id = ?", req.ProductID).First(&product).Error; err == nil {
			newCartItem := models.CartItem{
				CartID:         cart.ID,
				ProductID:      uint64(req.ProductID),
				Quantity:       uint64(req.Quantity),
				TotalItemPrice: float64(product.Price) * float64(req.Quantity),
			}
			if err := cartController.db.Create(&newCartItem).Error; err != nil {
				return c.JSON(http.StatusInternalServerError, "could not add product to cart")
			}
			return c.JSON(http.StatusOK, map[string]string{"message": "Product added to cart successfully"})
		}
		return c.JSON(http.StatusInternalServerError, "Invalid ProductId")
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Product added to cart successfully"})
}

func (cartController *CartController) RemoveFromCart(c echo.Context) error {
	panic("unimplemented")
}
func (cartController *CartController) ShowCart(c echo.Context) error {
	panic("unimplemented")
}

func (cartController *CartController) getCartIDByUserID(userId uint64) (*models.Cart, error) {
	cart := &models.Cart{}
	if err := cartController.db.Where("user_id = ?", userId).First(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}
