package CommentHandlers

import (
	"DevConnector/database"
	"DevConnector/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func AddComment(c echo.Context) error {
	myCookie, err := c.Cookie("Id")

	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "unauthorized",
		})
	}

	myStringId := myCookie.Value
	myId, _ := strconv.Atoi(myStringId)

	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "invalid post ID",
		})
	}

	var post models.Post
	if err := database.DB.First(&post, intId).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "failed to find post by its id",
		})
	}

	var body models.Comment
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "failed to bind comment",
		})
	}

	comment := models.Comment{
		ProfileID: uint(myId),
		PostID:    post.ID,
		Creator:   body.Creator,
		Text:      body.Text,
	}

	if err := database.DB.Create(&comment).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "failed to create a comment",
		})
	}

	post.Comments = append(post.Comments, comment)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "comment has been successfully added to a post",
	})
}
