package apifunc

import (
	"net/http"

	"github.com/labstack/echo"
)

func User_get(c echo.Context) error {
	return c.String(http.StatusOK, "user_get") // フロントに返す値
}

func User_post(c echo.Context) error {
	return c.String(http.StatusOK, "user_post") // フロントに返す値
}

func User_put(c echo.Context) error {
	return c.String(http.StatusOK, "user_put") // フロントに返す値
}
