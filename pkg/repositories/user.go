package repositories

import (
	"log"
	"time"

	"github.com/Bappy60/ecommerce_in_echo/pkg/consts"
	"github.com/Bappy60/ecommerce_in_echo/pkg/domain"
	"github.com/Bappy60/ecommerce_in_echo/pkg/models"
	"github.com/Bappy60/ecommerce_in_echo/pkg/tokens"
	"github.com/Bappy60/ecommerce_in_echo/pkg/types"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type userRepo struct {
	db *gorm.DB
}

func UserDBInstance(d *gorm.DB) domain.IUserRepo {
	return &userRepo{
		db: d,
	}
}

func (userRepo *userRepo) SignUp(user *types.SignInRequest) error {

	err := userRepo.db.Where("name = ? AND email = ?", user.Name, user.Email).First(&models.User{}).Error
	if err == nil {
		return &types.CustomError{
			Message: consts.DuplicateUser,
		}
	}
	password := HashPassword(user.Password)
	newUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: password,
		Address:  user.Address,
	}
	if err := userRepo.db.Create(&newUser).Error; err != nil {
		return &types.CustomError{
			Message: consts.SignUpUnSuccessful,
		}
	}
	return nil
}

func (repo *userRepo) Login(user *types.LoginRequset) (string, error) {
	foundUser := models.User{}
	if err := repo.db.Where("email = ?", user.Email).First(&foundUser).Error; err != nil {
		return "", &types.CustomError{
			Message: consts.IncorerctDetails,
		}
	}

	IsValidPassword, msg := VerifyPassword(user.Password, foundUser.Password)
	if !IsValidPassword {
		return "", &types.CustomError{
			Message: msg,
		}
	}
	token, err := tokens.TokenGenerator(foundUser.Email, foundUser.ID, foundUser.HasRole)
	if err != nil {
		return "", &types.CustomError{
			Message: consts.TokenGenErr,
		}
	}
	if err := repo.CreateCart(foundUser.ID); err != nil {
		return "", &types.CustomError{
			Message: consts.CartCreationErr,
		}
	}
	return token, nil
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}
func VerifyPassword(userpassword string, givenpassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(givenpassword), []byte(userpassword))
	valid := true
	msg := ""
	if err != nil {
		msg = consts.IncorerctDetails
		valid = false
	}
	return valid, msg
}
func (userRepo *userRepo) CreateCart(userId uint64) error {

	cart := models.Cart{}
	if err := userRepo.db.Where("user_id = ?", userId).First(&cart).Error; err == nil {
		return nil
	} else {
		newCart := models.Cart{
			UserID:    userId,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := userRepo.db.Create(&newCart).Error; err != nil {
			return err
		}
		return nil
	}
}
