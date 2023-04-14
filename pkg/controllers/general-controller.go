package controllers

import (
	"net/http"

	"github.com/Bappy60/ecommerce_in_echo/pkg/domain"
	"github.com/Bappy60/ecommerce_in_echo/pkg/types"
	"github.com/labstack/echo/v4"
)

// type GeneralController struct {
// 	db *gorm.DB
// }

//	func GeneralControllerInstance(db *gorm.DB) *GeneralController {
//		return &GeneralController{
//			db: db,
//		}
//	}
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

	searchReq := &types.SearchReqStruct{
		Id:       id,
		Name:     name,
		Price:    price,
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
