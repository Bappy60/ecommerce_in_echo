package middleware

import (
	"net/http"

	"github.com/Bappy60/ecommerce_in_echo/pkg/tokens"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		authHeadertoken := c.Request().Header.Get("token")

		if authHeadertoken == "" {
			return c.JSON(http.StatusUnauthorized, "token is needed Unauthorized:(")
		}
		claims, err := tokens.ValidateToken(authHeadertoken)
		if err != "" {
			return c.JSON(http.StatusUnauthorized, err)
		}
		c.Set("email",claims.Email)
		c.Set("userId",claims.UserId)
		return next(c)
	}
}
