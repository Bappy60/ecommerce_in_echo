package controllers

import (
	"net/http"
	"strconv"

	"github.com/Bappy60/ecommerce_in_echo/pkg/consts"
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
	requestProduct := types.CreateProduct{}
	if err := c.Bind(&requestProduct); err != nil {
		return c.JSON(http.StatusBadRequest, consts.BadRequest)
	}
	err := requestProduct.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, consts.ValidationErr)
	}

	if err := adminController.service.AddProduct(requestProduct); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, consts.ProductCreated)

}

func (adminController *AdminController) DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	ID, err := strconv.ParseUint(id, 10, 64)
	if err != nil || id == "" {
		return &types.CustomError{
			Message: consts.ParseErr,
			Err:     err,
		}
	}
	if err := adminController.service.DeleteProduct(ID); err != nil {
		return c.JSON(http.StatusInternalServerError, consts.DeleteUnsuccessful)
	}
	return c.JSON(http.StatusNoContent, consts.DeleteSuccessful)
}
