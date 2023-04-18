package services

import (
	"strconv"

	"github.com/Bappy60/ecommerce_in_echo/pkg/domain"
	"github.com/Bappy60/ecommerce_in_echo/pkg/models"
	"github.com/Bappy60/ecommerce_in_echo/pkg/types"
)

type GeneralService struct {
	repo domain.IGeneralRepo
}

func GeneralServiceInstance(generalRepo domain.IGeneralRepo) domain.IGeneralService {
	return &GeneralService{
		repo: generalRepo,
	}
}

// SearchProduct implements domain.IGeneralService
func (generalService *GeneralService) SearchProduct(searchReq *types.SearchReqStruct) ([]models.Product, error) {
	parsedId, err := strconv.ParseUint(searchReq.Id, 10, 64)
	if err != nil && searchReq.Id != "" {
		return nil, &types.CustomError{
			Message: "Invalid format of id",
			Err:     err,
		}
	}


	parsedPrice, err := strconv.ParseFloat(searchReq.Price, 64)
	if err != nil && searchReq.Price != "" {
		return nil, &types.CustomError{
			Message: "Invalid format of price",
			Err:     err,
		}
	}

	searchRepoStruct := &types.SearchRepoStruct{
		Id:       parsedId,
		Name:     searchReq.Name,
		Price:    parsedPrice,
		Category: searchReq.Category,
	}
	Products, err2 := generalService.repo.SearchProduct(searchRepoStruct)
	if err2 != nil {
		return nil, &types.CustomError{
			Message: "",
			Err:     err2,
		}
	}
	return Products, nil

}