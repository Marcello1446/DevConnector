package CommentHandlers

import (
	"DevConnector/database"
	"DevConnector/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func DeleteComment(c echo.Context) error {
	myCookie, err := c.Cookie("Id")

	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "Unauthorized",
		})
	}

	myID, _ := strconv.Atoi(myCookie.Value)
	postID, _ := strconv.Atoi(c.Param("postID"))
	commentID, _ := strconv.Atoi(c.Param("commentID"))

	var post models.Post

	if err := database.DB.First(&post, postID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "failed to find profile by this id",
		})
	}

	var comment models.Comment

	if err := database.DB.First(&comment, commentID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "failed to find profile by this id",
		})
	}

	if uint(myID) != comment.ProfileID {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "this comment doesnt belong to you",
		})
	}

	if err := database.DB.Unscoped().Delete(&comment).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "failed to delete comment",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "comment has been successfully deleted",
	})
}
