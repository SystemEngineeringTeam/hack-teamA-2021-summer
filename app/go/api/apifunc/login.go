package apifunc

import (
	"SystemEngineeringTeam/hack-teamA-2021-summer/dbfunc"

	"net/http"

	"github.com/labstack/echo/v4"
)

type LoginPostParams struct {
	Email string `json:"email"`
	Password string `json:"password"`

}

func LoginPost(c echo.Context) error {
	var params LoginPostParams
	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "パラメータが正しくありません: " + err.Error()})
	}
	user,err := dbfunc.PostLoginInfo(params.Email, params.Password) 


	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "データベースのユーザの取得に失敗しました: " + err.Error()})
	}

	if !dbfunc.ComparePassword(user.Password,params.Password) {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "パスワードが正しくありません: " })
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "成功しました"}) // フロントに返す値

}
