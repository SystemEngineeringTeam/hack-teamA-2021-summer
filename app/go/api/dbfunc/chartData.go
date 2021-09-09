package dbfunc

import (
	"SystemEngineeringTeam/hack-teamA-2021-summer/models"

	"github.com/labstack/echo/v4"
)

type ChartData struct {
	spin int `json:"spin"`
	coin int `json:"coin"`
}

func GetChartData(c echo.Context) (data []map[string]int, err error) {
	db := sqlConnect()
	defer db.Close()

	var u models.User
	u, err = GetUserFromToken(c)
	if err != nil {
		return nil, err
	}

	d := []models.ChartData{}
	err = db.Find(&d, "user_uid = ?", u.UUID).Error
	if err != nil {
		return nil, err
	}
	for _, v := range d {
		data = append(data, map[string]int{"spin": v.Spin, "coin": v.Coin})
	}

	return data, err
}

func PostChartData(c echo.Context, spin int, coin int) (err error) {
	db := sqlConnect()
	defer db.Close()

	var u models.User
	u, err = GetUserFromToken(c)
	if err != nil {
		return err
	}

	d := models.ChartData{UserUID: u.UUID, Spin: spin, Coin: coin}
	err = db.Create(&d).Error
	if err != nil {
		return err
	}

	return err
}
