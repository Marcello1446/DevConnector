package IDMatching

import (
	"DevConnector/database"
	"DevConnector/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CheckForID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		existedProfile := c.Get("existedProfile").(models.Profile)
		postID := c.Param("id")

		var profile models.Profile

		if err := database.DB.First(&profile, existedProfile.ID).Error; err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": "failed to find profile by id",
			})
		}

		var post models.Post

		if err := database.DB.First(&post, postID).Error; err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": "failed to find post by id",
			})
		}

		if profile.ID != post.ProfileID {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": "you cannot update this post, it doesnt belong to you",
			})
		}

		c.Set("post", post)
		return next(c)
	}
}
