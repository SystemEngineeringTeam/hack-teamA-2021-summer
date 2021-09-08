package dbfunc

import (
	"SystemEngineeringTeam/hack-teamA-2021-summer/models"
)

type ChartData struct {
	spin int `json:"spin"`
	coin int `json:"coin"`
}

func GetChartData(email string) (data []map[string]int, err error) {
	db := sqlConnect()
	defer db.Close()

	var u models.User
	if err = db.Where("email = ?", email).First(&u).Error; err != nil {
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

func PostChartData(email string, spin int, coin int) (err error) {
	db := sqlConnect()
	defer db.Close()

	var u models.User
	if err = db.Where("email = ?", email).First(&u).Error; err != nil {
		return err
	}

	d := models.ChartData{UserUID: u.UUID, Spin: spin, Coin: coin}
	err = db.Create(&d).Error
	if err != nil {
		return err
	}

	return err
}
