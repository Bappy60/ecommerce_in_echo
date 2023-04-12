package domain

import (
	"github.com/labstack/echo/v4"
)

type IGeneralController interface {
	SearchProduct(c echo.Context) error
}

// type IGeneralService interface {
// 	SearchProduct(c echo.Context) error
// }
// type IGeneralRepo interface {
// 	SearchProduct(c echo.Context) error
// }
