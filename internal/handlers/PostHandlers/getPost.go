package PostHandlers

import (
	"DevConnector/database"
	"DevConnector/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetPost(c echo.Context) error {
	id := c.Param("id")

	var post models.Post

	if err := database.DB.Find(&post, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "failed to find post by id",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"post": post,
	})
}
