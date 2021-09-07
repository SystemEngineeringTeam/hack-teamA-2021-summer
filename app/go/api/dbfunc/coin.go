package dbfunc

import (
	"SystemEngineeringTeam/hack-teamA-2021-summer/models"
)

func PostCoin(email string, total int) (user models.User, err error) { //引数と返り値
	db := sqlConnect()
	defer db.Close()

	// データベースからユーザーが存在するかを確認
	var u models.User
	if err = db.Where("email = ?", email).First(&u).Error; err != nil {
		return u, err
	}
	// データベースを更新
	err = db.Model(models.User{}).Where("email = ?", email).Select("Coin").Updates(models.User{Coin: total}).Error

	return u, err
}
