package PostHandlers

import (
	"DevConnector/database"
	"DevConnector/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func UpdatePost(c echo.Context) error {
	oldPost := c.Get("post").(models.Post)
	var newPost models.Post

	if err := c.Bind(&newPost); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "failed to bind data to new post",
		})
	}

	if newPost.Creator != "" {
		oldPost.Creator = newPost.Creator
	}
	if newPost.Text != "" {
		oldPost.Text = newPost.Text
	}

	if err := database.DB.Save(&oldPost).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "failed to save record in database",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "post has been successfully updated",
	})
}
