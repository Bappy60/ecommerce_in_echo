package services

import (
	"strconv"

	"github.com/Bappy60/ecommerce_in_echo/pkg/domain"
	"github.com/Bappy60/ecommerce_in_echo/pkg/models"
	"github.com/Bappy60/ecommerce_in_echo/pkg/types"
)

type AdminService struct {
	repo domain.IAdminRepo
}
func AdminServiceInstance(adminRepo domain.IAdminRepo) domain.IAdminService {
	return &AdminService{
		repo: adminRepo,
	}
}

// AddProduct implements domain.IAdminService
func (service *AdminService) AddProduct(reqproduct types.CreateProductStruct) error {

		product := models.Product{
		Name:     reqproduct.Name,
		Price:    reqproduct.Price,
		Category: reqproduct.Category,
	}
	if err:= service.repo.AddProduct(product); err!=nil {
		return err
	}
	return nil
}

// DeleteProduct implements domain.IAdminService
func (service *AdminService) DeleteProduct(id string) error {
	Id, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return &types.CustomError{
			Message: "Invalid Id",
			Err: err,
		}
	}

	if err:= service.repo.DeleteProduct(Id); err!= nil{
		return err
	}
	return nil
}

