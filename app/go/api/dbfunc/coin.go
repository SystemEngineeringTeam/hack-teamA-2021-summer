package dbfunc

import (
	"SystemEngineeringTeam/hack-teamA-2021-summer/models"

	"github.com/labstack/echo/v4"
)

func PostCoin(c echo.Context, total int) (err error) { //引数と返り値
	db := sqlConnect()
	defer db.Close()

	// データベースからユーザーが存在するかを確認
	user, err := GetUserFromToken(c)
	if err != nil {
		return err
	}
	// データベースを更新
	err = db.Model(models.User{}).Where("uuid = ?", user.UUID).Select("Coin").Updates(models.User{Coin: total}).Error

	return err
}
