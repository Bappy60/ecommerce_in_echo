package controllers

import (
	"log"
	"net/http"

	"github.com/Bappy60/ecommerce_in_echo/pkg/models"
	"github.com/Bappy60/ecommerce_in_echo/pkg/tokens"
	"github.com/Bappy60/ecommerce_in_echo/pkg/types"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// type UserController struct {
// 	service domain.IUserService
// }

// func UserControllerInstance(userService domain.IUserService) domain.IUserController {
// 	return &UserController{
// 		service: userService,
// 	}
// }

type UserController struct {
	db *gorm.DB
}

func UserControllerInstance(db *gorm.DB) *UserController {
	return &UserController{
		db: db,
	}
}

// SignUp implements domain.IUserController
func (userController *UserController) SignUp(c echo.Context) error {
	user := types.SignReqStruct{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err := user.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := userController.db.Where("name = ? AND email = ?", user.Name, user.Email).First(&models.User{}).Error; err == nil {
		return c.JSON(http.StatusBadRequest, "User already exists")
	}
	password := HashPassword(user.Password)
	newUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: password,
		Address:  user.Address,
	}
	if err := userController.db.Create(&newUser).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "SignUp failed")
	}
	return c.JSON(http.StatusCreated, "Successfully Signed Up!!")
}

// Login implements domain.IUserController
func (userController *UserController) Login(c echo.Context) error {
	user := types.LoginReqStruct{}
	foundUser := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err := user.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := userController.db.Where("email = ?", user.Email).First(&foundUser).Error; err != nil {
		return c.JSON(http.StatusBadRequest, "email or password incorrect")
	}

	IsValidPassword, msg := VerifyPassword(user.Password, foundUser.Password)
	if !IsValidPassword {
		return c.JSON(http.StatusInternalServerError, msg)
	}

	token, err := tokens.TokenGenerator(foundUser.Email)
	if err != nil {
		log.Println("Error Creating JWT token", err)
		return c.JSON(http.StatusInternalServerError, "something went wrong")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "You were logged in!",
		"token":   token,
	})

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
		msg = "email Or Passowrd is Incorerct"
		valid = false
	}
	return valid, msg
}
