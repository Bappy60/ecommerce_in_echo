package connection

import (
	"fmt"

	"github.com/Bappy60/ecommerce_in_echo/pkg/config"
	"github.com/Bappy60/ecommerce_in_echo/pkg/models"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Connect() {
	config := config.LocalConfig
	connectionString := fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local",
	config.DBUser, config.DBPass, config.DBIP, config.DbName)
	d, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}
	DB = d
}
func GetDB() *gorm.DB {
	return DB
}

func Initialize() *gorm.DB {
	Connect()
	db := GetDB()
	db.AutoMigrate(
		&models.User{}, 
		&models.ProductCategory{},
		&models.Product{},
		&models.Address{},
		&models.SelectedProduct{},
		&models.Order{},
		&models.Payment{},
	)
	return db
}
