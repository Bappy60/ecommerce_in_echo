package controllers

import (
	"net/http"
	"strconv"

	"github.com/Bappy60/ecommerce_in_echo/pkg/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type GeneralController struct {
	db *gorm.DB
}

func GeneralControllerInstance(db *gorm.DB) *GeneralController {
	return &GeneralController{
		db: db,
	}
}

func (generalController *GeneralController) SearchProduct(c echo.Context) error {
	id := c.QueryParam("id")
	Name := c.QueryParam("name")
	price := c.QueryParam("price")
	categoryName := c.QueryParam("category_name")

	parsedId, err := strconv.ParseUint(id, 0, 0)
	if err != nil && id != "" {
		return c.JSON(http.StatusBadRequest, err)
	}
	parsedPrice, err := strconv.ParseUint(price, 0, 0)
	if err != nil && price != "" {
		return c.JSON(http.StatusBadRequest, err)
	}

	products := []models.Product{}
	query := generalController.db.Model(&models.Product{}).Preload("Category")

	if parsedId == 0 && Name == "" && categoryName == "" && parsedPrice == 0 {
		query.Find(&products)
		return c.JSON(http.StatusOK, products)
	}

	if parsedId != 0 {
		query.Where("id =?", parsedId).Find(&products)
		return c.JSON(http.StatusOK, products)
	} else {
		if Name != "" {
			query = query.Where("name LIKE ?", "%"+Name+"%")
		}
		if categoryName != "" {
			query = query.Joins("JOIN product_categories ON products.category_id = product_categories.id")
			query = query.Where("product_categories.category_name LIKE ?", "%"+categoryName+"%")
		}
		if parsedPrice != 0 {
			query = query.Where("price =?", parsedPrice)
		}
		if err := query.Find(&products).Error; err != nil {
			return c.JSON(http.StatusOK, products)
		}
	}
	return c.JSON(http.StatusOK, products)
}
