package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Bappy60/ecommerce_in_echo/pkg/domain"
	"github.com/Bappy60/ecommerce_in_echo/pkg/models"
	"github.com/Bappy60/ecommerce_in_echo/pkg/types"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

type GeneralController struct {
	service     domain.IGeneralService
	redisClient *redis.Client
}

func GeneralControllerInstance(generalService domain.IGeneralService, redisClient *redis.Client) domain.IGeneralController {
	return &GeneralController{
		service:     generalService,
		redisClient: redisClient,
	}
}

var ctx = context.Background()

func (generalController *GeneralController) SearchProduct(c echo.Context) error {
	id := c.QueryParam("id")
	name := c.QueryParam("name")
	price := c.QueryParam("price")
	categoryName := c.QueryParam("category_name")

	cacheKey := fmt.Sprintf("%s:%s:%s:%s", id, name, price, categoryName)

	cachedData, err := generalController.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var products []models.Product
		err := json.Unmarshal([]byte(cachedData), &products)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &types.CustomError{
				Message: err.Error(),
				Err:     err,
			})
		}
		return c.JSON(http.StatusOK, products)
	}

	searchReq := &types.SearchRequest{
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
	jsonData, err := json.Marshal(products)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &types.CustomError{
			Message: err.Error(),
			Err:     err,
		})
	}
	_, err2 := generalController.redisClient.Set(ctx, cacheKey, jsonData,time.Hour).Result()
	if err2 != nil {
		return c.JSON(http.StatusInternalServerError, &types.CustomError{
			Message: err.Error(),
			Err:     err,
		})
	}
	return c.JSON(http.StatusOK, products)
}
