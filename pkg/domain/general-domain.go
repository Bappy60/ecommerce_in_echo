package domain

import (
	"github.com/Bappy60/ecommerce_in_echo/pkg/models"
	"github.com/Bappy60/ecommerce_in_echo/pkg/types"
	"github.com/labstack/echo/v4"
)

type IGeneralController interface {
	SearchProduct(c echo.Context) error
}

type IGeneralService interface {
	SearchProduct(searchReq *types.SearchReqStruct) ([]models.Product,error)
}
type IGeneralRepo interface {
	SearchProduct(searchRepo *types.SearchRepoStruct) ([]models.Product,error)
}
