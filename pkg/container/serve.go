package container

import (
	"log"

	"github.com/Bappy60/ecommerce_in_echo/pkg/config"
	"github.com/Bappy60/ecommerce_in_echo/pkg/connection"
	"github.com/Bappy60/ecommerce_in_echo/pkg/controllers"
	"github.com/Bappy60/ecommerce_in_echo/pkg/repositories"
	"github.com/Bappy60/ecommerce_in_echo/pkg/routes"
	"github.com/Bappy60/ecommerce_in_echo/pkg/services"
	"github.com/labstack/echo/v4"
)

func Serve() {
	e := echo.New()
	config.SetConfig()
	var db = connection.Initialize()
	
	userRepo := repositories.UserDBInstance(db)
	userService := services.UserServiceInstance(userRepo)
	userController := controllers.UserControllerInstance(userService)

	cartRepo := repositories.CartDBInstance(db)
	cartService := services.CartServiceInstance(cartRepo)
	cartController := controllers.CartControllerInstance(cartService)

	generalRepo := repositories.GeneralDBInstance(db)
	generalService := services.GeneralServiceInstance(generalRepo)
	generalController := controllers.GeneralControllerInstance(generalService)

	adminRepo := repositories.AdminDBInstance(db)
	adminService := services.AdminServiceInstance(adminRepo)
	adminController := controllers.AdminControllerInstance(adminService)


	 log.Println("Database Connected...")
	 routes.UserRoutes(e, userController)
	 routes.CartRoutes(e,cartController)
	 routes.GeneralRoutes(e,generalController)
	 routes.AdminRoutes(e,adminController)
	 
	e.Logger.Fatal(e.Start(":" + config.LocalConfig.Port))

}
