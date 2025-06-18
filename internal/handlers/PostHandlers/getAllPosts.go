package PostHandlers

import (
	"DevConnector/database"
	"DevConnector/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAllPosts(c echo.Context) error {
	var posts []models.Post

	if err := database.DB.Model(&models.Post{}).Find(&posts).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "failed to load all posts",
		})
	}

	type PostResponse struct {
		Creator string `json:"creator"`
		Text    string `json:"text"`
	}

	var response []PostResponse
	for _, post := range posts {
		response = append(response, PostResponse{
			Creator: post.Creator,
			Text:    post.Text,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"posts": response,
	})

}
