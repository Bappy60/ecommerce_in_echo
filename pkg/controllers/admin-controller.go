package controllers

import (
	"net/http"

	"github.com/Bappy60/ecommerce_in_echo/pkg/models"
	"github.com/Bappy60/ecommerce_in_echo/pkg/types"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type AdminController struct {
	db *gorm.DB
}

func AdminControllerInstance(db *gorm.DB) *AdminController {
	return &AdminController{
		db: db,
	}
}
func (adminController *AdminController) AddProduct(c echo.Context) error {
	reqproduct := types.CreateProductStruct{}
	if err := c.Bind(&reqproduct); err != nil {
		return c.JSON(http.StatusBadRequest, "err while binding")
	}
	err := reqproduct.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, "err whlie validating")
	}
	product := models.Product{
		Name:     reqproduct.Name,
		Price:    reqproduct.Price,
		Category: reqproduct.Category,
	}
	if err := adminController.db.Where("name = ?", product.Name).First(&models.Product{}).Error; err == nil {
		return c.JSON(http.StatusBadRequest, "product already exists")
	}

	if err := adminController.db.Create(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Not Created")
	}
	return c.JSON(http.StatusCreated, "Product Created Successfully!!")

}

func (adminController *AdminController) DeleteProduct(c echo.Context) error {
	panic("unimplemented")
}
