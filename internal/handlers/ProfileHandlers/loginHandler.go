package ProfileHandlers

import (
	"DevConnector/database"
	"DevConnector/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func Login(c echo.Context) error {
	body := c.Get("body").(models.Login)

	var profile models.Profile

	if err := database.DB.Where("email = ?", body.Email).First(&profile).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "incorrect email",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(profile.Password), []byte(body.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "incorrect password",
		})
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": profile.ID,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	})

	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": "Failed to sign token",
		})
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": profile.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": "Failed to sign token",
		})
	}

	accessCookie := &http.Cookie{
		Name:     "Access",
		Value:    accessTokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	refreshCookie := &http.Cookie{
		Name:     "Refresh",
		Value:    refreshTokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	idCookie := &http.Cookie{
		Name:     "Id",
		Value:    strconv.Itoa(int(profile.ID)),
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	c.SetCookie(accessCookie)
	c.SetCookie(refreshCookie)
	c.SetCookie(idCookie)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "successful login!",
	})
}
