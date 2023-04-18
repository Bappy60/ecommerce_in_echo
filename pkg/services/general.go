package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Bappy60/ecommerce_in_echo/pkg/domain"
	"github.com/Bappy60/ecommerce_in_echo/pkg/models"
	"github.com/Bappy60/ecommerce_in_echo/pkg/types"
	"github.com/redis/go-redis/v9"
)

type GeneralService struct {
	repo        domain.IGeneralRepo
	redisClient *redis.Client
}

func GeneralServiceInstance(generalRepo domain.IGeneralRepo, redisClient *redis.Client) domain.IGeneralService {
	return &GeneralService{
		repo:        generalRepo,
		redisClient: redisClient,
	}
}

var ctx = context.Background()

// SearchProduct implements domain.IGeneralService
func (generalService *GeneralService) SearchProduct(searchReq *types.SearchRequest) ([]models.Product, error) {

	cacheKey := fmt.Sprintf("%d:%s:%f:%s", searchReq.Id, searchReq.Name, searchReq.Price, searchReq.Category)
	cachedData, err := generalService.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var products []models.Product
		err := json.Unmarshal([]byte(cachedData), &products)
		if err != nil {
			return nil, &types.CustomError{
				Message: err.Error(),
				Err:     err,
			}
		}
		return products, nil
	}

	products, err := generalService.repo.SearchProduct(searchReq)
	if err != nil {
		return nil, &types.CustomError{
			Message: err.Error(),
			Err:     err,
		}
	}
	jsonData, err := json.Marshal(products)
	if err != nil {
		return nil, &types.CustomError{
			Message: err.Error(),
			Err:     err,
		}
	}
	if _, err := generalService.redisClient.Set(ctx, cacheKey, jsonData, time.Hour).Result(); err != nil {
		return nil, &types.CustomError{
			Message: err.Error(),
			Err:     err,
		}
	}
	return products, nil

}
