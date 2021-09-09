package apifunc

import (
	"SystemEngineeringTeam/hack-teamA-2021-summer/dbfunc"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// success: return (json){email: (string), name: (string)}
// error: return (json){"message": (string)}
func ChartDataGet(c echo.Context) error {
	// 送られてきたデータを元にDBから取得する
	data, err := dbfunc.GetChartData(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "データが取得できませんでした: " + err.Error()})
	}

	log.Println(data)

	return c.JSON(http.StatusOK, map[string]interface{}{"data": data})
}

type ChartDataPostParams struct {
	Spin int `json:"spin" validate:"required"`
	Coin int `json:"coin" validate:"required"`
}

// success: return (json){"message": (string)}
// error: return (json){"message": (string)}
func ChartDataPost(c echo.Context) error {
	// 送られてきたJSONを確かめる
	var params ChartDataPostParams
	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "パラメータが正しくありません: " + err.Error()})
	}

	if err := c.Validate(&params); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "パラメータが不足しています: " + err.Error()})
	}

	// 送られてきたデータを元にDBに登録する
	if err := dbfunc.PostChartData(c, params.Spin, params.Coin); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "データベースへの登録に失敗しました:" + err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "success"})
}
