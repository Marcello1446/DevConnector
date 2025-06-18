package ProfileHandlers

import (
	"DevConnector/database"
	"DevConnector/internal/models"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func CreateProfile(c echo.Context) error {
	body := c.Get("body").(models.Profile)

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "failed to hash password",
		})
	}

	profile := models.Profile{Email: body.Email, Password: string(hash), Username: body.Username, Bio: body.Bio, Github: body.Github}

	if err := database.DB.Create(&profile).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "failed to create an user",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Profile has been successfully added!",
	})
}
