package ProfileHandlers

import (
	"DevConnector/database"
	"DevConnector/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func FollowProfile(c echo.Context) error {
	FollowingId := c.Param("id")
	var profile models.Profile

	if err := database.DB.First(&profile, FollowingId).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "failed to find profile by id",
		})
	}

	MyCookie, err := c.Cookie("Id")

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "can not take your id",
		})
	}

	myId, _ := strconv.Atoi(MyCookie.Value)
	var me models.Profile
	
	if err := database.DB.First(&me, myId).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "failed to find your profile by id",
		})
	}

	if !contains(profile.Followers, me.Username) {
		profile.Followers = append(profile.Followers, me.Username)
		profile.FollowersCount++
	}

	if !contains(me.Following, profile.Username) {
		me.Following = append(me.Following, profile.Username)
	}

	if err := database.DB.Save(&profile).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "failed to save updated profile in database",
		})
	}

	if err := database.DB.Save(&me).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "failed to save updated profile in database",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "you began to follow user!",
	})
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
