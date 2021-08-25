package apifunc

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginPost(c echo.Context) error {
	return c.String(http.StatusOK, "login_post") // フロントに返す値
}
