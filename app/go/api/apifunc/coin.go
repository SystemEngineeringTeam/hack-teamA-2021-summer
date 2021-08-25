package apifunc

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CoinPost(c echo.Context) error {
	return c.String(http.StatusOK, "coin_post") // フロントに返す値
}
