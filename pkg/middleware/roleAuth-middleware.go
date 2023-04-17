package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role := c.Get("role").(string)
		// Check if user has the required role
		if role != "admin" {
			return c.JSON(http.StatusForbidden, "You do not have permission to access this resource")
		}
		return next(c)
	}
}

