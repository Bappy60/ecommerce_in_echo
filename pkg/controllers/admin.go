package controllers

import (
	"net/http"

	"github.com/Bappy60/ecommerce_in_echo/pkg/domain"
	"github.com/Bappy60/ecommerce_in_echo/pkg/types"
	"github.com/labstack/echo/v4"
)

type AdminController struct {
	service domain.IAdminService
}

func AdminControllerInstance(adminService domain.IAdminService) domain.IAdminController {
	return &AdminController{
		service: adminService,
	}
}

func (adminController *AdminController) AddProduct(c echo.Context) error {
	reqproduct := types.CreateProduct{}
	if err := c.Bind(&reqproduct); err != nil {
		return c.JSON(http.StatusBadRequest, "err while binding")
	}
	err := reqproduct.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, "err whlie validating")
	}

	if err := adminController.service.AddProduct(reqproduct); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, "Product created Successfully")

}

func (adminController *AdminController) DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	if id != "" {
		if err := adminController.service.DeleteProduct(id); err != nil {
			return c.JSON(http.StatusInternalServerError, "Could not delete the product")
		}
		return c.JSON(http.StatusNoContent, "Delete Successfull")
	}
	return c.JSON(http.StatusInternalServerError, "Delete Unsuccessfull")
}
