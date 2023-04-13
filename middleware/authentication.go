package middleware

import (
	"movie-app/helpers"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		verifyToken, err := helpers.VerifyToken(c)
		id := verifyToken.(jwt.MapClaims)["id"]

		if err != nil || id == nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error":   "Unauthenticated",
				"message": err.Error(),
			})
		}

		c.Set("userData", verifyToken)
		return next(c)
	}
}
