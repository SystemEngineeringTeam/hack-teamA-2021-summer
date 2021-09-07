package apifunc

import (
	"SystemEngineeringTeam/hack-teamA-2021-summer/dbfunc"
	"log"
	"reflect"

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

	u := c.Get("user").(*jwt.Token)
	claims := u.Claims.(jwt.MapClaims)
	uid := claims["uid"]
	log.Println(uid)

	user, err := dbfunc.PostCoin(params.Email, params.Total)
	if err != nil { //データベースでエラーが出た時の処理
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "データベースの更新に失敗しました: " + err.Error()})
	}

	log.Println(user.ID, uid, reflect.TypeOf(user.ID), reflect.TypeOf(uid))

	if user.ID != uid { //ユーザーIDが一致しなかった時の処理
		return c.JSON(http.StatusOK, map[string]interface{}{"message": "tokenが正しくありません"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "成功しました"}) // フロントに返す値
}
