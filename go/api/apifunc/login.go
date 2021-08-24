package apifunc

import (
	"net/http"

	"github.com/labstack/echo"
)

func Login_post(c echo.Context) error {
	return c.String(http.StatusOK, "login_post") // フロントに返す値
}
