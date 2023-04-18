package domain

import (
	"github.com/Bappy60/ecommerce_in_echo/pkg/types"
	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
}

type IUserService interface {
	SignUp(reqUser *types.SignInRequest) error
	Login(reqUser *types.LoginRequset) (string,error)
}
type IUserRepo interface {
	SignUp(user *types.SignInRequest) error
	Login(user *types.LoginRequset) (string,error)
}
