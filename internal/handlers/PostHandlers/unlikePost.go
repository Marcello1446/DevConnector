package PostHandlers

import (
	"DevConnector/database"
	"DevConnector/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func UnlikePost(c echo.Context) error {
	id := c.Param("id")

	var post models.Post

	if err := database.DB.Find(&post, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "failed to find post by id",
		})
	}

	post.Dislikes = post.Dislikes + 1

	if err := database.DB.Save(&post).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "failed to save record in database",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "post has been successfully disliked",
	})
}
