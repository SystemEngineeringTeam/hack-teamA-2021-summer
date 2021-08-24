package apifunc

import (
	"net/http"

	"github.com/labstack/echo"
)

func Coin_post(c echo.Context) error {
	return c.String(http.StatusOK, "coin_post") // フロントに返す値
}
