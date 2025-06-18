package validators

import (
	"DevConnector/internal/models"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

var validate *validator.Validate

func ValidateProfileCreating(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var profile models.Profile

		if err := c.Bind(&profile); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": "Failed to bind data(validating)",
			})
		}

		validate = validator.New()

		if err := validate.Struct(&profile); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": "Failed to validate data",
			})
		}

		c.Set("body", profile)

		return next(c)
	}
}

func ValidateProfileLoggining(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var login models.Login

		if err := c.Bind(&login); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": "Failed to bind data",
			})
		}

		validate = validator.New()

		if err := validate.Struct(&login); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": "Failed to validate data",
			})
		}

		c.Set("body", login)

		return next(c)
	}
}
