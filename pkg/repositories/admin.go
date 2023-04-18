package repositories

import (
	"github.com/Bappy60/ecommerce_in_echo/pkg/consts"
	"github.com/Bappy60/ecommerce_in_echo/pkg/domain"
	"github.com/Bappy60/ecommerce_in_echo/pkg/models"
	"github.com/Bappy60/ecommerce_in_echo/pkg/types"
	"github.com/jinzhu/gorm"
)

type AdminRepo struct {
	db *gorm.DB
}

func AdminDBInstance(d *gorm.DB) domain.IAdminRepo {
	return &AdminRepo{
		db: d,
	}
}

// AddProduct implements domain.IAdminRepo
func (repo *AdminRepo) AddProduct(product models.Product) error {
	if err := repo.db.Where("name = ?", product.Name).First(&models.Product{}).Error; err == nil {
		return &types.CustomError{
			Message: consts.ProductExists,
			Err:     err,
		}
	}

	if err := repo.db.Create(&product).Error; err != nil {
		return &types.CustomError{
			Message: consts.ProductNotCreated,
			Err:     err,
		}
	}
	return nil
}

// DeleteProduct implements domain.IAdminRepo
func (repo *AdminRepo) DeleteProduct(id uint64) error {
	product := models.Product{}
	product.ID = id

	if err := repo.db.Where("id = ?", product.ID).First(&models.Product{}).Error; err != nil {
		return &types.CustomError{
			Message: consts.InvalidID,
			Err:     err,
		}
	}
	if err := repo.db.Unscoped().Where("id =?", product.ID).Delete(product).Error; err != nil {
		return &types.CustomError{
			Message: consts.DeleteUnsuccessful,
			Err:     err,
		}
	}
	return nil
}
