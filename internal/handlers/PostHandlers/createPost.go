package PostHandlers

import (
	"DevConnector/database"
	"DevConnector/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func CreatePost(c echo.Context) error {
	MyCookie, err := c.Cookie("Id")

	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "Unauthorized",
		})
	}

	myId, _ := strconv.Atoi(MyCookie.Value)
	var me models.Profile

	if err := database.DB.First(&me, myId).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "failed to find your profile by id",
		})
	}

	var body models.Post

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "failed to bind post data",
		})
	}

	post := models.Post{
		ProfileID: uint(myId),
		Creator:   body.Creator,
		Text:      body.Text,
	}

	if err := database.DB.Create(&post).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "failed to create post",
		})
	}

	me.Posts = append(me.Posts, post)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "post has been successfully created",
	})
}
