package routes

import (
	"github.com/Bappy60/ecommerce_in_echo/pkg/domain"
	"github.com/labstack/echo/v4"
)

func AdminRoutes(e *echo.Echo, adminController domain.IAdminController) {
	AdminGroup := e.Group("/admin")

	AdminGroup.POST("/addproduct",adminController.AddProduct)
	AdminGroup.DELETE("/deleteproduct/:id",adminController.DeleteProduct)

}
