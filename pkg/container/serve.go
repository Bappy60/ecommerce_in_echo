package container

import (
	"log"

	"github.com/Bappy60/ecommerce_in_echo/pkg/config"
	"github.com/Bappy60/ecommerce_in_echo/pkg/connection"
	"github.com/Bappy60/ecommerce_in_echo/pkg/controllers"
	"github.com/Bappy60/ecommerce_in_echo/pkg/routes"
	"github.com/labstack/echo/v4"
)

func Serve() {
	e := echo.New()
	config.SetConfig()
	var db = connection.Initialize()
	userController := controllers.UserControllerInstance(db)
	adminController := controllers.AdminControllerInstance(db)
	generalController := controllers.GeneralControllerInstance(db)
	log.Println("Database Connected...")
	routes.UserRoutes(e, userController)
	routes.AdminRoutes(e,adminController)
	routes.GeneralRoutes(e,generalController)
	e.Logger.Fatal(e.Start(":" + config.LocalConfig.Port))

}
// userRepo := repositories.UserDBInstance(db)
	// userService := services.UserServiceInstance(userRepo)
	// userController := controllers.UserControllerInstance(userService)

	// bookRepo := repositories.BookDBInstance(db)
	// bookService := services.BookServiceInstance(bookRepo)
	// bookController := controllers.BookControllerInstance(bookService)

	// authorRepo := repositories.AuthorDBInstance(db)
	// authorService := services.AuthorServiceInstance(authorRepo)
	// authorController := controllers.AuthorControllerInstance(authorService)

	// r := mux.NewRouter()
	// routes.AuthorRoutes(r, authorController)
	// routes.BookRoutes(r,bookController)
	// http.Handle("/", r)
	// log.Println("Server Started...")
	// log.Fatal(http.ListenAndServe("localhost:9011", r))

