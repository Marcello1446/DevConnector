package ProfileHandlers

import (
	"DevConnector/database"
	"DevConnector/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetProfileById(c echo.Context) error {
	id := c.Param("id")
	var profile models.Profile

	if err := database.DB.Preload("Posts").First(&profile, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "failed to find profile in database",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"username":       profile.Username,
		"bio":            profile.Bio,
		"github":         profile.Github,
		"followersCount": profile.FollowersCount,
		"posts":          profile.Posts,
		"followers":      profile.Followers,
		"following":      profile.Following,
	})
}
