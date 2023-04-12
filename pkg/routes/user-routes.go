package routes

import (
	"github.com/Bappy60/ecommerce_in_echo/pkg/domain"
	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo, userController domain.IUserController ) {
	userGroup := e.Group("/users")
	userGroup.POST("/signup",userController.SignUp)
	userGroup.POST("/login",userController.Login)
	userGroup.GET("users/viewproduct",userController.ViewProduct)
	userGroup.GET("users/search",userController.SearchProduct)
}