package controllers

import (
	"net/http"
	"strconv"

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

	if err := cartController.db.Where("cart_id = ? AND product_id = ?", cart.ID, req.ProductID).Preload("Product").First(&cartItem).Error; err == nil {
		// if req.Quantity < 0 && int(cartItem.Quantity) < req.Quantity {
		// 	//TODO:delete the cartItem
		// }
		cartItem.Quantity += req.Quantity
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

func (cartController *CartController) ShowCart(c echo.Context) error {
	userid := c.Get("userId").(uint64)
	cart, err := cartController.getCartIDByUserID(userid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	cart.CartItems, err = cartController.PreloadCartItems(cart)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	var items []types.ShowCartStruct
	for _, item := range cart.CartItems {
		items = append(items, types.ShowCartStruct{
			CartItemID:      item.ID,
			ProductName:     item.Product.Name,
			ProductCategory: item.Product.Category.CategoryName,
			ProductQuantity: item.Quantity,
			Amount:          item.TotalItemPrice,
		})
	}
	return c.JSON(http.StatusOK, items)
}

func (cartController *CartController) RemoveFromCart(c echo.Context) error {
	cartItemId, err := strconv.ParseUint(c.Param("cartItemId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid cart item ID")
	}

	// Get the cart item from the database
	cartItem := models.CartItem{}
	if err := cartController.db.First(&cartItem, cartItemId).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Cart item not found")
	}

	// Check that the logged in user owns the cart that the item belongs to
	userid := c.Get("userId").(uint64)
	cart, err := cartController.getCartIDByUserID(userid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if cartItem.CartID != cart.ID {
		return c.JSON(http.StatusForbidden, "Cart item does not belong to the logged in user")
	}

	// Delete the cart item from the database
	if err := cartController.db.Delete(&cartItem).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Could not remove cart item")
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Cart item removed successfully"})
}

func (cartController *CartController) getCartIDByUserID(userId uint64) (*models.Cart, error) {
	cart := &models.Cart{}
	query := cartController.db.Model(&models.Cart{}).Preload("CartItems")
	if err := query.Where("user_id = ?", userId).First(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}
func (cartController *CartController) PreloadCartItems(cart *models.Cart) ([]models.CartItem, error) {
	cart_items := []models.CartItem{}
	query := cartController.db.Model(&cart_items).Preload("Product").Preload("Product.Category")
	if err := query.Where("cart_id = ?", cart.ID).Find(&cart_items).Error; err != nil {
		return nil, err
	}
	return cart_items, nil
}
