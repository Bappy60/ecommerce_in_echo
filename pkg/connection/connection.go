package connection

import (
	"fmt"

	"github.com/Bappy60/ecommerce_in_echo/pkg/config"
	"github.com/Bappy60/ecommerce_in_echo/pkg/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Connect() {
	config := config.LocalConfig
	var connectionString string
	if config.APP_MODE == "prod" {
		connectionString = fmt.Sprintf("%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.DBURL)
	} else {
		connectionString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName)
	}
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
	// db.DropTable(&models.CartItem{})
	db.AutoMigrate(
		&models.User{},
		&models.ProductCategory{},
		&models.Product{},
		&models.Address{},
		&models.Cart{},
		// &models.Order{},
		// &models.OrderItem{},
		&models.CartItem{},
	)
	return db
}
