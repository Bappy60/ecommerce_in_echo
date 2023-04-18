package controllers

import (
	"net/http"
	"strconv"

	"github.com/Bappy60/ecommerce_in_echo/pkg/consts"
	"github.com/Bappy60/ecommerce_in_echo/pkg/domain"
	"github.com/Bappy60/ecommerce_in_echo/pkg/types"
	"github.com/labstack/echo/v4"
)

type GeneralController struct {
	service domain.IGeneralService
}

func GeneralControllerInstance(generalService domain.IGeneralService) domain.IGeneralController {
	return &GeneralController{
		service: generalService,
	}
}

func (generalController *GeneralController) SearchProduct(c echo.Context) error {
	id := c.QueryParam("id")
	name := c.QueryParam("name")
	price := c.QueryParam("price")
	categoryName := c.QueryParam("category_name")

	parsedId, err := strconv.ParseUint(id, 10, 64)
	if err != nil && id != "" {
		return c.JSON(http.StatusBadRequest, consts.ParseErr)
	}
	parsedPrice, err := strconv.ParseFloat(price, 64)
	if err != nil && price != "" {
		return c.JSON(http.StatusBadRequest, consts.ParseErr)
	}
	searchReq := &types.SearchRequest{
		Id:       parsedId,
		Name:     name,
		Price:    parsedPrice,
		Category: categoryName,
	}
	products, err := generalController.service.SearchProduct(searchReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &types.CustomError{
			Message: err.Error(),
			Err:     err,
		})
	}
	return c.JSON(http.StatusOK, products)
}
