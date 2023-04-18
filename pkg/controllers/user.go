package controllers

import (
	"net/http"

	"github.com/Bappy60/ecommerce_in_echo/pkg/domain"
	"github.com/Bappy60/ecommerce_in_echo/pkg/types"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	service domain.IUserService
}

func UserControllerInstance(userService domain.IUserService) domain.IUserController {
	return &UserController{
		service: userService,
	}
}

// SignUp implements domain.IUserController
func (userController *UserController) SignUp(c echo.Context) error {
	user := types.SignInRequest{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err := user.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := userController.service.SignUp(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, "Successfully Signed Up!!")
}

// Login implements domain.IUserController
func (userController *UserController) Login(c echo.Context) error {
	user := types.LoginRequset{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err := user.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	token, err := userController.service.Login(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "You were logged in!",
		"token":   token,
	})

}
