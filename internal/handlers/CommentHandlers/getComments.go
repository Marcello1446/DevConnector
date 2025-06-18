package CommentHandlers

import (
	"DevConnector/database"
	"DevConnector/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetComments(c echo.Context) error {
	id := c.Param("id")
	postId, _ := strconv.Atoi(id)

	var post models.Post

	if err := database.DB.Preload("Comments").First(&post, postId).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "failed to find post by id",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"comments": post.Comments,
	})
}
