package routes

import (
	"github.com/Bappy60/ecommerce_in_echo/pkg/domain"
	"github.com/labstack/echo/v4"
)

func GeneralRoutes(e *echo.Echo, generalController domain.IGeneralController ) {
	
	e.GET("/search",generalController.SearchProduct)
}