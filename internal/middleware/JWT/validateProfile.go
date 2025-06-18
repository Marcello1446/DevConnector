package JWT

import (
	"DevConnector/database"
	"DevConnector/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"time"
)

func RequireProfileAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		accessCookie, err := c.Cookie("Access")

		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"error": "Unauthorized",
			})
		}

		tokenString := accessCookie.Value

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET")), nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": "Failed to parse token",
			})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			var profile models.Profile

			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				return c.JSON(http.StatusUnauthorized, err)
			}

			if err := database.DB.First(&profile, claims["sub"]).Error; err != nil {
				return c.JSON(http.StatusBadRequest, echo.Map{
					"error": "Failed to find user in a database",
				})
			}

			c.Set("existedProfile", profile)
		}

		return next(c)
	}

}
