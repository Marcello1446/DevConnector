package ProfileHandlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func Logout(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:     "Access",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	c.SetCookie(&http.Cookie{
		Name:     "Refresh",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	c.SetCookie(&http.Cookie{
		Name:     "Id",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	return c.JSON(http.StatusOK, echo.Map{
		"message": "successful logout",
	})
}
