package ProfileHandlers

import (
	"DevConnector/database"
	"DevConnector/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func UnfollowProfile(c echo.Context) error {
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

	if contains(profile.Followers, me.Username) {
		for i, v := range profile.Followers {
			if v == me.Username {
				profile.Followers = append(profile.Followers[:i], profile.Followers[i+1])
			}
		}
		for index, value := range me.Following {
			if value == profile.Username {
				me.Following = append(me.Following[:index], me.Following[index+1])
			}
		}
		profile.FollowersCount--
	} else {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "failed to unfollow user",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "you successfully unfollowed profile",
	})
}

