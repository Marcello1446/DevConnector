package ProfileHandlers

import (
	"DevConnector/database"
	"DevConnector/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func UpdateProfile(c echo.Context) error {
	id := c.Param("id")
	var profile models.Profile

	if err := database.DB.First(&profile, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Failed to find user with given id",
		})
	}

	var body models.Profile

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Failed to bind data",
		})
	}

	if body.Username != "" {
		profile.Username = body.Username
	}
	if body.Bio != "" {
		profile.Bio = body.Bio
	}
	if body.Github != "" {
		profile.Github = body.Github
	}

	if err := database.DB.Save(&profile).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "failed to save updated profile in database",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "profile was successfully updated",
	})
}
