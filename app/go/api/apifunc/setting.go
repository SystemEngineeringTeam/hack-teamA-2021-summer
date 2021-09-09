package apifunc

import (
	"SystemEngineeringTeam/hack-teamA-2021-summer/dbfunc"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SettingPostParams struct {
	Path string `json:"path" validate:"required"`
}

func SettingPost(c echo.Context) error {
	var params SettingPostParams
	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "パラメータが正しくありません: " + err.Error()})
	}

	if err := c.Validate(&params); err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{"message": "パラメータが不足しています: " + err.Error()})
	}

	err := dbfunc.SettingPost(params.Path, c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "DBエラー: " + err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "success"})
}

func SettingGet(c echo.Context) error {
	s, err := dbfunc.SettingGet(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "DBエラー: " + err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"setting": s.Path})
}
