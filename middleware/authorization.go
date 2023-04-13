package middleware

import (
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func AdminAuthorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userData := c.Get("userData").(jwt.MapClaims)
		isAdmin := userData["is_admin"].(bool)

		if !isAdmin {
			err := errors.New("not admin user")

			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "not admin user",
				"error":   err.Error(),
			})
		}

		return next(c)
	}
}
