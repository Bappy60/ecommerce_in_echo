package container

import (
	"github.com/Bappy60/ecommerce_in_echo/pkg/config"
)

func Serve() {
	config.SetConfig()
	//var db = connection.Initialize()

	// bookRepo := repositories.BookDBInstance(db)
	// bookService := services.BookServiceInstance(bookRepo)
	// bookController := controllers.BookControllerInstance(bookService)

	// authorRepo := repositories.AuthorDBInstance(db)
	// authorService := services.AuthorServiceInstance(authorRepo)
	// authorController := controllers.AuthorControllerInstance(authorService)

	// log.Println("Database Connected...")
	// r := mux.NewRouter()
	// routes.AuthorRoutes(r, authorController)
	// routes.BookRoutes(r,bookController)
	// http.Handle("/", r)
	// log.Println("Server Started...")
	// log.Fatal(http.ListenAndServe("localhost:9011", r))
}
