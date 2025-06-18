package ProfileHandlers

import (
	"DevConnector/database"
	"DevConnector/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetFollowers(c echo.Context) error {
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

	return c.JSON(http.StatusOK, echo.Map{
		"followers": me.Followers,
	})
}
