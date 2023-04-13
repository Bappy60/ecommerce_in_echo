package services

import (
	"github.com/Bappy60/ecommerce_in_echo/pkg/domain"
	"github.com/Bappy60/ecommerce_in_echo/pkg/types"
)

type userService struct {
	repo domain.IUserRepo
}

func UserServiceInstance(userRepo domain.IUserRepo) domain.IUserService {
	return &userService{
		repo: userRepo,
	}
}

// SignUp implements domain.IUserService
func (userService *userService) SignUp(reqUser *types.SignReqStruct) error {

	err := userService.repo.SignUp(reqUser)
	if  err != nil {
		return err
	}
	return nil
}

// Login implements domain.IUserService
func (userService *userService) Login(reqUser *types.LoginReqStruct) (string, error) {

	token, err := userService.repo.Login(reqUser)
	if err != nil {
		return "", err
	}
	return token, nil
}
