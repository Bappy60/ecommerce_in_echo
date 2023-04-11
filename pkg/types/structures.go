package types

import (
	"github.com/Bappy60/ecommerce_in_echo/pkg/models"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type RequestStruct struct {
	Name     string         `json:"name"`
	Password string         `json:"password"`
	Email    string         `json:"email"`
	Address  models.Address `json:"address"`
}
func (user RequestStruct) Validate() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&user.Password, validation.Required),
		validation.Field(&user.Email, validation.Required,is.Email),
		validation.Field(&user.Address),
	)
}

type LogReqStruct struct {
	Password string         `json:"password"`
	Email    string         `json:"email"`
}
func (logUser LogReqStruct) Validate() error {
	return validation.ValidateStruct(&logUser,
		validation.Field(&logUser.Password, validation.Required,validation.Length(4,20)),
		validation.Field(&logUser.Email, validation.Required,is.Email),
		
	)
}