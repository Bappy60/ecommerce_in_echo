package repositories

import (

	"github.com/Bappy60/ecommerce_in_echo/pkg/domain"
	"github.com/Bappy60/ecommerce_in_echo/pkg/models"
	"github.com/Bappy60/ecommerce_in_echo/pkg/types"
	"github.com/jinzhu/gorm"
)

type CartRepo struct {
	db *gorm.DB
}

func CartDBInstance(d *gorm.DB) domain.ICartRepo {
	return &CartRepo{
		db: d,
	}
}

// AddToCart implements domain.ICartRepo
func (repo *CartRepo) AddToCart(userid uint64, reqStruct *types.AddToCart) error {

	cart, err := repo.getCartByUserID(userid)
	if err != nil {
		return &types.CustomError{
			Message: err.Error(),
		}
	}

	cartItem := models.CartItem{}
	err2 := repo.db.Where("cart_id = ? AND product_id = ?", cart.ID, reqStruct.ProductID).Preload("Product").First(&cartItem).Error
	if err2 == nil {
		cartItem.Quantity += reqStruct.Quantity
		cartItem.TotalItemPrice = float64(cartItem.Product.Price) * float64(cartItem.Quantity)
		if err := repo.db.Save(&cartItem).Error; err != nil {
			return &types.CustomError{
				Message: "could not update cart item",
				Err: err,
			}
		}
		return nil
	} else {
		product := models.Product{}
		if err := repo.db.Where("id = ?", reqStruct.ProductID).First(&product).Error; err == nil {
			newCartItem := models.CartItem{
				CartID:         cart.ID,
				ProductID:      uint64(reqStruct.ProductID),
				Quantity:       uint64(reqStruct.Quantity),
				TotalItemPrice: float64(product.Price) * float64(reqStruct.Quantity),
			}
			if err := repo.db.Create(&newCartItem).Error; err != nil {
				return &types.CustomError{
					Message: "could not add product to cart",
					Err: err,
				}
			}
			return nil
		}
		// TODO:  statuscode conflict
		return &types.CustomError{
			Message: "Invalid ProductId",
			Err: err,
		}
	}
}

// RemoveFromCart implements domain.ICartRepo
func (repo *CartRepo) RemoveFromCart(cartItemId uint64, userid uint64) error {
	cartItem := models.CartItem{}
	if err := repo.db.First(&cartItem, cartItemId).Error; err != nil {
		return &types.CustomError{
			Message: "Cart item not found",
		}
	}
	cart, err := repo.getCartByUserID(userid)
	if err != nil {
		return &types.CustomError{
			Message: err.Error(),
		}
	}
	if cartItem.CartID != cart.ID {
		return &types.CustomError{
			Message: "Cart item does not belong to the logged in user",
		}
	}
	if err := repo.db.Delete(&cartItem).Error; err != nil {
		return &types.CustomError{
			Message: "Could not remove cart item",
		}
	}
	return nil
}

// ShowCart implements domain.ICartRepo
func (repo *CartRepo) ShowCart(userId uint64) ([]types.ShowCart, error) {
	cart, err := repo.getCartByUserID(userId)
	if err != nil {
		return nil, &types.CustomError{
			Message: "Invalid ProductId",
		}
	}
	cart.CartItems, err = repo.PreloadCartItems(cart)
	if err != nil {
		return nil, &types.CustomError{
			Message: "could not load cart items",
		}
	}

	var items []types.ShowCart
	for _, item := range cart.CartItems {
		items = append(items, types.ShowCart{
			CartItemID:      item.ID,
			ProductName:     item.Product.Name,
			ProductCategory: item.Product.Category.CategoryName,
			ProductQuantity: item.Quantity,
			Amount:          item.TotalItemPrice,
		})
	}
	return items, nil
}

func (repo *CartRepo) getCartByUserID(userId uint64) (*models.Cart, error) {
	cart := &models.Cart{}
	query := repo.db.Model(&models.Cart{}).Preload("CartItems")
	if err := query.Where("user_id = ?", userId).First(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}
func (repo *CartRepo) PreloadCartItems(cart *models.Cart) ([]models.CartItem, error) {
	cart_items := []models.CartItem{}
	query := repo.db.Model(&cart_items).Preload("Product").Preload("Product.Category")
	if err := query.Where("cart_id = ?", cart.ID).Find(&cart_items).Error; err != nil {
		return nil, err
	}
	return cart_items, nil
}
