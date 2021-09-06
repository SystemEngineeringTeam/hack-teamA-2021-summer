package apifunc

import (
	"SystemEngineeringTeam/hack-teamA-2021-summer/dbfunc"

	"net/http"

	"github.com/labstack/echo/v4"
)

type CoinPostParams struct { //
	Email string `json:"email" validate:"required"`
	Total int    `json:"total" validate:"required"`
}

func CoinPost(c echo.Context) error {
	var params CoinPostParams
	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "パラメータが正しくありません: " + err.Error()})
	}

	if err := c.Validate(&params); err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{"message": "パラメータが不足しています: " + err.Error()})
	}

	err := dbfunc.PostCoinInfo(params.Email, params.Total)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "データベースの更新に失敗しました: " + err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "成功しました"}) // フロントに返す値
}
