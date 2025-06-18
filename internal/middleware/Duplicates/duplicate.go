package Duplicates

import (
	"DevConnector/database"
	"DevConnector/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CheckForDuplicates(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		profile := c.Get("body").(models.Profile)

		var profiles []models.Profile
		if err := database.DB.Find(&profiles).Error; err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": "failed to parse all db profiles",
			})
		}

		for _, item := range profiles {
			if item.Email == profile.Email {
				return c.JSON(http.StatusBadRequest, echo.Map{
					"error": "profile with same email already exists",
				})
			} else if item.Username == profile.Username {
				return c.JSON(http.StatusBadRequest, echo.Map{
					"error": "profile with same username already exists",
				})
			}

		}

		c.Set("body", profile)

		return next(c)
	}
}
