package types

import (
	"fmt"

	"github.com/Bappy60/ecommerce_in_echo/pkg/models"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type SignInRequest struct {
	Name     string         `json:"name"`
	Password string         `json:"password"`
	Email    string         `json:"email"`
	Address  models.Address `json:"address"`
}

func (user SignInRequest) Validate() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&user.Password, validation.Required),
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Address),
	)
}

type LoginRequset struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (user LoginRequset) Validate() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Password, validation.Required, validation.Length(4, 20)),
		validation.Field(&user.Email, validation.Required, is.Email),
	)
}

type CreateProduct struct {
	Name     string                 `json:"name"`
	Price    float64                `json:"price"`
	Category models.ProductCategory `json:"category"`
}

func (product CreateProduct) Validate() error {
	return validation.ValidateStruct(&product,
		validation.Field(&product.Name, validation.Required),
		validation.Field(&product.Price, validation.Required),
		validation.Field(&product.Category),
	)
}

type AddToCart struct {
	ProductID int    `json:"product_id"`
	Quantity  uint64 `json:"quantity"`
}

func (atc AddToCart) Validate() error {
	return validation.ValidateStruct(&atc,
		validation.Field(&atc.ProductID, validation.Required),
		validation.Field(&atc.Quantity, validation.Required),
	)
}

type ShowCart struct {
	CartItemID      uint64  `json:"cart_item_id"`
	ProductName     string  `json:"product_name"`
	ProductCategory string  `json:"product_category"`
	ProductQuantity uint64  `json:"product_quantity"`
	Amount          float64 `json:"amount"`
}

type SearchRequest struct {
	Id       uint64  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}
type SearchRepo struct {
	Id       uint64  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

type CustomError struct {
	Message string
	Err     error
}

func (e *CustomError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}
