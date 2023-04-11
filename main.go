package main

import (
	"net/http"

	"github.com/Bappy60/ecommerce_in_echo/pkg/routes"
	"github.com/labstack/echo/v4"
)

func main(){
	e := echo.New()
	routes.UserRoutes(e)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}