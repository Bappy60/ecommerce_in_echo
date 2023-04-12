package types

import (
	"github.com/Bappy60/ecommerce_in_echo/pkg/models"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type SignReqStruct struct {
	Name     string         `json:"name"`
	Password string         `json:"password"`
	Email    string         `json:"email"`
	Address  models.Address `json:"address"`
	IsAdmin  bool			`json:"is_admin"`
}

func (user SignReqStruct) Validate() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&user.Password, validation.Required),
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Address),
	)
}

type LoginReqStruct struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (user LoginReqStruct) Validate() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Password, validation.Required, validation.Length(4, 20)),
		validation.Field(&user.Email, validation.Required, is.Email),
	)
}

type CreateProductStruct struct {
	Name     string                 `json:"name"`
	Price    uint64                 `json:"price"`
	Category models.ProductCategory `json:"category"`
}
func (product CreateProductStruct) Validate() error {

    // categoryRules := []*validation.FieldRules{
    //     validation.Field(&product.Category.CategoryName, validation.Required),
    //     validation.Field(&product.Category.Description, validation.Required),
    // }

    return validation.ValidateStruct(&product,
        validation.Field(&product.Name, validation.Required),
        validation.Field(&product.Price, validation.Required),
        validation.Field(&product.Category),
    )
}

// func categoryRulesValidator(rules []*validation.FieldRules) validation.RuleFunc {
// 	return func(value interface{}) error {
// 		return validation.ValidateStruct(nil, rules...)
// 	}
// }