package apifunc

import (
	"SystemEngineeringTeam/hack-teamA-2021-summer/dbfunc"
	"log"

	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type CoinPostParams struct { //受け取るデータの定義
	Email string `json:"email"`
	Total int    `json:"total"`
}

// requiredAuth
func CoinPost(c echo.Context) error {
	var params CoinPostParams
	if err := c.Bind(&params); err != nil { //送られてきたデータからエラーが出た時の処理
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "パラメータが正しくありません: " + err.Error()})
	}

	if err := c.Validate(&params); err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{"message": "パラメータが不足しています: " + err.Error()})
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	uid := claims["uid"]
	log.Println(uid)

	err := dbfunc.PostCoinInfo(params.Email, params.Total)
	if err != nil { //データベースでエラーが出た時の処理
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "データベースの更新に失敗しました: " + err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "成功しました"}) // フロントに返す値
}
