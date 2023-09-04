package handlers

import (
	"login-pro/models"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func LogIn(c echo.Context) (err error) {
	var user models.User
	if err = c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// TODO) 검증 로직 추가해야함.
	ts, err := CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	saveErr := CreateAuth(user.ID, ts)
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	}
	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}

	accessExpiry, _ := strconv.Atoi(os.Getenv("HTTP_COOKIE_ACCESS_EXPIRY"))
	refreshExpiry, _ := strconv.Atoi(os.Getenv("HTTP_COOKIE_REFRESH_EXPIRY"))

	c.SetCookie(&http.Cookie{
		Name:    "access",
		Value:   tokens["access_token"],
		Path:    "/",
		Expires: time.Now().Add(time.Minute * time.Duration(accessExpiry)),
		Domain:  "127.0.0.1",
		// Secure:  true, // https 일 경우에만 true
		// 현재 chrome에서는 https에서만 쿠키 저장이 가능하기 때문에 firefox에서 진행
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})

	c.SetCookie(&http.Cookie{
		Name:    "refresh",
		Value:   tokens["refresh_token"],
		Path:    "/",
		Expires: time.Now().Add(time.Minute * time.Duration(refreshExpiry)),
		Domain:  "127.0.0.1",
		// Secure: true, // https 일 경우에만 true
		// 현재 chrome에서는 https에서만 쿠키 저장이 가능하기 때문에 firefox에서 진행
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})

	return c.JSON(http.StatusOK, "success login")
}

func LogOut(c echo.Context) (err error) {
	au, err := ExtractTokenMetadata(c.Request())
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	deleted, delErr := DeleteAuth(au.AccessUuid)
	if delErr != nil || deleted == 0 {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	c.JSON(http.StatusOK, "Successfully logged out")
	return
}
