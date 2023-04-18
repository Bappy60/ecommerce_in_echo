package repositories

import (
	"github.com/Bappy60/ecommerce_in_echo/pkg/domain"
	"github.com/Bappy60/ecommerce_in_echo/pkg/models"
	"github.com/Bappy60/ecommerce_in_echo/pkg/types"
	"github.com/jinzhu/gorm"
)

type GeneralRepo struct {
	db *gorm.DB
}

func GeneralDBInstance(d *gorm.DB) domain.IGeneralRepo {
	return &GeneralRepo{
		db: d,
	}
}

// SearchProduct implements domain.IGeneralRepo
func (repo *GeneralRepo) SearchProduct(searchReq *types.SearchRequest) ([]models.Product, error) {
	products := []models.Product{}
	query := repo.db.Model(&models.Product{}).Preload("Category")

	if searchReq.Id == 0 && searchReq.Name == "" && searchReq.Category == "" && searchReq.Price == 0 {
		query.Find(&products)
		return products, nil
	}

	if searchReq.Id != 0 {
		query.Where("id =?", searchReq.Id).Find(&products)
		return products, nil
	} else {
		if searchReq.Name != "" {
			query = query.Where("name LIKE ?", "%"+searchReq.Name+"%")
		}
		if searchReq.Category != "" {
			query = query.Joins("JOIN product_categories ON products.category_id = product_categories.id")
			query = query.Where("product_categories.category_name LIKE ?", "%"+searchReq.Category+"%")
		}
		if searchReq.Price != 0 {
			query = query.Where("price =?", searchReq.Price)
		}
		if err := query.Find(&products).Error; err != nil {
			return products, nil
		}
	}
	return products, nil
}
