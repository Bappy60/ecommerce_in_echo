package routes

import (
	"github.com/Bappy60/ecommerce_in_echo/pkg/domain"
	"github.com/labstack/echo/v4"
	"github.com/Bappy60/ecommerce_in_echo/pkg/middleware"
)

func CartRoutes(e *echo.Echo, cartController domain.ICartController){
	cartGroup := e.Group("/auth")
	cartGroup.Use(middleware.JWTMiddleware)
	cartGroup.POST("/addtocart", cartController.AddToCart)
	cartGroup.POST("/showcart", cartController.ShowCart)
	cartGroup.POST("/removefromcart/:cartItemId", cartController.RemoveFromCart)
}