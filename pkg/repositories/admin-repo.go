package repositories

import (
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
			Message: "Product already exists",
			Err:     err,
		}
	}

	if err := repo.db.Create(&product).Error; err != nil {
		return &types.CustomError{
			Message: "Product not created",
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
			Message: "Invalid Id",
			Err:     err,
		}
	}
	if err := repo.db.Unscoped().Where("id =?", product.ID).Delete(product).Error; err != nil {
		return &types.CustomError{
			Message: "Could not delete the product",
			Err:     err,
		}
	}
	return nil
}
