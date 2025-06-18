package ProfileHandlers

import (
	"DevConnector/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ValidateExistance(c echo.Context) error {
	profile := c.Get("existedProfile").(models.Profile)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "logged in!",
		"user id": profile.ID,
	})
}
