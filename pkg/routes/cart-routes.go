package routes

import (
	"github.com/Bappy60/ecommerce_in_echo/pkg/domain"
	"github.com/labstack/echo/v4"
)

func CartRoutes(e *echo.Echo, cartController domain.ICartController){
	cartGroup := e.Group("/auth")
	cartGroup.POST("/addtocart", cartController.AddToCart)
	cartGroup.POST("/removefromcart", cartController.RemoveFromCart)
	cartGroup.POST("/showcart", cartController.ShowCart)
}